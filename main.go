package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	endpoint := os.Getenv("ENDPOINT")
	for {
		resp, err := http.Get(fmt.Sprintf("http://%s", endpoint))
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(resp.Status))
		}
		time.Sleep(2 * time.Second)
		resp, err = http.Get(fmt.Sprintf("http://%s/test", endpoint))
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(resp.Status)
		}
		time.Sleep(2 * time.Second)
	}
}
