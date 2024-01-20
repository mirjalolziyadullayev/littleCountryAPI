package handlers

import (
	"countryAPI/models"
	"encoding/json"
	"net/http"
	"os"
)

func PlacesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getPlaces(w, r)
	}
}
func getPlaces(w http.ResponseWriter, r *http.Request) {
	var placesData []models.PlacesModel
	ByteCollection, _ := os.ReadFile("db/places.json")
	json.Unmarshal(ByteCollection, &placesData)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(placesData)
}