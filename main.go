package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	//Just a message for ensuring that localhost is running
	fmt.Println("Localhost is running on port 8888")

	http.HandleFunc("/", home)
	http.HandleFunc("/features", features)
	http.HandleFunc("/docs", docs)

	//for serving files like css, js, images etc
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("addTolink"))))

	//Starting the server
	http.ListenAndServe(":8888", nil)
}

func home(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp.Execute(w, nil)
	//fmt.Fprintf(w, "This is a sample web page")

}
func features(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp, err = ptmp.ParseFiles("wpage/features.gohtml")
	if err != nil {
		fmt.Println(err.Error())
		ptmp.Execute(w, nil)
		//fmt.Fprintf(w, "This is a sample web page")
	}
	ptmp.Execute(w, nil)
}
func docs(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp, err = ptmp.ParseFiles("wpage/docs.gohtml")
	if err != nil {
		fmt.Println(err.Error())
		ptmp.Execute(w, nil)
		//fmt.Fprintf(w, "This is a sample web page")
	}
	ptmp.Execute(w, nil)
}
