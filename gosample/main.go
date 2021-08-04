package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	errs := make(chan error)
	r.HandleFunc("/ping", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte(`{"ping": "pong"}`))
	})
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT)
	go func() {
		fmt.Println("App starting")
		if err := http.ListenAndServe(":8080", r); err != nil {
			errs <- err
			return
		}
	}()
	go func() {
		sig := <-c
		errs <- fmt.Errorf("Application finish with status %v", sig)
	}()
	fmt.Println("App started in :8080")
	fmt.Println(<-errs)
}
