package main

import (
"net/http"
	"fmt"
	"encoding/json"
)
const (
	ALPHA=2.20462
	KILOGRAM="k"
	POUND="l"
)

type WeightResponse struct {
	Unit string
	Result string
}

func pToK(val string) WeightResponse {
	return WeightResponse{"Kilogram",val}
}

func kToP(val string) WeightResponse {
	return WeightResponse{"Pounds",val}
}


func weightHandler( w http.ResponseWriter, r *http.Request){
	fmt.Println("GET params:", r.URL.Query());
	wtype :=r.URL.Query().Get("type")
	weight :=r.URL.Query().Get("val")
	if  wtype != "" && weight != ""{
		switch (wtype) {
			case POUND:
				response:= pToK(weight)
				js, err := json.Marshal(response)
				if err == nil {
					w.Header().Set("Content-Type","application/json")
					w.Write(js)
				}else{
					fmt.Print(err)
				}
			case KILOGRAM:
				response:= kToP(weight)
				js, err := json.Marshal(response)
				if err == nil {
					w.Header().Set("Content-Type","application/json")
					w.Write(js)
				}else{
					fmt.Print(err)
				}
		default:
			panic(nil)
		}
	} else{
		http.Error(w, "bad request", 400)
	}
}

func main (){
	http.HandleFunc("/weight", weightHandler)
	http.ListenAndServe(":5000",nil)
}