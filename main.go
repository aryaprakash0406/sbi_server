package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var timestamp string = time.Now().Format(time.RFC3339)

type requestBodyAck struct {
	Serial_no     string `json:"serial_no"`
	ParticipantId string `json:"participantId"`
	Welfileflag   string `json:"welfileflag"`
	Exefileflag   string `json:"exefileflag"`
}
type requestBodyReg struct {
	Serial_no     string `json:"serial_no"`
	App_version   string `json:"app_version"`
	Audio_version string `json:"audio_version"`
	TMSIPPort     string `json:"TMSIPPort"`
	IMEI          string `json:"IMEI"`
}

func ackHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "text/plain")

	var requestData requestBodyAck

	err := json.NewDecoder(request.Body).Decode(&requestData)
	if err != nil {
		fmt.Println(timestamp, ":", "Error Decoding requestBody:", err)
	}
	fmt.Println(timestamp, ":", requestData)

	if requestData.Serial_no == "" || requestData.ParticipantId == "" || requestData.Welfileflag == "" || requestData.Exefileflag == "" {
		fmt.Println(timestamp, ":", "Missing required fields in requestBody")
		http.Error(response, "Missing required fields", http.StatusBadRequest)
		return
	}
	var welfileflag int
	welfileflag, err = strconv.Atoi(requestData.Welfileflag)
	fmt.Println(welfileflag)
	if err != nil {
		fmt.Println(timestamp, ":", "Error In Conversion From String To INT")
	}

	var exefileflag int
	exefileflag, err = strconv.Atoi(requestData.Exefileflag)
	if err != nil {
		fmt.Println(timestamp, ":", "Error In Conversion From String To INT")
	}
	fmt.Println(exefileflag)
	var responseData string
	if welfileflag >= 0 && exefileflag >= 0 {
		responseData = "00|Success"
	} else {
		responseData = "-1|Failed"
	}
	json.NewEncoder(response).Encode(responseData)
}

func regHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "text/plain")

	var requestData requestBodyReg

	err := json.NewDecoder(request.Body).Decode(&requestData)
	if err != nil {
		fmt.Println(timestamp, ":", "Error Decoding requestBody:", err)
	}
	fmt.Println(timestamp, ":", requestData)

	if requestData.Serial_no == "" || requestData.App_version == "" || requestData.Audio_version == "" || requestData.TMSIPPort == "" || requestData.IMEI == "" {
		fmt.Println(timestamp, ":", "Missing required fields in requestBody")
		http.Error(response, "Missing required fields", http.StatusBadRequest)
		return
	}

	check := requestData.Serial_no
	if check == "38240318750001" {
		responseData := "0000|35.200.193.157|1883|isu_client1|password1|0|| |6245fb7a-7701-42cb-936f-d7504b9a3133|F00KRQEVGV|0|0|1|http://34.93.102.191:8089/download/mp3/snd.mp3|1|http://34.93.102.191:8089/download/mp3/exc.txt/TMS/V1/SF/DEM/6245FB7A-7701-42CB-936F-D7504B9A3133_F00KRQEVGV_Exception.txt|1|0|"
		json.NewEncoder(response).Encode(responseData)
		return
	} else if check == "38240318750005" {
		responseData := "9999|35.200.193.157|1883|isu_client1|password1|0||http://34.93.102.191:8089/download/mp3/ota.bin|6245fb7a-7701-42cb-936f-d7504b9a3133|F00KRQEVGV|0|0|0| |1|https://xxxx:xxxx/TMS/V1/SF/DEM/6245FB7A-7701-42CB-936F-D7504B9A3133_F00KRQEVGV_Exception.txt|1|0|"
		json.NewEncoder(response).Encode(responseData)
		return
	} else if check == "38240424750001" {
		responseData := "9998|20.197.7.20|1884|||0|| |6245fb7a-7701-42cb-936f-d7504b9a3133|F00KRQEVGV|0|0|0| |1|https://xxxx:xxxx/TMS/V1/SF/DEM/6245FB7A-7701-42CB-936F-D7504B9A3133_F00KRQEVGV_Exception.txt|1|0|"
		json.NewEncoder(response).Encode(responseData)
		return
	} else if check == "38240424750005" {
		responseData := "9997|20.197.7.20|1884|||0|| |6245fb7a-7701-42cb-936f-d7504b9a3133|F00KRQEVGV|0|0|0| |1|https://xxxx:xxxx/TMS/V1/SF/DEM/6245FB7A-7701-42CB-936F-D7504B9A3133_F00KRQEVGV_Exception.txt|1|0|"
		json.NewEncoder(response).Encode(responseData)
		return
	} else {
		responseData := "Serial No Not Found"
		json.NewEncoder(response).Encode(responseData)
		return
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ack", ackHandler).Methods("POST")
	r.HandleFunc("/reg", regHandler).Methods("POST")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PATCH"},
		// AllowedMethods: []string{"*"},
		AllowedHeaders: []string{"Authorization", "Content-Type"},
		// AllowedHeaders: []string{"*"},
	})

	handler := c.Handler(r)
	port := ":8090"
	s := &http.Server{
		Addr:    port,
		Handler: handler,
	}
	log.Printf("Server Started On Port:%v", port)
	s.ListenAndServe()
}
