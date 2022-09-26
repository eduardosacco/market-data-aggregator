# MARKET-DATA-AGGREGATOR

## About
Found an online challenge for [Messari](https://engineering.messari.io/) a Web3 crypto research & data company
and thought it would be a good idea to attempt doing it to practice Go programming language which I am currently learning.

The purpose of the challenge is to write a tool that parses trade data as it comes in and compute various aggregate metrics from the provided data, completing a set of ten million trades in as little time as possible.

A binary that writes ten million trade objects as JSON to stdout (stdoutinator) is provided. The tool should accept input from stdin so output from the provided binary can be piped into it using a terminal window.

This is actually my second Go script so the language best practices and conventions may not be top notch here.

## Resources I used on my investigation
* [Challenge](https://engineering.messari.io/blog/messari-market-data-coding-challenge)
* [Slices](https://go.dev/blog/slices)
* [Variable Weighted Average Price](https://www.investopedia.com/terms/v/vwap.asp)
* [bufio Reader vs Scanner benchmarking](https://www.anycodings.com/1questions/3123588/bufioreader-and-bufioscanner-functionality-and-performance)
* [Reading line by lines in go](https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go)
* [Go perfbook](https://github.com/dgryski/go-perfbook)

## How to use
Make sure you have Go installed, more info [here](https://go.dev/doc/install).

From a terminal where the main.go file is located just run `go build` to build the market-data-aggregator executable

To pipe stdoutinator's stdout to market-data-aggregator's stdin, in MacOS use the following command. Note that the terminal must be open in the location that both executables are located.

`exec ./stdoutinator | ./market-data-aggregator`

### Output
You should see a terminal output like the one below.

```
➜  market-data-aggregator git:(master) ✗ exec ./stdoutinator | ./market-data-aggregator
Starting...
Duration of aggregate operation: 20.201481417s 
Total trades: 10000001 
Total markets: 8831 
Writing to file...
```

A markets.json file will be generated containing all aggregated markets data from all trades.
