package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	     
	"github.com/gorilla/mux"
	"github.com/google/uuid"
)

type Customer struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Contacted bool   `json:"contacted"`
}
// 2. Mock database (slice)
var customers []Customer

// 3. Seed function (preload data)
func seedData() {
	customers = []Customer{
		{
			ID:        uuid.New().String(),
			Name:      "Alice Johnson",
			Role:      "Manager",
			Email:     "alice@example.com",
			Phone:     "123-456-7890",
			Contacted: true,
		},
		{
			ID:        uuid.New().String(),
			Name:      "Bob Smith",
			Role:      "Developer",
			Email:     "bob@example.com",
			Phone:     "234-567-8901",
			Contacted: false,
		},
		{
			ID:        uuid.New().String(),
			Name:      "Charlie Brown",
			Role:      "Designer",
			Email:     "charlie@example.com",
			Phone:     "345-678-9012",
			Contacted: true,
		},
	}
}


func getCustomers(w http.ResponseWriter, r *http.Request) {
    fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(customers)
}


func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	for i, c := range customers {
		if c.ID == id {
			customers = append(customers[:i], customers[i+1:]...)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(customers)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"message": "Customer not found"})
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newCustomer Customer

	err := json.NewDecoder(r.Body).Decode(&newCustomer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	// Generate UUID
	newCustomer.ID = uuid.New().String()

	customers = append(customers, newCustomer)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCustomer)
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	var updatedCustomer Customer

	err := json.NewDecoder(r.Body).Decode(&updatedCustomer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	for i, c := range customers {
		if c.ID == id {
			updatedCustomer.ID = c.ID
			customers[i] = updatedCustomer

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(updatedCustomer)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"message": "Customer not found"})
}

func main() {
	seedData() // <-- IMPORTANT

	router := mux.NewRouter().StrictSlash(true)
    
	fileServer := http.FileServer(http.Dir("./static"))
	router.Handle("/", fileServer).Methods("GET")

	router.HandleFunc("/customers", getCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", getCustomers).Methods("GET")
	router.HandleFunc("/customers", createCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PUT")
	router.HandleFunc("/deleteCustomer/{id}", deleteCustomer).Methods("DELETE")
    

	fmt.Println("Server is starting on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", router))
}