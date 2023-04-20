package kpi_lab1

import (
	"net/http"
)

func main() {
	http.HandleFunc("/time", getTime)
	err := http.ListenAndServe(":8795", nil)
	handleError(err)
}

func GetTime(w http.ResponseWriter, r *http.Request) {

}

func handleError(err error) {

}
