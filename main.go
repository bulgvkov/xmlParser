package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ValCurs struct {
	Date   string `xml:"Date"`
	Valute []struct {
		Name  string `xml:"Name"`
		Value string `xml:"Value"`
	} `xml:"Valute"`
}

func get() ValCurs {
	var result ValCurs
	xmlFile, err := os.ReadFile("XML_daily_eng.xml")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully opened XML_daily_eng.xml")
	xml.Unmarshal(xmlFile, &result)
	return result
}

func parseFloat(ctringa string) (float64, error) {
	var flo float64
	ctringa = strings.Replace(ctringa, ",", ".", -1)
	ctringa = strings.TrimSuffix(ctringa, "\n")
	flo, err := strconv.ParseFloat(ctringa, 64)
	if err != nil {
		panic(err)
		return flo, err
	}
	return flo, nil
}

func main() {

	var avgCurrency, counter float64
	var min, max float64
	var minDate, maxDate, valuteNameMin, valuteNameMax string
	var err error
	var result ValCurs

	result = get()
	min, err = parseFloat(result.Valute[0].Value)
	if err != nil {
		panic(err)
		return
	}
	max, err = parseFloat(result.Valute[0].Value)
	if err != nil {
		panic(err)
		return
	}
	for j := 0; j < len(result.Valute); j++ {
		currentElement, err := parseFloat(result.Valute[j].Value)
		if err != nil {
			panic(err)
			return
		}
		if min > currentElement {
			min = currentElement
			valuteNameMin = result.Valute[j].Name
			minDate = result.Date
		}
		if max < currentElement {
			max = currentElement
			valuteNameMax = result.Valute[j].Name
			maxDate = result.Date
		}
		avgCurrency += currentElement
		counter++
	}

	avgCurrency /= counter

	fmt.Println(result.Date)
	fmt.Printf("Максимальное значение %f, название %s, дата %s", max, valuteNameMax, maxDate)
	fmt.Printf("Минимальное значение %f, название %s, дата %s", min, valuteNameMin, minDate)
	fmt.Printf("Среднее значение курса рубля за весь период по всем валютам: %f", avgCurrency)
}
