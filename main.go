package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/brian-armstrong/gpio"
)

// Connect a LED to GPIO11 on the Omega2
const ledPin = 11

var (
	led      gpio.Pin
	ledState bool
)

func sendJsonState(res http.ResponseWriter, message, state string) {
	body := struct {
		Message string `json:"message"`
		State   string `json:"state"`
	}{
		Message: message,
		State:   state,
	}
	json.NewEncoder(res).Encode(body)
}

func isJsonReq(req *http.Request) bool {
	return strings.Contains(req.Header.Get("content-type"), ("json"))
}

// Endpoint that turn on the LED
func handleOnRequest(res http.ResponseWriter, req *http.Request) {
	led.High()
	ledState = true
	if isJsonReq(req) {
		sendJsonState(res, "ok", "On")
	} else {
		http.Redirect(res, req, "/", 303)
	}
}

// Endpoint that turn off the LED
func handleOffRequest(res http.ResponseWriter, req *http.Request) {
	led.Low()
	ledState = false
	if isJsonReq(req) {
		sendJsonState(res, "ok", "Off")
	} else {
		http.Redirect(res, req, "/", 303)
	}
}

var homeTemplate = `
<html>
	<head>
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<link rel="stylesheet" href="https://fonts.googleapis.com/icon?family=Material+Icons">
		<link rel="stylesheet" href="https://code.getmdl.io/1.3.0/material.indigo-pink.min.css">		
		<title>Onion + Golang</title>
	</head>
	<body class="container">
		<center>		
			<h3> Onion + Golang = ❤️ </h3>		
			The LED is {{ .LedState }}
			<br/>
			<button
				class="mdl-button mdl-js-button mdl-button--raised mdl-js-ripple-effect mdl-button--accent"
				onClick="javascript:location.href = '/led/{{ .Pin }}/{{ .NextState }}';">			
				Turn {{ .NextState }}
			</button>
		</center>
	</body>
</html>
`

// Endpoint that handle the home page render
func handleHomeRequest(res http.ResponseWriter, req *http.Request) {
	t := template.New("main")
	t, _ = t.Parse(homeTemplate)

	nextState := "on"
	currentLedState := "off"
	if ledState {
		nextState = "off"
		currentLedState = "on"
	}

	t.Execute(res, struct {
		LedState  string
		NextState string
		Pin       int
	}{
		LedState:  currentLedState,
		NextState: nextState,
		Pin:       ledPin,
	})
}

func main() {
	led = gpio.NewOutput(ledPin, false) // Set GPIO as output
	defer led.Close()                   // Close when the program ends

	onURL := fmt.Sprintf("/led/%d/on", ledPin)
	offURL := fmt.Sprintf("/led/%d/off", ledPin)

	// Register functions to handle each url request
	http.HandleFunc(onURL, handleOnRequest)
	http.HandleFunc(offURL, handleOffRequest)
	http.HandleFunc("/", handleHomeRequest)

	fmt.Println("Server running at port 9090...")
	http.ListenAndServe(":9090", nil) // This blocks execution
}
