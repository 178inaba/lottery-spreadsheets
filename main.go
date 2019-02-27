package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strconv"
	"time"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	credFilePath := flag.String("c", "", "credentials")
	readRange := flag.String("r", "A2:B", "range")
	flag.Parse()
	spreadsheetID := flag.Arg(0)

	b, err := ioutil.ReadFile(*credFilePath)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v.", err)
	}

	conf, err := google.JWTConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets.readonly")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v.", err)
	}

	ctx := context.Background()
	srv, err := sheets.New(conf.Client(ctx))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v.", err)
	}

	resp, err := srv.Spreadsheets.Values.Get(spreadsheetID, *readRange).Context(ctx).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v.", err)
	}

	var items []lotteryItem
	if len(resp.Values) == 0 {
		log.Fatal("No data found.")
	}

	for _, row := range resp.Values {
		name := row[0].(string)
		weight, err := strconv.Atoi(row[1].(string))
		if err != nil {
			log.Fatal(err)
		}

		items = append(items, lotteryItem{name: name, weight: weight})
	}

	fmt.Println(lottery(items))
}

type lotteryItem struct {
	name   string
	weight int
}

func lottery(ls []lotteryItem) string {
	var total int
	for _, li := range ls {
		total += li.weight
	}

	rnd := rand.Intn(total)
	for _, li := range ls {
		if li.weight > rnd {
			return li.name
		}

		rnd -= li.weight
	}

	return ""
}
