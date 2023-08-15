package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type PersonInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var Persons []PersonInfo

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/app/v1/create", createHandler).Methods("POST")
	router.HandleFunc("/app/v1/read", readAllHandler).Methods("GET")
	router.HandleFunc("/app/v1/read/{name}", readHandler).Methods("GET")
	router.HandleFunc("/app/v1/update/{name}", updatehandler).Methods("PUT")
	router.HandleFunc("/app/v1/delete/{name}", deleteHandle).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}

func createHandler(writer http.ResponseWriter, request *http.Request) {
	var person PersonInfo
	if err := json.NewDecoder(request.Body).Decode(&person); err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		log.Fatalln("Error While Reading Request Body", err)
	}
	Persons = append(Persons, person)
	if err := json.NewEncoder(writer).Encode(&Persons); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Fatalln("Error While Encoding the initialize struct", err)
	}
	writer.Header().Set("Content-Type", "application/json")
}

func readAllHandler(writer http.ResponseWriter, request *http.Request) {

	if err := json.NewEncoder(writer).Encode(&Persons); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Fatalln("Error While Encoding the initialize struct", err)
	}
	writer.Header().Set("Content-Type", "application/json")
}

func readHandler(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	name := params["name"]
	var resultPerson PersonInfo

	for _, person := range Persons {
		if person.Name == name {
			resultPerson = person
		}
	}

	if err := json.NewEncoder(writer).Encode(&resultPerson); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Fatalln("Error While Encoding result", err)
	}
	writer.Header().Set("Content-Type", "application-json")
}

func updatehandler(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	name := params["name"]

	var person PersonInfo
	if err := json.NewDecoder(request.Body).Decode(&person); err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		log.Fatalln("Error While Reading Request Body", err)
	}

	for index, personDetails := range Persons {
		if personDetails.Name == name {
			Persons = append(Persons[:index], Persons[index+1:]...)
		}
	}
	Persons = append(Persons, person)
	var UpdatedPerson PersonInfo
	for _, personDetails := range Persons {
		if personDetails.Name == name {
			UpdatedPerson = personDetails
		}
	}
	if err := json.NewEncoder(writer).Encode(&UpdatedPerson); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Fatalln("Error While Encoding the initialize struct", err)
	}
	writer.Header().Set("Content-Type", "application/json")
}

func deleteHandle(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	name := params["name"]

	for index, personDetails := range Persons {
		if personDetails.Name == name {
			Persons = append(Persons[:index], Persons[index+1:]...)
		}
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.Write([]byte("Succussfully Deleted"))
}
