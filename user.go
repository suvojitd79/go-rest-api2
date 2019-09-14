package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	gorm.Model
	Email string `gorm:"column:username"`
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {

	db, _ := gorm.Open("postgres", dbPath)

	w.Header().Set("content-type", "application/json")

	userData := []User{}

	if err := db.Find(&userData).Error; err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)

	} else {

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&userData)
	}

	defer db.Close()

}

func GetUserById(w http.ResponseWriter, r *http.Request) {

	UrlData := mux.Vars(r)

	var user User

	db, _ := gorm.Open("postgres", dbPath)

	w.Header().Set("content-type", "application/json")

	if e := db.Where("id = ?", UrlData["id"]).Find(&user).Error; e != nil {

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(e)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&user)

	}

	defer db.Close()

}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	db, _ := gorm.Open("postgres", dbPath)

	w.Header().Set("content-type", "application/json")

	var data User

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	if e := db.Create(&data).Error; e != nil {

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(e)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)

	defer db.Close()

}
