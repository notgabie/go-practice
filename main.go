package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
		// This function is an HTTP handler function. It takes two arguments: an http.ResponseWriter and an http.Request. The http.ResponseWriter is used to write the response back to the client, and the http.Request contains information about the incoming request.

	// The function signature is func(w http.ResponseWriter, r *http.Request). This is a common pattern in Go for defining HTTP handler functions. The asterisk (*) in *http.Request is a pointer type, which means that the function receives a reference to the request object rather than a copy of it.

	// the w and r are just variable names. However, it's a common convention to use w for the response writer and r for the request.

	// []byte is a built-in type that is used to store a sequence of bytes. We use it here to convert a string to a byte slice.

	// Strings are immutable in Go, which means that we can't change the contents of a string once it's created. Byte slices, on the other hand, are mutable, which means that we can modify the contents of a byte slice after it's created.
	w.Write([]byte("Hello, World!"))
}



func main() {
	// Create a new ServeMux which acts as a router in our application. We use the HandleFunc method to register the home function as the handler for the "/" route.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)


	// http.ListenAndServe starts an HTTP server and listens for incoming requests. It takes two arguments: the TCP address to listen on and the ServeMux to use for routing requests.

	// The first argument is a string in the format "host:port". The host can be an IP address or a domain name. The port is a number that specifies the port to listen on.

	// If you omit the host, ListenAndServe listens on all available network interfaces. 

	// The second argument is a ServeMux, which is an HTTP request multiplexer. It matches the URL of each incoming request against a list of registered patterns and calls the handler for the pattern that most closely matches the URL.

	// In case of an error, ListenAndServe returns an error object. We log the error using log.Fatal, which logs the error and exits the program.
	log.Print("Server starting on localhost:4000")
	err := http.ListenAndServe("localhost:4000", mux)
	log.Fatal(err)
}