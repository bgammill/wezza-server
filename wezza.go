package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strings"
	"io"
	"log"
	"fmt"
)

const JSON = "application/json"
const API_KEY = "49715ecba29d9b1fe6c2ccd59ea03ebb"

type User []struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
	Email string `json:"email"`
	Address struct {
		   Street string `json:"street"`
		   Suite string `json:"suite"`
		   City string `json:"city"`
		   Zipcode string `json:"zipcode"`
		   Geo struct {
				  Lat string `json:"lat"`
				  Lng string `json:"lng"`
			  } `json:"geo"`
	   } `json:"address"`
	Phone string `json:"phone"`
	Website string `json:"website"`
	Company struct {
		   Name string `json:"name"`
		   CatchPhrase string `json:"catchPhrase"`
		   Bs string `json:"bs"`
	   } `json:"company"`
}

type WeatherResult struct {
	Coord struct {
		      Lon float64 `json:"lon"`
		      Lat float64 `json:"lat"`
	      } `json:"coord"`
	Weather []struct {
		ID int `json:"id"`
		Main string `json:"main"`
		Description string `json:"description"`
		Icon string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		      Temp float64 `json:"temp"`
		      Pressure int `json:"pressure"`
		      Humidity int `json:"humidity"`
		      TempMin float64 `json:"temp_min"`
		      TempMax float64 `json:"temp_max"`
	      } `json:"main"`
	Wind struct {
		      Speed float64 `json:"speed"`
		      Deg float64 `json:"deg"`
	      } `json:"wind"`
	Clouds struct {
		      All int `json:"all"`
	      } `json:"clouds"`
	Dt int `json:"dt"`
	Sys struct {
		      Type int `json:"type"`
		      ID int `json:"id"`
		      Message float64 `json:"message"`
		      Country string `json:"country"`
		      Sunrise int `json:"sunrise"`
		      Sunset int `json:"sunset"`
	      } `json:"sys"`
	ID int `json:"id"`
	Name string `json:"name"`
	Cod int `json:"cod"`
}

func main() {
	// declare router and set routes
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/users", MehHandler)
	r.HandleFunc("/weather/{location}", WeatherHandler)

	http.ListenAndServe(":3000", r)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<a href='/users'>/users</a>"))
}

func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	var weathers []WeatherResult
	v := mux.Vars(r)
	l := v["location"]
	url := "http://api.openweathermap.org/data/2.5/weather?q=" + l + "&APPID=" + API_KEY

	w.Header().Set("Content-Type", JSON)

	// get url
	resp, err := http.Get(url)

	// check for error
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)

	dec := json.NewDecoder(strings.NewReader(string(contents)))
	for {
		var m WeatherResult
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			w.Write([]byte("not found."))
			log.Print(err)
			break
		}

		weathers = append(weathers, m)
	}

	b, err := json.Marshal(weathers)
	if err != nil {
		fmt.Println("error:", err)
	}

	w.Write([]byte(b))
}

func MehHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSON)

	var users []User

	// get url
	resp, err := http.Get("http://jsonplaceholder.typicode.com/users")

	// check for error
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)

	dec := json.NewDecoder(strings.NewReader(string(contents)))
	for {
		var m User
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		users = append(users, m)
	}

	b, err := json.Marshal(users)
	if err != nil {
		fmt.Println("error:", err)
	}

	w.Write([]byte(b))
}