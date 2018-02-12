// +build ignore
package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	body, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("error reading stdin: %v", err)
	}
	json.NewEncoder(os.Stdout).Encode(struct {
		Body string
	}{
		Body: string(body),
	})
}
