package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"bytes"
)

type MpesaRequest struct{
	ShortCode string `json:"shortcode"`
	CommandID string `json:"commandid"`
	Amount string `json:"amount"`
	Msisdn string `json:"msisdn"`
	AccountReference string `json:"accountreference"`
	TransactionDesc string `json:"transactiondesc"`
}

type MpesaResponse struct{
	ConversationID string `json:"conversationid"`
	OriginatorCoversationID string `json:"originatorconversationid"`
	ResponseDescription string `json:"responsedescription"`
}

func main(){
	// create a new request
	mpesaRequest := &MpesaRequest{
		ShortCode: "174379",
		CommandID: "CustomerPayBillOnline",
		Amount: "1",
		Msisdn: "254708374149",
		AccountReference: "123456",
		TransactionDesc: "test",
	}
	// convert to json
	jsonValue, _ := json.Marshal(mpesaRequest)
	// create a new request
	request, err := http.NewRequest("POST", "https://sandbox.safaricom.co.ke/mpesa/stkpush/v1/processrequest", bytes.NewBuffer(jsonValue))
	if err != nil{
		log.Fatal(err)
	}
	// set the headers
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer bearer-token-here")
	// create a new client
	client := &http.Client{}
	// send the request
	response, err := client.Do(request)
	if err != nil{
		log.Fatal(err)
	}
	// read the response
	data, _ := ioutil.ReadAll(response.Body)
	// convert to json
	var mpesaResponse MpesaResponse
	json.Unmarshal(data, &mpesaResponse)
	// print the response
	fmt.Printf("ConversationID: %sOriginatorCoversationID: %sResponseDescription: %s", mpesaResponse.ConversationID, mpesaResponse.OriginatorCoversationID, mpesaResponse.ResponseDescription)
}