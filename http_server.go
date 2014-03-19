package main

import (
    "net/http"
    "io"
)

type String string

type Struct struct {
	Greeting string
   	Punct string
	Who string
}

func (s String) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello!\n")
}

func (s *Struct) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, s.Greeting)
}

func main() {
    // your http.Handle calls here
    http.Handle("/string", String("Gopher"))
    http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
    http.ListenAndServe("localhost:4000", nil)
}
