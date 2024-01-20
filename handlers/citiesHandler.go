package handlers

import (
	"countryAPI/models"
	"encoding/json"
	"net/http"
	"os"
)

func CitiesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getCities(w,r)
	}
}

func getCities(w http.ResponseWriter, r *http.Request) {
	var citiesData []models.CityModel
	ByteCollection, _ := os.ReadFile("db/cities.json")
	json.Unmarshal(ByteCollection, &citiesData)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(citiesData)
}