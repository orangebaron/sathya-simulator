package main

import (
	"io"
	"net/http"
	"log"
	"os"
	"strings"
)

//FOR FUTURE REFERENCE, FOOD IS REFERRED TO AS "GENDER"

func assumeGender(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, `
	<!DOCTYPE html>
	<html>
		<head>
			<title>wow you can order lunch im sathya edmmmmmmmmmadaka</title>
		</head>
		<body>
			<!-- wow this is my 10,000th line of code -->
			<form action='/come-out-of-the-closet.html'>
				waht ist thy naem?!<br>
				<input type="text" name="name" /><br>
				waht fdood do yao want?!<br>
				<input type="radio" name="gender" value="sandwich">yES,<br>
				<input type="radio" name="gender" value="pizza">boneless bizza<br>
				<input type="radio" name="gender" value="apple">give your soul to the devil<br>
				<input type='submit' value='plaec lnuch ordere'></input>
			</form>
		</body>
	</html>
	`)
}

func findOutGender(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		return
	}
	naem := strings.Replace(strings.Replace(req.FormValue("name"),"\n","mmmmmmmmm",-1),",","mmmmmmmmmm",-1)
	gender := strings.Replace(strings.Replace(req.FormValue("gender"),"\n","mmmmmmmmm",-1),",","mmmmmmmmmm",-1)
	f, err := os.OpenFile("orders.csv", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
	    panic(err)
	}
	
	defer f.Close()
	
	_, err = f.WriteString(naem+","+gender+"\n")
	
	if err != nil {
	    panic(err)
	}
	
	io.WriteString(w, `
	<!DOCTYPE html>
	<html>
		<head>
			<title>thanks for ordering lunch</title>
		</head>
		<body>
			thAnks FOR ORDERing lunch
		</body>
	</html>
	`)
}


func main() {
	http.HandleFunc("/what-is-your-gender.html", assumeGender)
	http.HandleFunc("/come-out-of-the-closet.html", findOutGender)
	log.Fatal(http.ListenAndServe(":80", nil))
}
