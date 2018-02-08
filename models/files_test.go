// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

func testFiles(t *testing.T) {
	t.Parallel()

	query := Files(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testFilesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	file := &File{}
	if err = randomize.Struct(seed, file, fileDBTypes, true, fileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize File struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = file.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = file.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Files(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFilesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	file := &File{}
	if err = randomize.Struct(seed, file, fileDBTypes, true, fileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize File struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = file.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Files(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Files(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFilesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	file := &File{}
	if err = randomize.Struct(seed, file, fileDBTypes, true, fileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize File struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = file.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := FileSlice{file}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Files(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testFilesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	file := &File{}
	if err = randomize.Struct(seed, file, fileDBTypes, true, fileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize File struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = file.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := FileExists(tx, file.ID)
	if err != nil {
		t.Errorf("Unable to check if File exists: %s", err)
	}
	if !e {
		t.Errorf("Expected FileExistsG to return true, but got false.")
	}
}
func testFilesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	file := &File{}
	if err = randomize.Struct(seed, file, fileDBTypes, true, fileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize File struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = file.Insert(tx); err != nil {
		t.Error(err)
	}

	fileFound, err := FindFile(tx, file.ID)
	if err != nil {
		t.Error(err)
	}

	if fileFound == nil {
		t.Error("want a record, got nil")
	}
}
func testFilesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	file := &File{}
	if err = randomize.Struct(seed, file, fileDBTypes, true, fileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize File struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = file.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Files(tx).Bind(file); err != nil {
		t.Error(err)
	}
}

func testFilesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	file := &File{}
	if err = randomize.Struct(seed, file, fileDBTypes, true, fileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize File struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = file.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Files(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testFilesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	fileOne := &File{}
	fileTwo := &File{}
	if err = randomize.Struct(seed, fileOne, fileDBTypes, false, fileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize File struct: %s", err)
	}
	if err = randomize.Struct(seed, fileTwo, fileDBTypes, false, fileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize File struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = fileOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = fileTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Files(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testFilesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	fileOne := &File{}
	fileTwo := &File{}
	if err = randomize.Struct(seed, fileOne, fileDBTypes, false, fileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize File struct: %s", err)
	}
	if err = randomize.Struct(seed, fileTwo, fileDBTypes, false, fileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize File struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = fileOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = fileTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Files(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func fileBeforeInsertHook(e boil.Executor, o *File) error {
	*o = File{}
	return nil
}

func fileAfterInsertHook(e boil.Executor, o *File) error {
	*o = File{}
	return nil
}

func fileAfterSelectHook(e boil.Executor, o *File) error {
	*o = File{}
	return nil
}

func fileBeforeUpdateHook(e boil.Executor, o *File) error {
	*o = File{}
	return nil
}

func fileAfterUpdateHook(e boil.Executor, o *File) error {
	*o = File{}
	return nil
}

func fileBeforeDeleteHook(e boil.Executor, o *File) error {
	*o = File{}
	return nil
}

func fileAfterDeleteHook(e boil.Executor, o *File) error {
	*o = File{}
	return nil
}

func fileBeforeUpsertHook(e boil.Executor, o *File) error {
	*o = File{}
	return nil
}

func fileAfterUpsertHook(e boil.Executor, o *File) error {
	*o = File{}
	return nil
}

func testFilesHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &File{}
	o := &File{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, fileDBTypes, false); err != nil {
		t.Errorf("Unable to randomize File object: %s", err)
	}

	AddFileHook(boil.BeforeInsertHook, fileBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	fileBeforeInsertHooks = []FileHook{}

	AddFileHook(boil.AfterInsertHook, fileAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	fileAfterInsertHooks = []FileHook{}

	AddFileHook(boil.AfterSelectHook, fileAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	fileAfterSelectHooks = []FileHook{}

	AddFileHook(boil.BeforeUpdateHook, fileBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	fileBeforeUpdateHooks = []FileHook{}

	AddFileHook(boil.AfterUpdateHook, fileAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	fileAfterUpdateHooks = []FileHook{}

	AddFileHook(boil.BeforeDeleteHook, fileBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	fileBeforeDeleteHooks = []FileHook{}

	AddFileHook(boil.AfterDeleteHook, fileAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	fileAfterDeleteHooks = []FileHook{}

	AddFileHook(boil.BeforeUpsertHook, fileBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	fileBeforeUpsertHooks = []FileHook{}

	AddFileHook(boil.AfterUpsertHook, fileAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	fileAfterUpsertHooks = []FileHook{}
}
func testFilesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	file := &File{}
	if err = randomize.Struct(seed, file, fileDBTypes, true, fileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize File struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = file.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Files(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFilesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	file := &File{}
	if err = randomize.Struct(seed, file, fileDBTypes, true); err != nil {
		t.Errorf("Unable to randomize File struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = file.Insert(tx, fileColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := Files(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFilesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	file := &File{}
	if err = randomize.Struct(seed, file, fileDBTypes, true, fileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize File struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = file.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = file.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testFilesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	file := &File{}
	if err = randomize.Struct(seed, file, fileDBTypes, true, fileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize File struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = file.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := FileSlice{file}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testFilesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	file := &File{}
	if err = randomize.Struct(seed, file, fileDBTypes, true, fileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize File struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = file.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Files(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	fileDBTypes = map[string]string{`Alias`: `text`, `ContentType`: `text`, `CreatedAt`: `timestamp without time zone`, `Hash`: `text`, `ID`: `uuid`, `Name`: `text`, `Size`: `integer`, `Token`: `text`, `UpdatedAt`: `timestamp without time zone`}
	_           = bytes.MinRead
)

func testFilesUpdate(t *testing.T) {
	t.Parallel()

	if len(fileColumns) == len(filePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	file := &File{}
	if err = randomize.Struct(seed, file, fileDBTypes, true, fileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize File struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = file.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Files(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, file, fileDBTypes, true, fileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize File struct: %s", err)
	}

	if err = file.Update(tx); err != nil {
		t.Error(err)
	}
}

func testFilesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(fileColumns) == len(filePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	file := &File{}
	if err = randomize.Struct(seed, file, fileDBTypes, true, fileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize File struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = file.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Files(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, file, fileDBTypes, true, filePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize File struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(fileColumns, filePrimaryKeyColumns) {
		fields = fileColumns
	} else {
		fields = strmangle.SetComplement(
			fileColumns,
			filePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(file))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := FileSlice{file}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testFilesUpsert(t *testing.T) {
	t.Parallel()

	if len(fileColumns) == len(filePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	file := File{}
	if err = randomize.Struct(seed, &file, fileDBTypes, true); err != nil {
		t.Errorf("Unable to randomize File struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = file.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert File: %s", err)
	}

	count, err := Files(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &file, fileDBTypes, false, filePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize File struct: %s", err)
	}

	if err = file.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert File: %s", err)
	}

	count, err = Files(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}