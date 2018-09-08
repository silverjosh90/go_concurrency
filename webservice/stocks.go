package main

import (
	"runtime"
	"time"
	"fmt"
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

type quoteResponse struct {
	Status string
	Name string
	LastPrice float32
	Change float32
	ChangePercent float32
	TimeStamp string
	MSDate float32
	MarketCap int
	Volume int
	ChangeYTD float32
	ChangePercentYTD float32
	High float32
	Low float32
	Open float32
}

func main() {
	runtime.GOMAXPROCS(4)
	start := time.Now()

	symbols := []string {
		"googl",
		"msft",
		"aapl",
		"bbry",
		"hpq",
		"vz",
		"t",
		"tmus",
		"s",
	}

	mapOfStocks := make(map[string]*quoteResponse)
	numComplete := 0

	for _, symbol := range symbols {
		go func(symbol string) {
			resp, _ := http.Get("http://dev.markitondemand.com/Api/v2/Quote?symbol=" + symbol)
			defer resp.Body.Close()
		
			body, _ := ioutil.ReadAll(resp.Body)	
			quote := new(quoteResponse)
			xml.Unmarshal(body, &quote)
	
			fmt.Printf("%s: %.2f", quote.Name, quote.LastPrice)
			fmt.Println(" ")
	
	
			mapOfStocks[symbol] = quote
			numComplete++
		}(symbol)
	}

	for numComplete < len(symbols) {
		time.Sleep(10 * time.Millisecond)
	}

	// fmt.Println(mapOfStocks)

	// fmt.Printf("%s: %.2f", mapOfStocks["googl"].Name, mapOfStocks["googl"].LastPrice)

	elapsed := time.Since(start)
	fmt.Println("\nExecution time:", elapsed)

}
