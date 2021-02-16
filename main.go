package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	routes := mux.NewRouter()
	routes.Handle("/metrics", promhttp.Handler()).Methods("GET")
	routes.HandleFunc("/foo", HelloWorld).Methods("GET")

	log.Println("Application is running on: 8080 .....")
	http.ListenAndServe(":8080", routes)
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello Foo!")
}
