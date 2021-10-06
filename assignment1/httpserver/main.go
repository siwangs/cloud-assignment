package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
)

func healthz(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "200\n")
}

type StatusRecorder struct {
	http.ResponseWriter
	Status int
}

func (r *StatusRecorder) WriteHeader(status int) {
	r.Status = status
	r.ResponseWriter.WriteHeader(status)
}

func DisplayInfo(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		recorder := &StatusRecorder{
			ResponseWriter: w,
			Status:         200,
		}
		h.ServeHTTP(recorder, r)
		log.Printf("Handling request for %s from %s, status: %d", r.URL.Path, r.RemoteAddr, recorder.Status)
	})
}
func headers(w http.ResponseWriter, req *http.Request) {

	if req.URL.Path != "/headers" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	fmt.Printf("header get request\n")

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
	fmt.Fprintf(w, "%v: %v\n", "VERSION", runtime.Version())
	fmt.Printf(req.RemoteAddr)

}

func main() {
	fmt.Printf("Starting Server at 8080\n")
	headerhandler := DisplayInfo(http.HandlerFunc(headers))
	http.HandleFunc("/healthz", healthz)
	http.Handle("/headers", headerhandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
