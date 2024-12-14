package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/jk.chen/posts/internal/models/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

type app struct {
	post *sqlite.PostModel
}

func main() {
	fmt.Println("Hello World")
	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Fatal(err)
	}

	app := app{post: &sqlite.PostModel{DB: db}}

	srv := http.Server{
		Addr:    ":8000",
		Handler: app.routes(),
	}

	log.Println("Listening on :8000")

	srv.ListenAndServe()
}
