package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func main() {
	// serverv1()
	// serverv2()
	serverv3()
}

func serverv1() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "req.url.path=%q", r.URL.Path)
	})
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}

var counter int
var mu sync.Mutex

func serverv2() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		counter++
		mu.Unlock()
		fmt.Fprintf(w, "req.url.path=%q", r.URL.Path)
	})
	http.HandleFunc("/counter", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		fmt.Fprintf(w, "counter: %d", counter)
		mu.Unlock()
	})
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}

func serverv3() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s\t%s\t%s\n", r.URL, r.Method, r.Proto)

		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q]=%q\n", k, v)
		}

		fmt.Fprintf(w, "Host=%q\n", r.Host)
		fmt.Fprintf(w, "RemoteAddress=%q\n", r.RemoteAddr)

		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		for k, v := range r.Form {
			fmt.Fprintf(w, "Form[%q]=%q\n", k, v)
		}
	})
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
