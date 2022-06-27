package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

	var history = make([]string, 0)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/calc/sum/{a}/{b}", Sum).Methods("GET")
	router.HandleFunc("/calc/sub/{a}/{b}", Sub).Methods("GET")
	router.HandleFunc("/calc/div/{a}/{b}", Div).Methods("GET")
	router.HandleFunc("/calc/mul/{a}/{b}", Mul).Methods("GET")
	router.HandleFunc("/calc/history", GetHistory).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}


func GetNumbers(r *http.Request) (float64, float64, error) {
	x, errorX := strconv.ParseFloat(mux.Vars(r)["a"], 32)
	y, errorY := strconv.ParseFloat(mux.Vars(r)["b"], 32)
	if errorX == nil && errorY == nil  {
		return x, y, nil
	} else {
		return 0, 0, errors.New("Falha, request contém letras.")
	}
}

func Sum(w http.ResponseWriter,  r *http.Request) {
	numberX, numberY, errorConvert := GetNumbers(r) 
	if errorConvert == nil {
		result := numberX + numberY
		history = append(history, fmt.Sprintf("%.2f + %.2f = %.2f", numberX, numberY, result))
		json.NewEncoder(w).Encode(result)
	} else {
		json.NewEncoder(w).Encode(errorConvert)
	}
}

func Sub(w http.ResponseWriter,  r *http.Request) {
	numberX, numberY, errorConvert := GetNumbers(r) 
	if errorConvert == nil {
		result := numberX - numberY
		history = append(history, fmt.Sprintf("%.2f - %.2f = %.2f", numberX, numberY, result))
		json.NewEncoder(w).Encode(result)
	} else {
		json.NewEncoder(w).Encode(errorConvert)
	}
}

func Mul(w http.ResponseWriter,  r *http.Request) {
	numberX, numberY, errorConvert := GetNumbers(r) 
	if errorConvert == nil {
		result := numberX * numberY
		history = append(history, fmt.Sprintf("%.2f * %.2f = %.2f", numberX, numberY, result))
		json.NewEncoder(w).Encode(result)
	} else {
		json.NewEncoder(w).Encode(errorConvert)
	}
}

func Div(w http.ResponseWriter,  r *http.Request) {
	var numberX, numberY, errorConvert = GetNumbers(r) 
	if errorConvert == nil {
		result := numberX / numberY
		if result == math.Inf(0) {
			json.NewEncoder(w).Encode("Infinity")	
		} else {
			json.NewEncoder(w).Encode(result)
		}
		history = append(history, fmt.Sprintf("%.2f / %.2f = %.2f", numberX, numberY, result))
	} else {
		json.NewEncoder(w).Encode(errorConvert)
	}
}

func GetHistory(w http.ResponseWriter,  r *http.Request) {
	json.NewEncoder(w).Encode(history)
}