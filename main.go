package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	println("running")
	router := mux.NewRouter()
	router.HandleFunc("/health", Health).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK) // Set 200 OK
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`) // Send json to the ResponseWriter
}
