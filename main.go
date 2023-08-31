package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

type Customer struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email,omitempty"`
	Phone     uint64 `json:"phone,omitempty"`
	Contacted bool   `json:"contacted"`
}

var customerDatabase = make(map[string]Customer)

func addInitialCustomerData() {
	customerDatabase["431baecf-6535-452f-884e-1da18ff0d5a2"] = Customer{
		Id:        "431baecf-6535-452f-884e-1da18ff0d5a2",
		Name:      "Foo",
		Role:      "ADMIN",
		Email:     "foo@company.com",
		Phone:     56792834,
		Contacted: false,
	}

	customerDatabase["4eb4af25-b104-4aa2-a321-9ab671c4fa36"] = Customer{
		Id:        "4eb4af25-b104-4aa2-a321-9ab671c4fa36",
		Name:      "Alice",
		Role:      "USER",
		Email:     "alice@company.com",
		Phone:     376849763,
		Contacted: true,
	}

	customerDatabase["78ee45a9-7626-4c83-9a97-b572c9c318eb"] = Customer{
		Id:        "78ee45a9-7626-4c83-9a97-b572c9c318eb",
		Name:      "Bob",
		Role:      "USER",
		Email:     "bob@company.com",
		Phone:     273598645,
		Contacted: false,
	}
}

func isCustomerIdExist(customerId string) bool {
	_, ok := customerDatabase[customerId]
	if ok {
		return true
	}
	return false
}

func getCustomer(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request)
	customerId := vars["id"]

	if isCustomerIdExist(customerId) {
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(customerDatabase[customerId])
	} else {
		writer.WriteHeader(http.StatusNotFound)
	}

}

func getAllValuesOfCustomerMap() []Customer {
	allCustomer := make([]Customer, len(customerDatabase))

	index := 0
	for customerId := range customerDatabase {
		allCustomer[index] = customerDatabase[customerId]
		index++
	}
	return allCustomer
}

func getCustomers(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(getAllValuesOfCustomerMap())
}

func addCustomer(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)

	var customerNewEntry Customer
	requestBody, _ := io.ReadAll(request.Body)
	json.Unmarshal(requestBody, &customerNewEntry)

	newCustomerId := uuid.New().String()
	customerNewEntry.Id = newCustomerId
	customerDatabase[newCustomerId] = customerNewEntry
	json.NewEncoder(writer).Encode(customerDatabase[newCustomerId])
}

func updateCustomer(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var updateCustomerEntry Customer
	requestBody, _ := io.ReadAll(request.Body)
	json.Unmarshal(requestBody, &updateCustomerEntry)

	vars := mux.Vars(request)
	customerId := vars["id"]

	if isCustomerIdExist(customerId) {
		customerDatabase[customerId] = updateCustomerEntry
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(customerDatabase[customerId])
	} else {
		writer.WriteHeader(http.StatusNotFound)
	}
}

func deleteCustomer(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	customerId := vars["id"]
	if isCustomerIdExist(customerId) {
		delete(customerDatabase, customerId)
		writer.WriteHeader(http.StatusNoContent)
	} else {
		writer.WriteHeader(http.StatusNotFound)
	}
}

func main() {
	addInitialCustomerData()

	router := mux.NewRouter()
	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	router.HandleFunc("/customers", getCustomers).Methods("GET")
	router.HandleFunc("/customers", addCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PUT")
	router.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")

	fmt.Println("Server is starting at port 3000...")
	err := http.ListenAndServe(":3000", router)

	if err != nil {
		fmt.Println("Server unable to start due to", err)
	}
}
