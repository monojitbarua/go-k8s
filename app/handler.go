package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/monojitbarua/go-util-lib/zlogger"
)

type Customer struct {
	Name string `json:"name"`
	City string `json:"city"`
	Zip  int    `json:"zip"`
}

var customers = []Customer{
	{Name: "Monojit", City: "Bangalore", Zip: 560000},
	{Name: "Nilay", City: "Boston", Zip: 87000},
}

// healthcheck handler for k8s
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	zlogger.Info("Healthcheck calling")
	w.WriteHeader(http.StatusOK)
}

// readiness prob handler for k8s
func ReadinessHandler(w http.ResponseWriter, r *http.Request) {
	zlogger.Info("Readiness calling")
	w.WriteHeader(http.StatusOK)
}

// get handler for greet
func Handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Guest"
	}
	zlogger.Info(fmt.Sprintf("Received request for %s\n", name))

	w.Write([]byte(fmt.Sprintf("Hello, %s\n", name)))
}

// get all list of customer handler
func GetCustomersHandler(w http.ResponseWriter, r *http.Request) {
	zlogger.Info(fmt.Sprintf("Get customers calling, number of customer: %d", len(customers)))
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
