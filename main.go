package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var from string
var to string
var amount float64

type Rate struct {
	Date string `json:"date"`
	Rates  map[string]interface{} `json:"rates"`
	Base string `json:"base"`
}

type SingleForexData struct {
	Symbol  string `json:"symbol"`
	Bid float64 `json:"bid"`
	Ask float64 `json:"ask"`
	Price float64 `json:"price"`
	Timestamp int `json:"timestamp"`
}

type ForexData []SingleForexData


func getRates(from string, to string) (error, float64) {
	url := "https://forex.1forge.com/1.0.3/quotes"

	req, _ := http.NewRequest("GET", url, nil)

	q := req.URL.Query()
	q.Add("api_key", os.Getenv("FORGE_ONE_API"))
	q.Add("pairs", strings.ToUpper(from) + strings.ToUpper(to))
	req.URL.RawQuery = q.Encode()

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	data := ForexData{
		SingleForexData{},
	}

	err := json.Unmarshal(body, &data)
	return err, data[0].Price
}


func main() {
	flag.StringVar(&from, "f","usd", "which currency to convert from")
	flag.StringVar(&to, "t", "rub" , "which currency to convert to")
	flag.Float64Var(&amount, "a", 1 , "amount of currency to convert")

	flag.Parse()
	_, result := getRates(from, to)
	fmt.Println(int(result * amount))
}