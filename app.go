package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/joho/godotenv"
)

var dbPath string

func init() {

	e := godotenv.Load()

	if e != nil{

		fmt.Println("Error loading .env file....")
		os.Exit(1)
	}


	dbPath = os.Getenv("DB_PATH")

	db, err := gorm.Open("postgres", dbPath)

	if err != nil {

		fmt.Println(err)
		os.Exit(1)
	}

	db.CreateTable(&User{})

}

func main() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/index", index).Methods("GET")
	router.HandleFunc("/users", GetAllUsers).Methods("GET")
	router.HandleFunc("/user/{id:[a-zA-Z0-9_]+}", GetUserById).Methods("GET")

	router.HandleFunc("/user", CreateUser).Methods("POST")

	fmt.Println("server starting....")
	log.Fatal(http.ListenAndServe(":8000", router))

}

func index(w http.ResponseWriter, r *http.Request) {

	response := make(map[string]string)
	response["status"] = "200"
	response["info"] = "Home page"

	w.Header().Set("content-type", "application/json") // important
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}
