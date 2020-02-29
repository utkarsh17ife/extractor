package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	unarr "github.com/gen2brain/go-unarr"
	"github.com/go-xmlfmt/xmlfmt"
)

func extractor(zippedFileName, unzippedFolderName string) {

	var err error

	a, err := unarr.NewArchive(zippedFileName)
	if err != nil {
		panic(err)
	}
	defer a.Close()

	_, err = a.Extract(fmt.Sprintf("./%s", unzippedFolderName))
	if err != nil {
		panic(err)
	}

}

func beautifier(unformattedFileName, formattedFileName string) {

	dat, err := ioutil.ReadFile(unformattedFileName)
	if err != nil {
		panic(err)
	}

	x := xmlfmt.FormatXML(string(dat), "\t", "  ")

	err = ioutil.WriteFile(formattedFileName, []byte(x), 0644)
	if err != nil {
		panic(err)
	}

}

func main() {

	if len(os.Args) < 2 {
		panic("Please provide the file name")
	}

	zippedFileName := os.Args[1]

	timeStamp := time.Now().Unix()

	unzippedFolderName := fmt.Sprintf("unzipped_%d", timeStamp)

	formattedFileName := fmt.Sprintf("formatted_%d.xml", timeStamp)

	extractor(zippedFileName, unzippedFolderName)

	// search for unformatted file
	var files []string

	err := filepath.Walk(unzippedFolderName, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}

	// look for the .gtc file for beautification
	var gtcFilePath = ""

	for _, file := range files {

		fileNameAndExt := strings.Split(file, ".")

		if len(fileNameAndExt) > 1 {

			if fileNameAndExt[len(fileNameAndExt)-1] == "gtc" {
				gtcFilePath = file
			}

		}
		fmt.Printf("File identfied: %s", file)
	}

	if gtcFilePath != "" {
		beautifier(gtcFilePath, formattedFileName)
	}
}
