package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"sync"
)

func main() {
	http.Handle("/foo", new(countHandler))

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/meow", handleMeow)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

type countHandler struct {
	mu sync.Mutex // guards nfmt.Fprintf(w, "Hello, foo")
	n  int
}

func (h *countHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.n++
	type result struct {
		N int  `json:"visited"`
		H bool `json:"is_even"`
	}

	resp := result{N: h.n, H: bool(h.n%2 == 0)}
	bytes, err := json.Marshal(resp)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

func handleMeow(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
