package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"tsis1-aot-api/api"

	"github.com/gorilla/mux"
)

var persons = []api.Person{
	{1, "Shayaa Bin ", " Abraham-Joseph", "21savage", "Bank Account"},
	{2, "Aubrey", "Drake", "Drake", "One Dance"},
	{3, "Adil", "Zhalelov", "skriptonit", "Jit` kak ya jivy"},
	{4, "Grigorey", "Lyahov", "Og Buda", "Kapli"},
	{5, "Tyler", "Leland ", "Metro Booming", "Space Cadet"},
	{6, " DeMun ", "Nayvadius", "Future", "Mask Off"},
	{7, "Olivier", "Noah ", "Yeat", "Out thë way"},
	{8, "Kendrick", "Lamar", "Kendrick Lamar", "HUMBLE."},
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "This is the health check endpoint. Thank u for visiting this website!")
}

func Persons(w http.ResponseWriter, r *http.Request) {
	response := api.Response{Persons: persons}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetPersonByNickname(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	nickname := params["nickname"]

	var targetPerson api.Person
	for _, person := range persons {
		if person.Nickname == nickname {
			targetPerson = person
			break
		}
	}

	if targetPerson.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(targetPerson)
}
