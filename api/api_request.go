package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)


func Api_request() {
	client := &http.Client{}
	req, err := http.NewRequest(
		"GET", "https://api.waqi.info/search/", nil,
	)
	// добавляем заголовки
	req.Header.Add("Accept", "application/json")   // добавляем заголовок Accept
	req.Header.Add("User-Agent", "MSIE/15.0")   // добавляем заголовок User-Agent


	q := req.URL.Query()
	q.Add("token", "")
	q.Add("keyword", "Moscow")
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)

	var result map[string]interface{}


	//type result struct {  примерная структура json ответа
	//	Status string `json:"status"`
	//	Data   []struct {
	//		UID  int    `json:"uid"`
	//		Aqi  string `json:"aqi"`
	//		Time struct {
	//			Tz    string `json:"tz"`
	//			Stime string `json:"stime"`
	//			Vtime int    `json:"vtime"`
	//		} `json:"time"`
	//		Station struct { //много разных станций
	//			Name    string    `json:"name"`
	//			Geo     []float64 `json:"geo"`
	//			URL     string    `json:"url"`
	//			Country string    `json:"country"`
	//		} `json:"station,omitempty"`
	//	} `json:"data"`
	//}


	json.NewDecoder(resp.Body).Decode(&result)

	fmt.Println(result)
	fmt.Println(result["data"])



	//return result
}