package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type OpenWeatherResponse struct {
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Name string `json:"name"`
}

const (
	apiKeyName = "OPENWEATHER_API_KEY"
)

func getWeatherFromExternalAPI(city string) (OpenWeatherResponse, error) {
	apiKey := os.Getenv(apiKeyName)
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return OpenWeatherResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return OpenWeatherResponse{}, fmt.Errorf("external API error: %d", resp.StatusCode)
	}

	var data OpenWeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return OpenWeatherResponse{}, err
	}
	return data, nil
}

func Weather(w http.ResponseWriter, req *http.Request) {
	city := req.URL.Query().Get("city")
	weatherData, err := getWeatherFromExternalAPI(city)
	if err != nil {
		log.Printf("Failed to get weather: %v", err)
		http.Error(w, "failed to get weather", http.StatusBadGateway)
		return
	}

	response := map[string]interface{}{
		"city":                weatherData.Name,
		"temperature":         weatherData.Main.Temp,
		"feels_like":          weatherData.Main.FeelsLike,
		"humidity":            weatherData.Main.Humidity,
		"weather_description": weatherData.Weather[0].Description,
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	encoder.Encode(response)
}
