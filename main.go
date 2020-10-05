package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Planet struct {
	XMLName xml.Name `xml:"fed2-map"`
	Name    string   `xml:"title,attr"`
	Locs    []Loc    `xml:"location"`
}

type Loc struct {
	XMLName xml.Name `xml:"location"`
	Num     int      `xml:"num,attr"`
	Name    string   `xml:"name"`
	Desc    string   `xml:"desc"`
}

func main() {
	xmlFile, err := os.Open("planet.xml")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully opened xml file")

	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var pl Planet

	xml.Unmarshal(byteValue, &pl)
	fmt.Println("This is planet: " + pl.Name)
	for i := 0; i < len(pl.Locs); i++ {
		fmt.Println("Location Num: ", pl.Locs[i].Num)
		fmt.Println("Location Name: " + pl.Locs[i].Name)
		fmt.Println("Location Description: " + pl.Locs[i].Desc)
	}
}
