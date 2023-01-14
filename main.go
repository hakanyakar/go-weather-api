package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type WeatherResponse struct {
	Name string `json:"name"`
	Main Main   `json:"main"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
}

func main() {
	response, err := http.Get(`https://api.openweathermap.org/data/2.5/weather?lat=36.79&lon=34.62&units=metric&appid=f1d16224e2508911600a52ccd944a4ca`)

	if err != nil {
		log.Fatal(err)
	}

	bytes, err_read := ioutil.ReadAll(response.Body)

	defer func() {
		e := response.Body.Close()
		if e != nil {
			log.Fatal(e)
		}
	}()

	if err_read != nil {
		log.Fatal(err_read)
	}

	var weather_reponse WeatherResponse
	err_unmarshall := json.Unmarshal(bytes, &weather_reponse)
	if err_unmarshall != nil {
		log.Fatal(err_unmarshall)
	}

	log.Printf("%+v", weather_reponse)
}
