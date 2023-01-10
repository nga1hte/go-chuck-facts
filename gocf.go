package main

import (
	"flag"
	"fmt"
	"go-chuck-facts/client"
	"go-chuck-facts/model"
	"log"
	"strings"
	"time"
)

func main() {
	categories := model.Categories
	cat := categories.MapValues()

	factAmt := flag.Int("n", 1, "Number of facts to fetch")
	category := flag.String("c", "", "categories: "+categories.PrintCategories())
	flag.Parse()

	if _, ok := cat[*category]; !ok {
		*category = ""
	}

	cnClient := client.NewCNClient()
	cnClient.SetTimeout(time.Duration(30) * time.Second)

	facts := []string{}
	for i := 0; i < *factAmt; i++ {
		fact, err := cnClient.Fetch(*category)
		if err != nil {
			log.Println(err)
		}
		facts = append(facts, "> "+fact.Value)
	}

	fmt.Printf("%s\n", strings.Join(facts, "\n"))
}
