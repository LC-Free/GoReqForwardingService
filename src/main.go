package main

import (
	"encoding/json"
	"log"
	"net/http"
	"io/ioutil"
	"os"
	"gopkg.in/yaml.v2"
	"fmt"
)

/*type tickerDataRaw struct {
	Dataset struct {
		ID                  int             `json:"id"`
		DatasetCode         string          `json:"dataset_code"`
		DatabaseCode        string          `json:"database_code"`
		Name                string          `json:"name"`
		Description         string          `json:"description"`
		RefreshedAt         string          `json:"refreshed_at"`
		NewestAvailableDate string          `json:"newest_available_date"`
		OldestAvailableDate string          `json:"oldest_available_date"`
		ColumnNames         []string        `json:"column_names"`
		Frequency           string          `json:"frequency"`
		Premium             bool            `json:"premium"`
		StartDate           string          `json:"start_date"`
		EndDate             string          `json:"end_date"`
		Data                [][]interface{} `json:"data"`
		DatabaseID          int             `json:"database_id"`
	} `json:"dataset"`
}

type tickerDataParsed struct {
	StartDate   string      `json:"start_date"`
	EndDate     string      `json:"end_date"`
	RefreshedAt string      `json:"refreshed_at"`
	Data        [][]float64 `json:"data"`
}*/

type Config struct {
	BaseUrl string `yaml:"baseUrl"`
}

var conf Config

func main() {
	readConfig(&conf)
	http.HandleFunc("/", allHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func readConfig(conf *Config) {
	f, err := os.Open("config.yml")
	if err != nil {
		log.Fatalf("yaml config read error: %v\n", err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(conf)
	if err != nil {
		log.Fatalf("yaml decode error: %v\n", err)
	}
}

func allHandler(w http.ResponseWriter, r *http.Request) {
	// should probably handle if r.URL.Path is "/"
	var url = conf.BaseUrl + r.URL.Path
	fmt.Println(url)

	res, err := http.Get(url)
	if err != nil {
		log.Fatalf("fetch err: %v\n", err)
	}
	defer res.Body.Close()
	
	bodyBytes, _ := ioutil.ReadAll(res.Body)
	bodyString := string(bodyBytes)
	// still need to decide how to write out the res, but tired now so stopping here
	err = json.NewEncoder(w).Encode(bodyString)
	if err != nil {
		log.Fatalf("json encode error: %v\n", err)
	}

}
