package main

import (
"net/http"
	"io"
	"strconv"
)
const (
	ALPHA=2.20462
)
func weightHandler( w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	var c string
	c=r.FormValue("kg")
	if c != "" {
		num,_:=strconv.ParseFloat(c,64)
		io.WriteString(w,strconv.Itoa(int(num*ALPHA)))

		return
	}
	c=r.FormValue("lbs")
	if ( c != ""){
		num,_:=strconv.ParseFloat(c,64)
		io.WriteString(w,strconv.Itoa(int(num/ALPHA)))
		return
	}
	http.Error(w, "bad request", 400)
}

func main (){
	http.HandleFunc("/weight", weightHandler)
	http.ListenAndServe(":5000",nil)
}