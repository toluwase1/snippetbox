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
	addr:= flag.String("addr", ":4000", "HTTP Network Address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog:= log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	//i could have used log.Llongfile  full file path in the log output
	//I can also force the logger to use UTC datetimes, instead of local ones, by using the log.LUTC flag

	app:= &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	// Create a file server which serves files out of the "./ui/static" directo
	// Note that the path given to the http.Dir function is relative to the pro
	// directory root.
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server:= &http.Server{
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: mux,
	}


	infoLog.Printf("Starting server on %s", *addr)
	err:= server.ListenAndServe()

	//err := http.ListenAndServe(*addr, mux)
	errorLog.Fatal(err)

}


//As a rule of thumb, you should avoid using the Panic() and Fatal()
//variations outside of your main() function — it’s good practice to
//return errors instead, and only panic or exit directly from main().