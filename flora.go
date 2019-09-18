/*
Serve is a very simple static file server in go
Usage:
	-p="8300": port to serve on
	-d=".":    the directory of static files to host
Navigating to http://localhost:8300 will display the index.html or directory
listing file.
*/

// Reference: https://github.com/madhanganesh/taskpad 
package main

import (
	"flag"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/user/new", nil).Methods("POST")


	port := flag.String("p", "8300", "port to serve on")
	directory := flag.String("d", ".", "the directory of static file to host")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*directory)))

	log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
	log.Fatal(http.ListenAndServe(":"+*port, router))
}


