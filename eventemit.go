package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// func init() {
// 	var events []event
// }

var events []eventType
var messageChannel chan eventType

type eventType struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

//fire event
func createEventForFire(w http.ResponseWriter, r *http.Request) {
	newEvent := &eventType{}
	newEvent.ID = "ID-" + strconv.Itoa(rand.Intn(100000))
	newEvent.Title = "FireEvent"
	newEvent.Description = "Notifying for a fire spread in the area"

	jsonAsBytes, _ := json.Marshal(*newEvent)
	err := sendEventToHandler(jsonAsBytes)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}
	fmt.Println(*newEvent)
	events = append(events, *newEvent)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newEvent)
	fmt.Println("POST succesfull for fire event")
}

//theft event
func createEventForTheft(w http.ResponseWriter, r *http.Request) {
	newEvent := &eventType{}
	newEvent.ID = "ID-" + strconv.Itoa(rand.Intn(100000))
	newEvent.Title = "TheftEvent"
	newEvent.Description = "Notifying for a theft in the area"

	jsonAsBytes, _ := json.Marshal(*newEvent)
	err := sendEventToHandler(jsonAsBytes)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}
	fmt.Println(*newEvent)
	events = append(events, *newEvent)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newEvent)
	fmt.Println("POST succesfull for THEFT event")
}

//murder event
func createEventForMurder(w http.ResponseWriter, r *http.Request) {
	newEvent := &eventType{}
	newEvent.ID = "ID-" + strconv.Itoa(rand.Intn(100000))
	newEvent.Title = "MurderEvent"
	newEvent.Description = "Notifying for a murder occurred in the area"

	jsonAsBytes, _ := json.Marshal(*newEvent)
	err := sendEventToHandler(jsonAsBytes)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}
	fmt.Println(*newEvent)
	events = append(events, *newEvent)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newEvent)
	fmt.Println("POST succesfull for MURDER event")
}

func sendEventToHandler(jsonValue []byte) error {
	response, err := http.Post("http://0.0.0.0:8081/event", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
		return nil
	}
}

func getOneEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	for _, singleEvent := range events {
		if singleEvent.ID == eventID {
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
	fmt.Println("GET succesfull")
}

func getAllEvents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(events)
	fmt.Println("GET ALL succesfull")
}

func main() {
	// emitter := New()
	// e := &emitter.Emitter{}
	// var eventMessage chan eventType
	router := mux.NewRouter().StrictSlash(true)
	fmt.Println("API SERVER RUNNING ON localhost:8080")
	router.HandleFunc("/fire", createEventForFire)
	router.HandleFunc("/theft", createEventForTheft)
	router.HandleFunc("/murder", createEventForMurder)
	router.HandleFunc("/events", getAllEvents).Methods("GET")
	router.HandleFunc("/events/{id}", getOneEvent).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
	go func() {
		eventMessage := <-messageChannel
		fmt.Println("eventMessage:", eventMessage)
	}()
}
