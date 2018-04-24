package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"html/template"

)



func main() {
	http.HandleFunc("/1", path1)
	http.HandleFunc("/2", path2)
	http.HandleFunc("/3", path3)
	http.HandleFunc("/4", path4)
	http.HandleFunc("/5", path5)
	http.ListenAndServe(":8080", nil)
}



type St1 struct {
	Dato                string `json:"Dato"`
	Klokkeslett         string `json:"Klokkeslett"`
	Sted                string `json:"Sted"`
	Latitude            string `json:"Latitude"`
	Longitude           string `json:"Longitude"`
	Antall_Ledige_Plasser string `json:"Antall_ledige_plasser"`
}


type St2 struct {
	Entries []struct {
		East      string `json:"east"`
		Zone      string `json:"zone"`
		North     string `json:"north"`
		Latitude  string `json:"latitude"`
		Longitude string `json:"longitude"`
	} `json:"entries"`
}


type St3 struct {
	Entries []struct {
		Latitude    string `json:"latitude"`
		Name        string `json:"name"`
		Adressenavn string `json:"adressenavn"`
		Longitude   string `json:"longitude"`
	} `json:"entries"`
}


type St4 struct {
	Entries []struct {
		Name   string `json:"Navn"`
		Nummer string `json:"Nummer"`
	} `json:"entries"`
}

type St5 struct {
	Countries []struct {
		Name string `json:"name"`
		Code string `json:"code"`
	} `json:"countries"`
}


func path1(w http.ResponseWriter, r *http.Request) {

	page, err := http.Get("https://opencom.no/dataset/36ceda99-bbc3-4909-bc52-b05a6d634b3f/resource/d1bdc6eb-9b49-4f24-89c2-ab9f5ce2acce/download/parking.json")
	if err != nil {
		log.Fatal(err)
	}
	jSonInfo, err := ioutil.ReadAll(page.Body)
	page.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var steder []St1
	er := json.Unmarshal(jSonInfo, &steder)
	if err != nil {
		fmt.Println("error:", er)
	}

	tmpl, err := template.ParseFiles("st1.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, steder); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
//slutt path1

func path2(w http.ResponseWriter, r *http.Request) {

	page, err := http.Get("https://hotell.difi.no/api/json/stavanger/lekeplasser?")
	if err != nil {
		log.Fatal(err)
	}
	jSonInfo, err := ioutil.ReadAll(page.Body)
	page.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var lekeplasser St2
	er := json.Unmarshal(jSonInfo, &lekeplasser)
	if err != nil {
		fmt.Println("error:", er)
	}
	tmpl, err := template.ParseFiles("st2.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, lekeplasser); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func path3(w http.ResponseWriter, r *http.Request) {

	page, err := http.Get("https://hotell.difi.no/api/json/stavanger/utsiktspunkt?")
	if err != nil {
		log.Fatal(err)
	}
	jSonInfo, err := ioutil.ReadAll(page.Body)
	page.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var punkt St3

	er := json.Unmarshal(jSonInfo, &punkt)
	if err != nil {
		fmt.Println("error:", er)
	}
	tmpl, err := template.ParseFiles("st3.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, punkt); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func path4(w http.ResponseWriter, r *http.Request) {

	page, err := http.Get("https://hotell.difi.no/api/json/difi/geo/fylke")
	if err != nil {
		log.Fatal(err)
	}
	jSonInfo, err := ioutil.ReadAll(page.Body)
	page.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var fylker St4

	er := json.Unmarshal(jSonInfo, &fylker)
	if err != nil {
		fmt.Println("error:", er)
	}
	tmpl, err := template.ParseFiles("st4.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, fylker); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func path5(w http.ResponseWriter, r *http.Request) {
	page, err := http.Get("http://api.nobelprize.org/v1/country.json")
	if err != nil {
		log.Fatal(err)
	}
	jSonInfo, err := ioutil.ReadAll(page.Body)
	page.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var nobel St5

	er := json.Unmarshal(jSonInfo, &nobel)
	if err != nil {
		fmt.Println("error:", er)
	}
	tmpl, err := template.ParseFiles("st5.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nobel); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
