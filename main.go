package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/ptrv/go-gpx"
)

func createNewGpxFile(folderPath string, gpx gpx.Gpx) error {
	fileName := folderPath + "/" + gpx.Tracks[0].Name + ".gpx"
	return ioutil.WriteFile(fileName, gpx.ToXML(), os.ModePerm)
}

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		panic("Provide a file path")
	}

	pathToGpx := args[0]

	if _, err := os.Stat(pathToGpx); err != nil {
		panic("Provided file path is not correct.")
	}
	folderPath := strings.Split(pathToGpx, ".gpx")[0]

	gpxFile, err := gpx.ParseFile(pathToGpx)
	if err != nil {
		fmt.Println("Couldn't parse gpx file located at $v", pathToGpx)
		panic(err)
	}

	if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
		fmt.Println("Couldn't create folder at location $v", folderPath)
		panic(err)
	}

	createdFiles := 0
	for _, track := range gpxFile.Tracks {
		newGpx := *gpxFile.Clone()
		newGpx.Tracks = []gpx.Trk{track}

		if err := createNewGpxFile(folderPath, newGpx); err != nil {
			log.Printf("Failed to create file for %v", track.Name)
			continue
		}

		createdFiles++
	}

	log.Printf("Successfuly created %d file at %v", createdFiles, folderPath)
}
