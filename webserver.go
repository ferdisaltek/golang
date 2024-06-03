package main

import (
	"fmt"
	"net/http"
)

// ana sayfa işleyici
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Home Page!")
}

// hakkında sayfası işleyici
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the About Page!")
}

func main() {
	// "/" yoluna gelen tüm istekleri homeHandler işlevine yönlendir
	http.HandleFunc("/", homeHandler)

	// "/about" yoluna gelen tüm istekleri aboutHandler işlevine yönlendir
	http.HandleFunc("/about", aboutHandler)

	// 8080 portunda web sunucusunu başlat
	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
