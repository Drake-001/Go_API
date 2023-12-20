package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// Register the Endpoint
	http.HandleFunc("/batterydata", getBatteryData)
	// 'http.HandleFunc' to handle HTTP requests.
	// 1 argument is path '/' which is the root of web server
	// 2 argument is function that defines how to handle requests to path
	// Function has 2 parameters: 1 'http.ResponseWriter' is used to write the HTTP response
	// 2 parameter is 'r', an '*http.Request' which represents incoming request
	// Inside function, 'fmt.Fprintf(w, "Data API Server")' writes string to response writer 'w'
	// which sends this string back to the client as part of HTTP response
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "EV Battery Data API Server")
	})
	// 'fmt.Println("...") prints message to console indicating where server is running
	// 'http.ListenAndServe(":8080", nil)' starts an HTTP server that listens to the TCP network address ":8080"
	// ":8080" specifies the PORT NUMBER on which the server listens. The ":" denotes a network address
	// "nil" means is uses the default HTTP request multiplexer ("http.DefaultServMux").
	// Since there is already a handler for root path ("/") using "http.HandleFunc" will handle any requests to that path
	fmt.Println("Server is running on http://localhost:8080")
	// This function call 'ListenAndServe' is a method of the http package
	http.ListenAndServe(":8080", nil)
}

// Visiting 'http://localhost:8080' should display "EV Battery Data API Server".

// Define struct for EV Battery Data
type BatteryData struct {
	ID                    int     `json:"id"`
	BatteryHealthPercent  float64 `json:"batteryHealthPercent"`
	ChargeCyclesNum       int     `json:"chargeCyclesNum"`
	TemperatureFahrenheit float64 `json:"temperatureFahrenheit"`
}

// Define function that reads JSON file and unmarshals data into slice of BatteryData struct
func getBatteryData(w http.ResponseWriter, r *http.Request) {
	// Open Json file
	jsonFile, err := os.Open("battery_data.json")
	if err != nil {
		fmt.Fprintf(w, "Error opening JSON file: %v", err)
		return
	}
	defer jsonFile.Close()

	// Read JSON data
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Fprintf(w, "Error reading JSON file: %v", err)
		return
	}

	// Unmarshal JSON data into a slice of BatteryData structs
	var batteries []BatteryData
	json.Unmarshal(byteValue, &batteries)

	// Send the JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(batteries)
}
