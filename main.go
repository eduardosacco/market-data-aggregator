package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type trade struct {
	Id     int
	Market int
	Price  float64
	Volume float64
	Is_buy bool
}

type market struct {
	Market_id        int
	Total_volume     float64
	Mean_price       float64
	Sum_price_volume float64
	VWAP             float64
	Total_trades     int
	Total_buys       int
	Percentage_buy   float64
}

const max_markets = 12000
const print_to_file = true

// main function
func main() {
	var markets []market = make([]market, max_markets)
	scanner := bufio.NewScanner(os.Stdin)
	done := make(chan bool)
	c := make(chan []byte)

	fmt.Println("Starting...")
	beginTime := time.Now()

	// STDIN Scanner goroutine
	go func() {
		for scanner.Scan() {
			buffer := new(bytes.Buffer)
			//buffering is needed since
			//The underlying array may point to data that will be overwritten by a subsequent call to Scan
			buffer.Write(scanner.Bytes())
			c <- buffer.Bytes()
		}
	}()

	// Message Processor goroutine
	go func() {
		for message := range c {
			if !processMessage(message, markets) {
				if string(message) == "END" {
					close(done)
				}
			}
		}
	}()

	// Wait for goroutines to finish
	<-done

	// Print stats
	duration := time.Since(beginTime)
	fmt.Printf("Duration of aggregate operation: %s \n", duration.String())
	markets = trimEmptyMarkets(markets)
	tradeCount, marketCount := calculateStats(markets)
	fmt.Printf("Total trades: %d \n", tradeCount)
	fmt.Printf("Total markets: %d \n", marketCount)

	if print_to_file {
		printToFile(markets)
	}
}

func processMessage(message []byte, markets []market) bool {
	if json.Valid(message) {
		var tx *trade
		err := json.Unmarshal(message, &tx)

		if err == nil {
			processTrade(*tx, markets)
		}

		return true
	}
	return false
}

func processTrade(tx trade, markets []market) {
	var m *market = &markets[tx.Market-1]

	if m.Market_id == 0 {
		m.Market_id = tx.Market
	}

	m.Total_trades++
	if tx.Is_buy {
		m.Total_buys++
	}
	m.Mean_price = (m.Mean_price + tx.Price) / 2
	m.Total_volume += tx.Volume
	m.Sum_price_volume += tx.Price * tx.Volume
	m.VWAP = m.Sum_price_volume / m.Total_volume
	m.Percentage_buy = float64(m.Total_buys) / float64(m.Total_trades) * 100
}

func printToFile(markets []market) {
	fmt.Println("Writing to file...")
	file, _ := json.MarshalIndent(markets, "", " ")
	_ = os.WriteFile("markets.json", file, 0644)
}

func trimEmptyMarkets(markets []market) []market {
	for i, market := range markets {
		if market.Market_id == 0 {
			return markets[:i]
		}
	}

	return markets
}

func calculateStats(markets []market) (int, int) {
	tradeCount := 0
	marketCount := 0
	for _, market := range markets {
		tradeCount += market.Total_trades
		marketCount++
	}

	return tradeCount, marketCount
}
