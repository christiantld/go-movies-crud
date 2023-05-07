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
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			break
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var movie Movie

	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			movie.ID = params["id"]
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
		}
	}

}

func main() {
	r := mux.NewRouter()
	const port = ":8000"
	const moviesRoute = "/movies"

	movies = append(movies, Movie{
		ID:    "1",
		Isbn:  "438277",
		Title: "Guardians of The Galaxy 3",
		Director: &Director{
			FirstName: "James",
			LastName:  "Gunn",
		}})

	movies = append(movies, Movie{
		ID:    "2",
		Isbn:  "45455",
		Title: "The Irishman",
		Director: &Director{
			FirstName: "Martin",
			LastName:  "Scorsese",
		}})

	r.HandleFunc(moviesRoute, getMovies).Methods("GET")
	r.HandleFunc(moviesRoute+"/{id}", getMovie).Methods("GET")
	r.HandleFunc(moviesRoute, createMovie).Methods("POST")
	r.HandleFunc(moviesRoute+"/{id}", updateMovie).Methods("PUT")
	r.HandleFunc(moviesRoute+"/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Server Start at por %v\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
