package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// cat locales.json | jq keys > localencodes.out
	data, err := ioutil.ReadFile("./localecodes.out")
	if err != nil {
		panic(err)
	}
	var codes []string
	json.Unmarshal([]byte(data), &codes)
	codes = append(codes, "tlh_aa") // https://minecraft-archive.fandom.com/wiki/Languages
	for _, langCode := range codes {

		resp, err := http.Get(fmt.Sprintf("http://localhost:3005/assets/i18n/%s.json", langCode))
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("Response: %v\tLanguage: %v\t%v\n", resp.StatusCode, langCode, resp.Request.URL)
		defer resp.Body.Close()
	}
}
