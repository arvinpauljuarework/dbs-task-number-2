package main

import (

	"fmt"
	"github.com/gorilla/mux"
  	"net/http"
  	"encoding/json"
	Yelp "./yelp"
	  
)

func main() {

	var port = "8000"
  	router := mux.NewRouter()
	router.HandleFunc("/search/{term}/{location}", getYelp).Methods("GET")
	fmt.Println("Listening to port:", port);
	http.ListenAndServe(":"+port, router)
}

func getYelp(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)
	fmt.Println(params)

	json.NewEncoder(w).Encode(Yelp.HomeLink(params["term"], params["location"]))
}