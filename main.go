package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/posener/complete"
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
	flag.StringVar(&from, "from","usd", "which currency to convert from")
	flag.StringVar(&to, "to", "rub" , "which currency to convert to")
	flag.Float64Var(&amount, "amount", 1 , "amount of currency to convert")

	cmp := complete.New(
		"econv",
		complete.Command{Flags: complete.Flags{
			"-from": complete.PredictAnything,
			"-to": complete.PredictAnything,
			"-amount": complete.PredictAnything,
		}},
	)

	cmp.AddFlags(nil)

	flag.Parse()

	if cmp.Complete() {
		return
	}

	_, result := getRates(from, to)
	fmt.Println(int(result * amount))
}