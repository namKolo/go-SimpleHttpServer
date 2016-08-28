package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to my world")
}

func NoteIndex(w http.ResponseWriter, r *http.Request) {
	notes := Notes{
		Note{Name: "Kolo"},
		Note{Name: "Kolo #@"},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(notes); err != nil {
		panic(err)
	}

}
