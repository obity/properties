package properties_test

import (
	"fmt"
	"log"

	"github.com/obity/properties"
)

func Example() {
	p := properties.NewProperties()
	p.SetProperty("HttpPort", "8081")
	p.SetProperty("MongoServer", "mongodb://10.11.1.5,10.11.1.6,10.11.1.7/?replicaSet=mytest")
	p.SetPropertySlice("LogLevel", "Debug", "Info", "Warn")
	err := p.StoreToFile("config.properties")
	if err != nil {
		log.Println(err)
		return
	}
	// Output:
	//// config.properties
	// HttpPort = 8081
	// MongoServer = mongodb://10.11.1.5,10.11.1.6,10.11.1.7/?replicaSet=mytest
	// LogLevel = Debug,Info,Warn
}

func ExampleProperties_LoadFromFile() {
	file := "./config.properties"
	p := properties.NewProperties()
	err := p.LoadFromFile(file)
	if err != nil {
		log.Println(err)
		return
	}

	httpPort, isExist := p.Property("HttpPort")
	if !isExist {
		log.Println("HttpPort not exist")
	}
	fmt.Println(httpPort)

	loglevel, isExist := p.PropertySlice("LogLevel")
	if !isExist {
		log.Println("LogLevel not exist")
	}
	fmt.Println(loglevel)

	// Output:
	// 8081
	// [Debug Info Warn]
}

func ExampleProperties_PropertySlice() {
	p := properties.NewProperties()
	err := p.StoreToFile("webconfig.properties")
	if err != nil {
		log.Println(err)
		return
	}

	loglevel, isExist := p.PropertySlice("LogLevel")
	if !isExist {
		log.Println("LogLevel not exist")
	}
	fmt.Println(loglevel)
	// Output:
	// [Debug Info Warn]
}

func ExampleProperties_SetProperty() {
	p := properties.NewProperties()
	p.SetProperty("HttpPort", "8081")

}

func ExampleProperties_SetPropertySlice() {
	p := properties.NewProperties()
	p.SetPropertySlice("LogLevel", "Debug", "Info", "Warn")
}

func ExampleProperties_StoreToFile() {
	p := properties.NewProperties()
	p.SetProperty("HttpPort", "8081")
	p.SetProperty("MongoServer", "mongodb://10.11.1.5,10.11.1.6,10.11.1.7/?replicaSet=mytest")
	p.SetPropertySlice("LogLevel", "Debug", "Info", "Warn")
	err := p.StoreToFile("webconfig.properties")
	if err != nil {
		log.Println(err)
		return
	}
	// Output:
	//// webconfig.properties
	// HttpPort = 8081
	// MongoServer = mongodb://10.11.1.5,10.11.1.6,10.11.1.7/?replicaSet=mytest
	// LogLevel = Debug,Info,Warn
}
