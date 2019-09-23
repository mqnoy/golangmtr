package main

import (
	"fmt"

	"log"

	"github.com/jda/routeros-api-go"
)

// Crud management for mikrotik hotspot

func createhotspot() {
	c, err := routeros.New("172.16.10.136:8728")
	if err != nil {
		fmt.Errorf("Error parsing address: %s\n", err)
	}

	err = c.Connect("userapi", "userapi")
	if err != nil {
		fmt.Errorf("Error connecting to device: %s\n", err)
	}

	// name="darigolang" interface=wlan_go
	var pairs []routeros.Pair
	pairs = append(pairs, routeros.Pair{Key: "name", Value: "darigolang"})
	pairs = append(pairs, routeros.Pair{Key: "interface", Value: "wlan_go"})
	// fmt.Println(pairs)
	_, err = c.Call("/ip/hotspot/add", pairs)
	if err != nil {
		fmt.Errorf("Error adding hotspot: %s\n", err)
	}

}

func readhotspot() {
	c, err := routeros.New("172.16.10.136:8728")
	if err != nil {
		fmt.Errorf("Error parsing address: %s\n", err)
	}

	err = c.Connect("userapi", "userapi")
	if err != nil {
		fmt.Errorf("Error connecting to device: %s\n", err)
	}

	res, err := c.Call("/ip/hotspot/getall", nil)
	if err != nil {
		fmt.Errorf("Error getting system resources: %s\n", err)
	}

	// outputan := res.SubPairs
	for _, each := range res.SubPairs {
		fmt.Println(each["name"])

	}
}
func updatehotspot() {
	c, err := routeros.New("172.16.10.136:8728")
	if err != nil {
		fmt.Errorf("Error parsing address: %s\n", err)
	}

	err = c.Connect("userapi", "userapi")
	if err != nil {
		fmt.Errorf("Error connecting to device: %s\n", err)
	}

	var pairs []routeros.Pair
	///ip hotspot set name="newname" [find where name ="oldname"]
	pairs = append(pairs, routeros.Pair{Key: "name", Value: "darigolang30"})
	pairs = append(pairs, routeros.Pair{Key: ".id", Value: "darigolang"})
	// ip hotspot remove [find name="server12"]
	_, err = c.Call("/ip/hotspot/set", pairs)
	if err != nil {
		log.Print("error updating hotspot: %s\n", err)
	}
}

func deletehotspot(hotspotname string) {
	c, err := routeros.New("172.16.10.136:8728")
	if err != nil {
		fmt.Errorf("Error parsing address: %s\n", err)
	}

	err = c.Connect("userapi", "userapi")
	if err != nil {
		fmt.Errorf("Error connecting to device: %s\n", err)
	}

	var pairs []routeros.Pair
	// var q routeros.Query

	pairs = append(pairs, routeros.Pair{Key: ".id", Value: hotspotname})
	// ip hotspot remove [find name="server12"]
	_, err = c.Call("/ip/hotspot/remove", pairs)
	if err != nil {
		fmt.Errorf("error removing hotspot: %s\n", err)
	}
}
