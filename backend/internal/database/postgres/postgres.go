package postgres

import (
	"database/sql"
	"log"
)

func Connect(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(
		"CREATE TABLE IF NOT EXISTS users (id serial primary key, name text, email text)",
	)
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}
