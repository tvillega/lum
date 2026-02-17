package db

import (
	"database/sql"
	"errors"
)

func ConnectOrCreateDatabase() (*sql.DB, error) {
	var errMsg string
	db, err := sql.Open("sqlite3", "./lum.db?_pragma=foreign_keys(1)")
	if err != nil {
		errMsg = "Database connection failed"
		return nil, errors.New(errMsg)
	}

	return db, nil
}

func CreateTables() error {

	var err error
	var db *sql.DB

	db, err = ConnectOrCreateDatabase()
	if err != nil {
		return err
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	err = _createBooks(db)
	if err != nil {
		return err
	}

	return nil
}

func _createBooks(db *sql.DB) error {

	var err error
	query :=
		`CREATE TABLE IF NOT EXISTS bookshelf (
        id         INTEGER PRIMARY KEY AUTOINCREMENT,
        title      TEXT NOT NULL,
        creator    TEXT NOT NULL,
        publisher  TEXT,
        isbn       INTEGRER,
        path       TEXT NOT NULL);`

	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	return nil

}
