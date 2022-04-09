package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP Network Address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	//i could have used log.Llongfile  full file path in the log output
	//I can also force the logger to use UTC datetimes, instead of local ones, by using the log.LUTC flag

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	server := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err := server.ListenAndServe()

	//err := http.ListenAndServe(*addr, mux)
	errorLog.Fatal(err)

}

//As a rule of thumb, you should avoid using the Panic() and Fatal()
//variations outside of your main() function — it’s good practice to
//return errors instead, and only panic or exit directly from main().
