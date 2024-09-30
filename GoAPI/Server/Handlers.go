package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Hello there, %s", "visitor")
}

func GetCountry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(countries)
}

func CreateCountry(w http.ResponseWriter, r *http.Request) {
	country := &Country{}
	err := json.NewDecoder(r.Body).Decode(country)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid request payload")
		return
	}
	countries = append(countries, country)
	fmt.Fprintf(w, "Country created successfully")

}
