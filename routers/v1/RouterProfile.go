package v1

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	v1 "github.com/suhriar/submission_maps_api/controllers/v1"
	"github.com/suhriar/submission_maps_api/errHandler"
)

func RouterProfile(r *mux.Router) {
	// set attendance prefix
	maps := r.PathPrefix("/map").Subrouter()

	// attendance
	maps.HandleFunc("", v1.GetPlaces).Methods("GET")

	//PING
	maps.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := map[string]string{
			"Code":    "1",
			"Message": "PONG",
		}
		err := json.NewEncoder(w).Encode(response)
		errHandler.ErrHandler("Error json response: ", err)
	}).Methods("GET")
}
