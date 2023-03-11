package main

import (
	"fmt"
	"net/http"
)

type Human struct {
	FName string
	LName string
	Age   int
}

func (hum Human) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hum.FName = "Metehan"
	hum.LName = "Sarı"
	hum.Age = 23
	// Formu parse etmek için
	r.ParseForm()

	// Sunucdan Form bilgisi almak için
	fmt.Println(r.Form)

	//URL' in path bilgisini almak için
	fmt.Println("path", r.URL.Path)

	fmt.Fprintf(w, "<table>"+
		"<tr>"+
		"<td><b>Ad</b></td>"+
		"<td><b>SoyAd</b></td>"+
		"<td><b>Yaş</b></td>"+
		"</tr>"+
		"<tr>"+
		"<td>"+hum.FName+"</td>"+
		"<td>"+hum.LName+"</td>"+
		"<td>"+fmt.Sprint(hum.Age)+"</td>"+
		"</tr>"+
		"<tr></tr>"+
		"<tr></tr>"+
		"<tr></tr>"+
		"<tr>"+
		"<td><b> Path</b></td>"+
		"<td> "+r.URL.Path+"</td>"+
		"</tr>"+
		"</table>")

}
func main() {
	var hum Human
	err := http.ListenAndServe("localhost:9000", hum)
	CheckError(err)
}
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
