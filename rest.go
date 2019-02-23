package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// user struct (Model)
type user struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Name *Name `json:"Name"`
}

// Name struct
type Name struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init users var as a slice user struct
var users []User

// Get all users
func getusers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Get single user
func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	// Loop through users and find one with the id from the params
	for _, item := range users {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&user{})
}

// Add new user
func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user user
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.ID = strconv.Itoa(rand.Intn(100000000)) // Mock ID - not safe
	users = append(users, user)
	json.NewEncoder(w).Encode(user)
}

// Update user
func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range users {
		if item.ID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			var user user
			_ = json.NewDecoder(r.Body).Decode(&user)
			user.ID = params["id"]
			users = append(users, user)
			json.NewEncoder(w).Encode(user)
			return
		}
	}
}


// Main function
func main() {
	// Init router
	r := mux.NewRouter()

	// Hardcoded data - @todo: add database
	users = append(users, user{ID: "1", Isbn: "438227", Title: "user One", Name: &Name{Firstname: "John", Lastname: "Doe"}})
	users = append(users, user{ID: "2", Isbn: "454555", Title: "user Two", Name: &Name{Firstname: "Steve", Lastname: "Smith"}})

	// Route handles & endpoints
	r.HandleFunc("/users", getusers).Methods("GET")
	r.HandleFunc("/users/{id}", getuser).Methods("GET")
	r.HandleFunc("/users", createuser).Methods("POST")
	r.HandleFunc("/users/{id}", updateuser).Methods("PUT")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}

