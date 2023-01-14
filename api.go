package main

import (
	"fmt"
	"log"
	"net/http"
)

var alarm_state bool = false

func alarmSetHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/set_alarm" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	if alarm_state {
		alarm_state = false
		fmt.Fprintln(w, "Setting Alarm to off")
	} else {
		alarm_state = true
		fmt.Fprintln(w, "Setting Alarm to on")
	}
}

func alarmStateHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/alarm_state" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	if alarm_state {
		fmt.Fprintf(w, "ON")
	} else {
		fmt.Fprintf(w, "OFF")
	}
}

func main() {
	http.HandleFunc("/set_alarm", alarmSetHandler)
	http.HandleFunc("/alarm_state", alarmStateHandler)

	fmt.Printf("Starting server at port 8111\n")
	if err := http.ListenAndServe(":8111", nil); err != nil {
		log.Fatal(err)
	}
}
