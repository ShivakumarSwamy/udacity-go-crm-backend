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

// Utility methods

func isCustomerIdExist(customerId string) bool {
	_, ok := customerDatabase[customerId]
	if ok {
		return true
	}
	return false
}

func setJsonContentType(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

func setStatus200Ok(writer http.ResponseWriter) {
	writer.WriteHeader(http.StatusOK)
}

func setStatus201Created(writer http.ResponseWriter) {
	writer.WriteHeader(http.StatusCreated)
}

func setStatus204NoContent(writer http.ResponseWriter) {
	writer.WriteHeader(http.StatusNoContent)
}

func setStatus404NotFound(writer http.ResponseWriter) {
	writer.WriteHeader(http.StatusNotFound)
}

func writeJsonBody(writer http.ResponseWriter, body any) error {
	return json.NewEncoder(writer).Encode(body)
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

// functions mapped with routes

func getCustomer(writer http.ResponseWriter, request *http.Request) {
	customerId := mux.Vars(request)["id"]

	if isCustomerIdExist(customerId) {
		setJsonContentType(writer)
		setStatus200Ok(writer)
		writeJsonBody(writer, customerDatabase[customerId])
	} else {
		setStatus404NotFound(writer)
		writeJsonBody(writer, "Unable to find the customer")
	}
}

func getCustomers(writer http.ResponseWriter, request *http.Request) {
	setJsonContentType(writer)
	setStatus200Ok(writer)
	writeJsonBody(writer, getAllValuesOfCustomerMap())
}

func addCustomer(writer http.ResponseWriter, request *http.Request) {
	setJsonContentType(writer)
	setStatus201Created(writer)

	var customerNewEntry Customer
	requestBody, _ := io.ReadAll(request.Body)
	json.Unmarshal(requestBody, &customerNewEntry)

	newCustomerId := uuid.New().String()
	customerNewEntry.Id = newCustomerId
	customerDatabase[newCustomerId] = customerNewEntry
	writeJsonBody(writer, customerDatabase[newCustomerId])
}

func updateCustomer(writer http.ResponseWriter, request *http.Request) {
	var updateCustomerEntry Customer
	requestBody, _ := io.ReadAll(request.Body)
	json.Unmarshal(requestBody, &updateCustomerEntry)

	customerId := mux.Vars(request)["id"]
	if isCustomerIdExist(customerId) {
		setJsonContentType(writer)
		setStatus200Ok(writer)
		customerDatabase[customerId] = updateCustomerEntry
		writeJsonBody(writer, customerDatabase[customerId])
	} else {
		setStatus404NotFound(writer)
		writeJsonBody(writer, "Unable to find the customer")
	}
}

func deleteCustomer(writer http.ResponseWriter, request *http.Request) {
	customerId := mux.Vars(request)["id"]
	if isCustomerIdExist(customerId) {
		delete(customerDatabase, customerId)
		setStatus204NoContent(writer)
	} else {
		setStatus404NotFound(writer)
		writeJsonBody(writer, "Unable to find the customer")
	}
}

// path mappings

func addCustomersHandlerMapping(router *mux.Router) {
	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	router.HandleFunc("/customers", getCustomers).Methods("GET")
	router.HandleFunc("/customers", addCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PUT")
	router.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")
}

func index(writer http.ResponseWriter, request *http.Request) {
	http.ServeFile(writer, request, "./static/index.html")
}

func main() {
	addInitialCustomerData()

	router := mux.NewRouter()
	addCustomersHandlerMapping(router)
	router.HandleFunc("/", index)
	fmt.Println("Server is starting at port 3000...")
	err := http.ListenAndServe(":3000", router)

	if err != nil {
		fmt.Println("Server unable to start due to", err)
	}
}
