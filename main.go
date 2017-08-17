package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
	"regexp"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//declare data model
type Issue struct {
	gorm.Model
	Customer      string
	Description   string
	Tags          string //[]Tag - slice of tags. would need separate table + foreign key
	Status        string
	Devices       uint
	Contact_email string
	Contact_name  string
}

//regex -- url with id on end
var urlNum = regexp.MustCompile(`/api/issue/\d`) // url with digit(s) at tail

func main() {

	//open DB connection
	db, err := gorm.Open("mysql", "root:issue-tracker@/issue_tracker?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println(err)
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Issue{})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "OPTIONS" {
			if origin := r.Header.Get("Origin"); origin != "" {
				w.Header().Set("Access-Control-Allow-Origin", origin)
			}
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PATCH, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
		}

		//////////////---- LIST ISSUES ---- /////////////
		if r.Method == "GET" && r.URL.Path == "/api/issue" {
			//CORS
			if origin := r.Header.Get("Origin"); origin != "" {
				w.Header().Set("Access-Control-Allow-Origin", origin)
			}
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PATCH, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
			w.Header().Set("Access-Control-Allow-Credentials", "true")

			var issues []Issue
			// pull all issues from the database
			iss := db.Find(&issues)

			//marshal issues into json
			issj, err := json.Marshal(iss)
			if err != nil {
				log.Println(err)
				w.WriteHeader(500)
				//json error
			}

			//serve issues
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			fmt.Fprintf(w, "%s", issj)

			//////////////---- CREATE ISSUE ---- /////////////
		} else if r.Method == "POST" && r.URL.Path == "/api/issue" {
			//CORS
			if origin := r.Header.Get("Origin"); origin != "" {
				w.Header().Set("Access-Control-Allow-Origin", origin)
			}
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PATCH, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
			w.Header().Set("Access-Control-Allow-Credentials", "true")

			issue := Issue{}
			//read the request body and decode the struct into json
			err = json.NewDecoder(r.Body).Decode(&issue)
			if err != nil {
				log.Println(err)
				w.WriteHeader(500)
				//json error
			} else {
				//create issue in mysql
				db.Create(&issue)
			}
			//////////////---- READ ISSUE BY ID ---- /////////////
		} else if r.Method == "GET" && urlNum.MatchString(r.URL.Path) {
			//CORS
			if origin := r.Header.Get("Origin"); origin != "" {
				w.Header().Set("Access-Control-Allow-Origin", origin)
			}
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PATCH, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
			w.Header().Set("Access-Control-Allow-Credentials", "true")

			var issue Issue
			//get base url id and pull that issue from database
			iss := db.First(&issue, path.Base(r.URL.Path))

			//convert the data to json
			issj, err := json.Marshal(iss)
			if err != nil {
				log.Println(err)
				w.WriteHeader(500)
				//json error
			}

			//serve the issue
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			fmt.Fprintf(w, "%s", issj)

			//////////////---- UPDATE ISSUE ---- /////////////
		} else if r.Method == "PATCH" && urlNum.MatchString(r.URL.Path) {
			//CORS
			if origin := r.Header.Get("Origin"); origin != "" {
				w.Header().Set("Access-Control-Allow-Origin", origin)
			}
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PATCH, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
			w.Header().Set("Access-Control-Allow-Credentials", "true")

			//get issue modificiations from request body
			issueMod := Issue{}
			err = json.NewDecoder(r.Body).Decode(&issueMod)
			if err != nil {
				log.Println(err)
				w.WriteHeader(500)
				//json error
			}
			//get issue id from request url and select data
			issue := Issue{}
			iss := db.First(&issue, path.Base(r.URL.Path))
			//update database with changes
			iss.Updates(&issueMod)

			//////////////---- DELETE ISSUE ---- /////////////
		} else if r.Method == "DELETE" && urlNum.MatchString(r.URL.Path) {
			//CORS
			if origin := r.Header.Get("Origin"); origin != "" {
				w.Header().Set("Access-Control-Allow-Origin", origin)
			}
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PATCH, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
			w.Header().Set("Access-Control-Allow-Credentials", "true")

			var issue Issue
			//select and delete issue
			db.First(&issue, path.Base(r.URL.Path))
			db.Delete(&issue)

			w.WriteHeader(200)

			//////////////---- CATCH EVERYTHING ELSE ---- /////////////
		} else {
			log.Println("request not matched")
		}

	})

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
