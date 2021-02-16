package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	routes := mux.NewRouter()
	routes.HandleFunc("/foo", HelloWorld).Methods("GET")

	log.Println("Application is running on: 8080 .....")
	http.ListenAndServe(":8080", routes)
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello Foo!")
}
