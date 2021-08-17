package main

import (
	"fmt"
	"log"

	"github.com/obity/properties"
)

func main() {
	// store to file
	p := properties.NewProperties()
	p.SetProperty("HttpPort", "8081")
	p.SetProperty("MongoServer", "mongodb://10.11.1.5,10.11.1.6,10.11.1.7/?replicaSet=mytest")
	p.SetPropertySlice("LogLevel", "Debug", "Info", "Warn")
	err := p.StoreToFile("config.properties")
	if err != nil {
		log.Println(err)
		return
	}
	// load from file
	file := "./config.properties"
	p2 := properties.NewProperties()
	err = p2.LoadFromFile(file)
	if err != nil {
		log.Println(err)
		return
	}
	httpPort, isExist := p2.Property("HttpPort")
	if !isExist {
		log.Println("HttpPort not exist")
	}
	fmt.Println(httpPort)
	loglevel, isExist := p2.PropertySlice("LogLevel")
	if !isExist {
		log.Println("LogLevel not exist")
	}
	fmt.Println(loglevel)
}
