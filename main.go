// Start HTTP server: this script starts server

// Every Go file starts with a package declaration.
// Main package defines a standalone executable (as opposed to library).
// Main package is entry point of the program (see main function below).

package main

// Imports two packages
// 'fmt' package used for formatted I/O operations, like printing to console.
// 'net/http' package provides HTTP clinet and server implementations.
// used to build web servers and handle HTTP requests and responses.

import (
	"fmt"
	"net/http"
)

// The main function is entry point of a Go program, wherein code inside
// main is executed.
// 'func' is used to declare a function. Here, 'main' is the function name.

func main() {

	// 'http.HandleFunc' is used to handle HTTP requests. Takes 2 arguments
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
