package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Vehicle struct {
	Id    int
	Make  string
	Model string
	Price int
}

var vehicles = []Vehicle{
	{1, "Toyota", "Corolla", 10000},
	{2, "Toyota", "Camry", 15000},
	{3, "Honda", "Accord", 13000},
}

func getAllCars(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vehicles)
}

func getCarsByBrand(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	carMaker := vars["make"]
	cars := &[]Vehicle{}
	for _, car := range vehicles {
		if car.Make == carMaker {
			*cars = append(*cars, car)
		}
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cars)
}

func getCarById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	carId, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println("Unable to convert id to string")
	}
	for _, car := range vehicles {
		if car.Id == carId {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(car)
		}
	}
}

func updateCar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	carId, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println("Unable to convert id to string")
	}
	var updatedCar Vehicle
	json.NewDecoder(r.Body).Decode(&updatedCar)
	for index, car := range vehicles {
		if car.Id == carId {
			vehicles = append(vehicles[:index], vehicles[index+1:]...)
			vehicles = append(vehicles, updatedCar)
		}
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vehicles)
}

func createCar(w http.ResponseWriter, r *http.Request) {
	var newCar Vehicle
	json.NewDecoder(r.Body).Decode(&newCar)
	vehicles = append(vehicles, newCar)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vehicles)
}

func deleteCar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	carId, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println("Unable to convert id to string")
	}
	for index, car := range vehicles {
		if car.Id == carId {
			vehicles = append(vehicles[:index], vehicles[index+1:]...)
		}
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vehicles)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/cars", getAllCars).Methods("GET")
	router.HandleFunc("/cars/make/{make}", getCarsByBrand).Methods("GET")
	router.HandleFunc("/cars/{id}", getCarById).Methods("GET")
	router.HandleFunc("/cars/{id}", updateCar).Methods("PUT")
	router.HandleFunc("/cars", createCar).Methods("POST")
	router.HandleFunc("/cars/{id}", deleteCar).Methods("DELETE")

	http.ListenAndServe(":8081", router)
}
