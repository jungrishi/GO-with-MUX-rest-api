package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jungrishi/rest_api/resources"
)

func main() {
	//Init Router
	r := mux.NewRouter()

	//Route Handler
	r.HandleFunc("/api/books", resources.GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", resources.GetBook).Methods("GET")
	r.HandleFunc("/api/books", resources.CreateBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", resources.UpdateBooks).Methods("PUT")
	r.HandleFunc("/api/books/{id}", resources.DeleteBooks).Methods("DELETE")
	r.HandleFunc("/api/author/{id}", resources.GetAuthor).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
