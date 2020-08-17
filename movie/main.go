package main

import (
	"encoding/json"
	"log"
	"net/http"
)

//Addr var is for server port assignment
var Addr = ":8080"

func main() {
	http.HandleFunc("/movies", movieListHandler)
	log.Printf("Starting on port %s", Addr)
	log.Fatal(http.ListenAndServe(Addr, nil))
}

//Movie struct for movie object
type Movie struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Poster   string `json:"poster"`
	MovieURL string `json:"movie_url"`
	IsPaid   bool   `json:"is_paid"`
}

func movieListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf8")
	mm := []Movie{
		Movie{0, "Бойцовский клуб", "/static/posters/fightclub.jpg", "https://youtu.be/qtRKdVHc-cE", true},
		Movie{1, "Крестный отец", "/static/posters/father.jpg", "https://youtu.be/ar1SHxgeZUc", false},
		Movie{2, "Криминальное чтиво", "/static/posters/pulpfiction.jpg", "https://youtu.be/s7EdQ4FqbhY", true},
	}
	err := json.NewEncoder(w).Encode(mm)
	if err != nil {
		log.Printf("Render response error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	return
}
