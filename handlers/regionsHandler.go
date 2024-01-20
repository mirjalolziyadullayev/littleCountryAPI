package handlers

import (
	"countryAPI/models"
	"encoding/json"
	"net/http"
	"os"
)

func RegionsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getRegions(w,r)
	}
}
func getRegions(w http.ResponseWriter, r *http.Request) {
	var regionsData []models.CityModel
	ByteCollection, _ := os.ReadFile("db/regions.json")
	json.Unmarshal(ByteCollection, &regionsData)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(regionsData)
}