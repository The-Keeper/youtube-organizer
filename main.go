package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
)

const server_address = ":8000"

var config_dir = "~/youtube-organizer"
var data_dir = "~/youtube-organizer"

func main() {
	database_path := path.Join(data_dir, "database.sqlite")

	println(database_path)
	http.Handle("/static/",
		http.StripPrefix("/static", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./templates/index.html"))
		tmpl.Execute(w, nil)
	})

	// media route handler
	http.HandleFunc("/media/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "")
	})

	// watch route handler
	http.HandleFunc("/watch", func(w http.ResponseWriter, r *http.Request) {
		video_id := r.FormValue("v")
		fmt.Println(video_id)
		tmpl := template.Must(template.ParseFiles("./templates/video.html"))
		tmpl.Execute(w, nil)
	})

	log.Printf("App running on %s", server_address)
	log.Fatal(http.ListenAndServe(server_address, nil))
}
