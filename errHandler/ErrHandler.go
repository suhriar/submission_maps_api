package errHandler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/suhriar/submission_maps_api/models"
)

func ErrHandler(s string, e error) {
	if e != nil {
		log.Printf("%s : %s", s, e)
	}
}

func ErrorResponse(w http.ResponseWriter, code int, httpCode int, message string) {
	// set the http code response
	w.WriteHeader(httpCode)

	// set the response content
	response := models.ResultKeywords{
		Code:    code,
		Message: message,
	}

	// encode the result to json
	err := json.NewEncoder(w).Encode(response)
	ErrHandler("Error json response: ", err)
}
