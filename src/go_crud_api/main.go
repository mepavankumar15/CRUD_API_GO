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
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

// get function

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// delete movie function

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {

		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

// get by id function

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	for _, item := range movies {

		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

//create function

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)

}

// update function
func updateMovie(w http.ResponseWriter, r *http.Request) {
	//set the json type
	w.Header().Set("Content-Type", "application/json")
	// params
	params := mux.Vars(r)
	//loop over the movie , range
	for index, item := range movies {
		//delete the movie wiith the id you've sent
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			// add the new movie with the same id
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{
		ID: "1", Isbn: "438227", Title: "the hero",
		Director: &Director{Firstname: "rambo", Lastname: "Zambo"},
	})
	movies = append(movies, Movie{
		ID: "2", Isbn: "438228", Title: "the villian",
		Director: &Director{Firstname: "kili", Lastname: "Zambo"},
	})
	movies = append(movies, Movie{
		ID: "3", Isbn: "438229", Title: "the funny",
		Director: &Director{Firstname: "koni", Lastname: "Zambo"},
	})
	movies = append(movies, Movie{
		ID: "4", Isbn: "438220", Title: "the fantozy",
		Director: &Director{Firstname: "gibi", Lastname: "Zambo"},
	})
	movies = append(movies, Movie{
		ID: "5", Isbn: "438221", Title: "the show",
		Director: &Director{Firstname: "shibi", Lastname: "Zambo"},
	})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
