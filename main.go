package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Receipt struct {
	RetailerName string  `json:"retailer_name"`
	Total        float64 `json:"total"`
	PurchaseDate string  `json:"purchase_date"`
	PurchaseTime string  `json:"purchase_time"`
	Items        []Item  `json:"items"`
}

type Item struct {
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type ProcessedReceipt struct {
	ID string `json:"id"`
}

type PointsResponse struct {
	Points int `json:"points"`
}

func processReceiptsHandler(w http.ResponseWriter, r *http.Request) {
	var receipt Receipt
	err := json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	id := generateID()
	processedReceipt := ProcessedReceipt{
		ID: id,
	}

	jsonResponse, err := json.Marshal(processedReceipt)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func getPointsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	receiptID := vars["id"]

	// Look up receiptID in in-memory data structure or database
	// Calculate and retrieve the points for the receiptID
	points := calculatePoints(receiptID)

	pointsResponse := PointsResponse{
		Points: points,
	}

	jsonResponse, err := json.Marshal(pointsResponse)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/receipts/process", processReceiptsHandler).Methods("POST")
	r.HandleFunc("/receipts/{id}/points", getPointsHandler).Methods("GET")

	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func calculatePoints(receiptID string) int {
	// Look up receiptID in in-memory data structure or database
	// Calculate and return the points for the receiptID
	return 32 // Replace with your logic to get the points for the receiptID
}

func generateID() string {
	return uuid.New().String()
}
