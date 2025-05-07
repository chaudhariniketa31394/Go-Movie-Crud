package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	Id       string    `json:"id"`
	Isbn     string    `json:"isbn:`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)

}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {

		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.Id = strconv.Itoa(rand.Intn(100000000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)

}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.Id = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}

}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	idToDelete := params["id"]

	found := false
	for index, item := range movies {
		if item.Id == idToDelete {
			movies = append(movies[:index], movies[index+1:]...)
			found = true
			break
		}
	}

	if found {
		json.NewEncoder(w).Encode(map[string]string{"message": "Movie deleted successfully"})
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Movie not found"})
	}
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{Id: "1", Isbn: "438221", Title: "Movie One", Director: &Director{Firstname: "john", Lastname: "Done"}})
	movies = append(movies, Movie{Id: "2", Isbn: "438222", Title: "Movie Two", Director: &Director{Firstname: "john1", Lastname: "Done"}})

	movies = append(movies, Movie{Id: "3", Isbn: "438223", Title: "Movie Three", Director: &Director{Firstname: "john2", Lastname: "Done"}})

	movies = append(movies, Movie{Id: "4", Isbn: "438224", Title: "Movie Four", Director: &Director{Firstname: "john3", Lastname: "Done"}})

	movies = append(movies, Movie{Id: "5", Isbn: "438225", Title: "Movie Five", Director: &Director{Firstname: "john4", Lastname: "Done"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movie", createMovie).Methods("POSt")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Start server at pot\rt 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
