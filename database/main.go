package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

func main() {
	db := NewDB()

	log.Println("Listeting on :8080")
	http.ListenAndServe(":8080", ShowBooks(db))
}

func ShowBooks(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var title, author string
		rows, err := db.Query("SELECT title, author FROM books")
		defer rows.Close()
		if err != nil {
			panic(err)
		}

		for rows.Next() {
			err := rows.Scan(&author, &title)
			if err != nil {
				panic(err)
			}

			fmt.Fprintf(rw, "The first book is '%s' by '%s'\n", title, author)
		}

		if err := rows.Err(); err != nil {
			panic(err)
		}
	})
}

func NewDB() *sql.DB {
	db, err := sql.Open("sqlite3", "example.sqlite")

	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS books(title text, author text)")
	if err != nil {
		panic(err)
	}

	return db
}
