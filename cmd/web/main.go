package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	// create default address like an env varaible
	addr := flag.String("addr", ":4000", "HTTP netwwork addewss")
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySWL data source name")
	flag.Parse()

	// Use log.New() to create a logger for writing information messages. This takes
	// three parameters: the destination to write the logs to (os.Stdout), a string
	// prefix for message (INFO followed by a tab), and flags to indicate what
	// additional information to include (local date and time). Note that the flags
	// are joined using the bitwise OR operator |.

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// creating error log
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
