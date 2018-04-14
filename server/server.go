package server

import (
	"crypto/tls"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/crypto/acme/autocert"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/zqz/upl/filedb"
)

type Server struct {
	config   config
	database *sql.DB
	logger   *log.Logger
}

func (s Server) Log(x ...interface{}) {
	s.logger.Println(x...)
}

func Init(path string) (Server, error) {
	s := Server{}
	s.logger = log.New(os.Stdout, "", log.LstdFlags)

	cfg, err := parseConfig(path)
	if err != nil {
		return s, err
	}

	if cfg.LogFile != "" {
		f, err := os.OpenFile(cfg.LogFile, os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			return s, err
		}

		s.logger = log.New(f, "", log.LstdFlags)
	}

	s.logger.Println("Parsed Config")

	db, err := cfg.DBConfig.loadDatabase()
	if err != nil {
		return s, err
	}

	s.logger.Println("Connected to DB")

	s.database = db
	s.config = cfg

	return s, nil
}

func (s Server) Close() {
	s.database.Close()
}

func (s Server) runInsecure(r chi.Router) error {
	listenPort := fmt.Sprintf(":%d", s.config.Port)

	s.logger.Println("[server] listening for HTTP traffic on port", listenPort)

	return http.ListenAndServe(listenPort, r)
}

func (s Server) runSecure(r chi.Router) error {
	c := autocert.DirCache("./")
	m := autocert.Manager{
		Cache:  c,
		Prompt: autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(
			"zqz.ca",
		),
	}

	tlsPort := fmt.Sprintf(":%d", s.config.TLSPort)

	h := &http.Server{
		Addr:      tlsPort,
		TLSConfig: &tls.Config{GetCertificate: m.GetCertificate},
		Handler:   r,
		ErrorLog:  s.logger,
	}

	listenPort := fmt.Sprintf(":%d", s.config.Port)
	go http.ListenAndServe(listenPort, m.HTTPHandler(s.secureRedirect()))

	s.logger.Println("[server] listening for TLS traffic on port", s.config.TLSPort)
	s.logger.Println("[server] redirecting HTTP traffic on port", s.config.Port, "to HTTPS")

	// start https server
	return h.ListenAndServeTLS("", "")
}

func (s Server) Run() error {
	fdb := filedb.NewServer(
		filedb.NewFileDB(
			filedb.NewDiskPersistence(),
			filedb.NewDBMetaStorage(s.database),
			filedb.NewDBThumbnailStorage(s.database),
		),
	)

	r := chi.NewRouter()

	logger := middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: s.logger})
	r.Use(logger)

	r.Mount("/api", fdb.Router())

	r.Get("/*", serveIndex)
	serveAssets(r)

	s.logger.Println("Listening for web traffic")

	return s.run(r)
}

func (s Server) run(r chi.Router) error {
	if s.config.Secure {
		return s.runSecure(r)
	} else {
		return s.runInsecure(r)
	}
}

func (s Server) secureRedirect() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		redir := "https://" + req.Host + req.RequestURI

		s.logger.Println("[server] redirected request to", redir)
		http.Redirect(w, req, redir, http.StatusMovedPermanently)
	}
}