package main

import (
	"fmt"
	route2 "github.com/frsbrito/Imersao-Full-Stack/application/route"
)

func main() {
	route := route2.Route{
		ID:			"1",
		ClientID:	"1",
		// Positions:	nil,
	}
	route.LoadPositions()
	stringjson, _ := route.ExportJsonPositions()
	fmt.Println(stringjson[0])
}