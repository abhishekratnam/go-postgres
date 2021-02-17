package handlers

import (
	"encoding/json"
	"log"
	"microservice/data"
	"microservice/models"
	"net/http"
)

//json.NewEncoder(w)
//	return e.Encode(i)
type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func (p *Products) AddProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handler POST Product")
	var user models.Product
	// prod := &data.Product{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}
	// insertID := data.Insertuser(user)
	data.Insertuser(user)
	// return insertID
	// return string("Product added successfully", InsertID)
}
