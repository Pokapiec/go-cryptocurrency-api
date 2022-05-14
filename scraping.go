package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type Crypto struct {
	Name       string
	Price      float64
	HourChange string
	DayChange  string
	WeekChange string
}

func GetAllPrices() []Crypto {
	c := colly.NewCollector(
		colly.AllowedDomains("coinmarketcap.com"),
	) // td[2]/div/a[2]

	currencies := []Crypto{}

	fmt.Println("Scraping")

	c.OnHTML("#__next > div > div > div > div > div > div > div:nth-child(3) > div > table > tbody > tr", func(e *colly.HTMLElement) {
		// e.Request.Visit(e.Attr("href"))
		// fmt.Print(e.ChildText("td:nth-child(2)"))
		// fmt.Print(" | ")
		// fmt.Print(e.ChildText("td:nth-child(5)"))
		// fmt.Print(" | ")
		// fmt.Print(e.ChildText("td:nth-child(8)"))
		// fmt.Print(" | ")
		// fmt.Print(e.ChildText("td:nth-child(9)"))
		// fmt.Print(" | ")
		// fmt.Print(e.ChildText("td:nth-child(10)"))
		// fmt.Println("")

		rawString := strings.Replace(e.ChildText("td:nth-child(5)"), "$", "", 1)

		price, _ := strconv.ParseFloat(rawString, 32)

		// if err != nil {
		// 	fmt.Println("Failed to parse!")
		// }

		singleCrypto := Crypto{
			Name:       e.ChildText("td:nth-child(2)"),
			Price:      price,
			HourChange: e.ChildText("td:nth-child(8)"),
			DayChange:  e.ChildText("td:nth-child(9)"),
			WeekChange: e.ChildText("td:nth-child(10)"),
		}

		currencies = append(currencies, singleCrypto)

	})

	c.Visit("https://coinmarketcap.com/all/views/all/")

	fmt.Println(currencies[:10])
	return currencies
}
