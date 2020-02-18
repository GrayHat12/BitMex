package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type ResponseStruct struct {
	S string    `json:"s"`
	T []int     `json:"t"`
	C []float64 `json:"c"`
	O []float64 `json:"o"`
	H []float64 `json:"h"`
	L []float64 `json:"l"`
	V []int     `json:"v"`
}

func main() {
	now := time.Now()
	to := now.Unix()
	from := to - 86400
	url := fmt.Sprintf("https://www.bitmex.com/api/udf/history?symbol=.BXBT&resolution=1&from=%d&to=%d", from, to)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var m ResponseStruct
	json.Unmarshal(body, &m)
	file, _ := json.MarshalIndent(m, "", " ")
	_ = ioutil.WriteFile("resp.json", file, 0644)
}
