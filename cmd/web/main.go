package main

import (
	"database/sql"
	"log/slog"
	"net/http"
	"os"

	"github.com/ColeBurch/snippetbox/internal/models"
	_ "github.com/lib/pq"
)

type application struct {
	logger   *slog.Logger
	snippets *models.SnippetModel
}

func main() {
	connStr := "user=myuser dbname=mydb password=mypassword host=localhost port=5432 sslmode=disable"
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))

	db, err := openDB(connStr)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()

	app := &application{
		logger:   logger,
		snippets: &models.SnippetModel{DB: db},
	}

	logger.Info("starting server on :8000")
	err = http.ListenAndServe(":8000", app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
