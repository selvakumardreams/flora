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
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Flora!\n"))
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/test", YourHandler)
	log.Fatal(http.ListenAndServe(":8000", r))
}
