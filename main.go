package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte ("Hello from Snippet"))
}

func showSnippet(w http.ResponseWriter, r *http.Request)  {
	id, err:=strconv.Atoi(r.URL.Query().Get("ID"))
	if err !=nil || id<0 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
	//w.Write([]byte("print show snippet"))
}
func createSnippet(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Allow", "POST")
	if r.Method != "POST" {
		//w.WriteHeader(405)
		//w.Write([]byte("Method Not allowed"))
		http.Error(w, "Method Not allowed", 405)
		return
	}

	w.Write([]byte ("create snippet"))
}

func main()  {
	mux:= http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Starting server on port :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}