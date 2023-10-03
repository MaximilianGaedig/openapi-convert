package main

import (
	"encoding/json"
	"fmt"
	"github.com/getkin/kin-openapi/openapi2"
	"github.com/getkin/kin-openapi/openapi2conv"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		panic("Usage: openapi-convert <input-file (json)>")
	}

	input, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	var doc openapi2.T
	if err = json.Unmarshal(input, &doc); err != nil {
		panic(err)
	}
	if doc.Swagger != "2.0" {
		panic(`doc.ExternalDocs was parsed incorrectly!`)
	}
	converted, err := openapi2conv.ToV3(&doc)
	if err != nil {
		panic(err)
	}

	outputJSON, err := json.Marshal(converted)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(outputJSON))
}
