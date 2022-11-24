package main

import (
	"fmt"
	"net/http"
)

func main() {
	//using standard library means no external framework or lib.
	// default mux has two func - one to regsiter routes (Handle func) and other to start server
	//Handle Func takes two parameters, first is route/ endpoint and second is handler func.
	// in handler func we have two params, response writer that writes back to client and second is  request reader that takes request from client
	//we will use Fprint func to send the response to the client, thid takes two params: one is IO writer and other is data that we want to send to IO writer
	//http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "Hello World\n") })
	http.HandleFunc("/greet", greet)
	//Listen and Serve starts the server at localhost:port and second parameter is handler which is a request multiplexer, and since we are relying on default so we will keep it as nil.
	http.ListenAndServe("localhost:9000", nil)
	// after running main. go
	// curl http://localhost:9000/greet
	//	Hello World
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World\n")
}
