package main

import (
	"fmt"

	"github.com/jda/routeros-api-go"
)

func cobalisten() {
	c, err := routeros.New("172.16.10.136:8728")
	if err != nil {
		fmt.Errorf("Error parsing address: %s\n", err)
	}

	err = c.Connect("userapi", "userapi")
	if err != nil {
		fmt.Errorf("Error connecting to device: %s\n", err)
	}

	// user active print
	res, err := c.Call("/user/active/print", nil)
	if err != nil {
		fmt.Errorf("error reiteving user active : %s\n", err)
	}

	// fmt.Println(res)
	for _, each := range res.SubPairs {
		fmt.Println(each)

	}
}
