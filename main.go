package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		println("Usage: input.csv delimiter")
		os.Exit(1)
	}

	// Open file
	inFile, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	// Parse it
	csvReader := csv.NewReader(inFile)
	outMap := make(map[string][]string)
	for {
		// csv reader fluff
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		// populate our map
		if len(rec) > 1 {
			key := strings.Split(rec[0], os.Args[2])
			if len(key) > 1 {
				outMap[strings.TrimSpace(key[1])] = []string(rec[1:])
			}
		}
	}
	outJson, err := json.MarshalIndent(outMap, "", "	")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(outJson))
}
