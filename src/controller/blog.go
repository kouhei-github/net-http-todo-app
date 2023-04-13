package controller

import (
	"encoding/json"
	"net-http/myapp/repository"
	"net-http/myapp/service"
	"net/http"
)

func CreateBlogHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "POST" {
		response := Response{Status: 405, Text: "Method Not Allowed"}
		// Header情報の追加方法
		//header := w.Header()
		//header.Set("cookie", "aaaa")
		w.WriteHeader(405)
		json.NewEncoder(w).Encode(response)
		return
	}
	var blogEntity repository.BlogEntity

	if err := json.NewDecoder(r.Body).Decode(&blogEntity); err != nil {
		service.WriteLogFile(err.Error())
		return
	}
	if err := blogEntity.Create(); err != nil {
		service.WriteLogFile(err.Error())
		return
	}
	service.WriteLogFile("完了しました")
}

func FindByTitleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		response := Response{Status: 405, Text: "Method Not Allowed"}
		// Header情報の追加方法
		//header := w.Header()
		//header.Set("cookie", "aaaa")
		w.WriteHeader(405)
		json.NewEncoder(w).Encode(response)
	}
	query := r.URL.Query()
	// このやり方もある
	//titles := query["title"]
	//title := titles[0]

	title := query.Get("title")

	blogEntity := repository.BlogEntity{Title: title}
	entities, err := blogEntity.FindByTitle()
	if err != nil {
		service.WriteLogFile(err.Error())
		return
	}
	if err := json.NewEncoder(w).Encode(entities); err != nil {
		service.WriteLogFile(err.Error())
		return
	}
}
