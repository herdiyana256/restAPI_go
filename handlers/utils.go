package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func writeJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(data)
}

func readJSONRequest(r *http.Request, data interface{}) error {
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(data)
}

func checkAndParseOrderID(w http.ResponseWriter, r *http.Request) (int, error) {
	orderIDStr := mux.Vars(r)["orderID"]
	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		http.Error(w, "Invalid order ID format", http.StatusBadRequest)
		return 0, err
	}
	return orderID, nil
}
