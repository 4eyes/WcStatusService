package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/stianeikeland/go-rpio"
)

type Status struct {
    Occupied int `json:"occupied"`
    Time int `json:"time"`
}

func main() {
	http.HandleFunc("/", handleMainPage)

	log.Println("Starting webserver on :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal("http.ListendAndServer() failed with %s\n", err)
	}
}

func handleMainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// choose pin where switch is connected
	pin := rpio.Pin(10)

	// define mode
	pin.Input()

	// pull down the IO so the initial state is definded
	pin.PullDown()

	// get state of switch
	state := pin.Read()

	occupied := 0
	if(state == 0){
		occupied = 0
	} else {
		occupied = 1
	}
	s := Status{
		Occupied:occupied,
		Time:0,
	}

	// create json
	b, _:= json.Marshal(s)

	// deliver response
	fmt.Fprintf(w, string(b))
}
