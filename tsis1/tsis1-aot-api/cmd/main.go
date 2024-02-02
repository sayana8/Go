package main

import (
	"log"
	"net/http"
	handlers "tsis1-aot-api/pkg"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("starting API server")
	router := mux.NewRouter()
	log.Println("Creating routes!")

	router.HandleFunc("/health-check", handlers.HealthCheck).Methods("GET")
	router.HandleFunc("/persons", handlers.Persons).Methods("GET")
	router.HandleFunc("/persons/{nickname}", handlers.GetPersonByNickname).Methods("GET")

	http.Handle("/", router)

	http.ListenAndServe(":8080", nil)
}
