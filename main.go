package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//HomeHandler ... todo
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Orbit Maneuver Homepage!\n"))
}

func main() {
	var dir string

	flag.StringVar(&dir, "dir", "./http", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()

	r := mux.NewRouter()
	// Routes consist of a path and a handler function.

	r.HandleFunc("/", HomeHandler)
	// This will serve files under http://localhost:8000/static/<filename>
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
