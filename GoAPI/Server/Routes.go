package server

import (
	"fmt"
	"net/http"
)

func InitRoutes() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/countries", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			GetCountry(w, r)
		case http.MethodPost:
			CreateCountry(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method not allowed")
			return
		}
	})
}
