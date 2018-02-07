// go-wfs project main.go
package main

import (
	//	"encoding/json"
	//	"fmt"

	"github.com/go-spatial/go-wfs/gpkg"
	"github.com/go-spatial/go-wfs/server"
)

func main() {
	filepath := "sandbox/athens-osm-20170921.gpkg"
	athens := gpkg.OpenGPKG(filepath)
	defer gpkg.CloseGPKG(filepath)

	//	for _, tname := range athens.FeatureTables() {
	//		fmt.Println(tname)
	//	}
	server.StartServer(":9000", &athens)
}