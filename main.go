package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://realtime.paragaranti.com/asp/xml/icpiyasa.asp")
	if err != nil {
		log.Fatal("HATA: %s", err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	var s icpiyasa
	xml.Unmarshal(data, &s)
	fmt.Println(s.Stocks)
}

type stock struct {
	XMLName xml.Name `xml:"STOCK"`
	SYMBOL  string   `xml:"SYMBOL"`
	DESC    string   `xml:"DESC"`
	LAST    string   `xml:"LAST"`
	PERNC   string   `xml:"PERNC"`
}

type icpiyasa struct {
	XMLName xml.Name `xml:"ICPIYASA"`
	Stocks  []stock  `xml:"STOCK"`
}

func (s stock) String() string {
	return fmt.Sprintf("\t Sembol: %s - Açıklama: %s - Son Değer: %s - Değişim: %s \n", s.SYMBOL, s.DESC, s.LAST, s.PERNC)
}
