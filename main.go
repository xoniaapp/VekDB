package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
)

var Data = []Value{}

type Value struct {
	Key   int    `json:"key"`
	Value string `json:"value"`
}

func main() {
	log.Println("vekDB Starting...")

	http.HandleFunc("/create", postValue)
	http.HandleFunc("/get", getValue)

	log.Println("vekDB Started!")
	http.ListenAndServe(":3000", nil)
}

func getValue(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(Data)
}

func postValue(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Panic(err)
	}

	var Value Value
	json.Unmarshal(body, &Value)

	Value.Key = rand.Intn(100)

	Data = append(Data, Value)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("key created!")
}
