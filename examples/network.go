package examples

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func Requester() {
	// creating a client instance
	client := &http.Client{
		Timeout: time.Second * 20,
	}
	resp, err := client.Get("https://dummyjson.com/carts")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var data map[string]any
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		panic(err)
	}

	fmt.Println(data)

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	fmt.Println("response Body:", resp.Body)

}
