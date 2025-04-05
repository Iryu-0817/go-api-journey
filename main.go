package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", HelloHandler)
	http.HandleFunc("/article", PostArticleHandler)
	http.HandleFunc("/article/list", ArticleListHandler)
	http.HandleFunc("/article/1", ArticleDetailHandler)
	http.HandleFunc("/article/nice", PostNiceHandler)
	http.HandleFunc("/comment", PostCommentHandler)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
