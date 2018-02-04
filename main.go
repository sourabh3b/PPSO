package main

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"fmt"
)

//TestRoute - test route
func TestRoute(w http.ResponseWriter, r *http.Request) {
	//render := render.New()
	fmt.Fprint(w, "Hello World !")
	//render.JSON(w, http.StatusOK, nil)
	return
}

func main() {
	log.Info("Starting PPSO ....")
	http.HandleFunc("/test", TestRoute)
	http.ListenAndServe(":8889", nil)
	//todo : create API to calculate PSO fitness function value
}