package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
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

func getRates(from string, to string) (error, float64) {
	url := "https://exchangeratesapi.io/api/latest"

	req, _ := http.NewRequest("GET", url, nil)

	q := req.URL.Query()
	q.Add("base", strings.ToUpper(from))
	q.Add("symbols", strings.ToUpper(to))
	req.URL.RawQuery = q.Encode()

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	data := Rate{}
	err := json.Unmarshal(body, &data)

	return err, data.Rates[strings.ToUpper(to)].(float64)
}


func main() {
	flag.StringVar(&from, "f","usd", "which currency to convert from")
	flag.StringVar(&to, "t", "rub" , "which currency to convert to")
	flag.Float64Var(&amount, "a", 1 , "amount of currency to convert")

	flag.Parse()

	_, result := getRates(from, to)
	fmt.Println(int(result * amount))
}