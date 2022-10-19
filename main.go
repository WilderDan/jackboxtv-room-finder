package main

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"log"
	"net/http"
	"flag"
	"strings"
)

type Room struct {
	RoomId           string
	Server           string
	AppTag           string
	AppId            string
	NumAudience      int
	AudienceEnabled  bool
	JoinAs           string
	RequiresPassword bool
}

var filterStrPtr *string

// makes api calls to endpoint using every code in roomCodes[]
func findRooms(roomCodes []string) {
	// base endpoint
	url := "https://blobcast.jackboxgames.com/room/"

	for _, code := range roomCodes {

		fullUrl := url + code

		resp, err := http.Get(fullUrl)
		if err != nil {
			log.Fatalln(err)
		}

		// if 404, continue
		if resp.Status == "404 Not Found" {
			continue
		}

		// if 500, continue
		if resp.Status == "500 Internal Server Error" {
			continue
		}

		// get json response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		// if body is empty, continue
		if len(body) == 0 {
			fmt.Println("Empty response body. Your IP might be throttled/blacklisted")
			continue
		}

		// deserialize json into struct
		var room Room
		err = json.Unmarshal([]byte(body), &room)
		if err != nil {
			fmt.Println(fullUrl)
			fmt.Println(resp.Status)
			fmt.Println(body)
			log.Fatalln(err)
		}

		if room.JoinAs == "audience" {
			continue
		}

		if room.RequiresPassword {
			continue
		}

		if room.RoomId == "" {
			continue
		}

		if strings.Contains(strings.ToLower(room.AppTag), strings.ToLower(*filterStrPtr)) {
			fmt.Printf("Room Code: %s, Game: %s \n", color.GreenString(room.RoomId), color.GreenString(room.AppTag))
		}
	}
}

func main() {
	filterStrPtr = flag.String("filter", "", "display only room codes for games whose name contain the given string")
	flag.Parse()

	fmt.Println("Finding open rooms...")

	// concurrency to speed up the process of finding rooms
	go findRooms(RoomCodes[50000:100000])
	go findRooms(RoomCodes[100001:150000])
	go findRooms(RoomCodes[150001:200000])
	go findRooms(RoomCodes[200001:250000])
	go findRooms(RoomCodes[250001:300000])
	go findRooms(RoomCodes[300001:350000])
	go findRooms(RoomCodes[350001:400000])
	go findRooms(RoomCodes[400001:456976])
	findRooms(RoomCodes) // begins from 0
}
