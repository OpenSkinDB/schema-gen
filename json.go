package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func ExportToJsonFile(v any, fname string) {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetEscapeHTML(false) // Disable HTML escaping
	encoder.SetIndent("", "  ")

	// Marshal the data to JSON format
	jsonData, err := json.MarshalIndent(v, "", "  ")

	if err != nil {
		fmt.Println("Error marshalling music kits to JSON:", err)
		return
	}

	if fname == "" {
		fname = "music_kits"
	}

	// dump to file
	var fileName = fmt.Sprintf(`exported/%s.json`, fname)
	err = os.WriteFile(fileName, jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing data to file:", err)
		return
	}
}
