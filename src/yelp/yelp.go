package yelp

import (
	"log"
	"net/http"
	"encoding/json"
	"strings"
)

func HomeLink(term string, location string) map[string]interface{} {

	req, err := http.NewRequest("GET", "https://api.yelp.com/v3/businesses/search", nil)
	req.Header.Add("Authorization", "Bearer u_DoEa_3_Sn6PJHx51ZcQKaYNR4ld_8mI-MYXt-gXHHZ4DVkKOM-YVd4snyx0wxLfEyiJ7d7h89GnGVttMZfk1XdeBJv8H4esKKVa5_6RNmDQZTrwsS2iQuviBhaXXYx")
	
	q := req.URL.Query()
	q.Add("term", strings.ToLower(term))
    q.Add("location", strings.ToLower(location))
	
	req.URL.RawQuery = q.Encode()

    resp, err := http.DefaultClient.Do(req)
	
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)
	
	return result
}