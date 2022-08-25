package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var randomKillTime = rand.Intn(300-60) + 60

func ping(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Pong\n")
}

func main() {
	http.HandleFunc("/ping", ping)
	go func() {
		fmt.Printf("A WARRIORS DEATH COME IN %d sec\n", randomKillTime)
		time.Sleep(time.Duration(randomKillTime) * time.Second)
		panic(fmt.Sprintf("SELF KILL AFTER %d sec\n", randomKillTime))
	}()
	http.ListenAndServe(":8080", nil)
}
