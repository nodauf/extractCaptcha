package main

import (
	_ "embed"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

//go:embed "captcha.tmpl"
var captchaTemplate string

func main() {
	key := flag.String("key", "", "Captcha's key")
	formTarget := flag.String("formTarget", "", "Target of the form. Webhook, requestbin,...")
	flag.Parse()
	if *key == "" || *formTarget == "" {
		fmt.Println(`key and formTarget arguments are required. Use -key ... -formTarget ...
Example: ./captcha -key 6LfT49ocAAAAANbrj7pDsLtlFe5maqwvceH8k3Dr -formTarget https://requestbin.io/w1xz2bw1`)
		os.Exit(1)
	}
	t := template.New("")
	tmpl := template.Must(t.Parse(captchaTemplate))
	data := struct {
		Key        string
		FormTarget string
	}{
		Key:        *key,
		FormTarget: *formTarget,
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Serving captcha.tmpl")
		err := tmpl.Execute(w, data)
		if err != nil {
			log.Println("Error: " + err.Error())
		}
	})
	log.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)

}
