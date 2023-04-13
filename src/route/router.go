package route

import (
	"net-http/myapp/controller"
	"net/http"
)

func GetRouter() *http.ServeMux {
	mux := http.NewServeMux()
	// 練習
	mux.HandleFunc("/", controller.Handler)
	mux.HandleFunc("/two", controller.HandlerTwo)

	// ブログ
	mux.HandleFunc("/blog", controller.FindByTitleHandler)
	mux.HandleFunc("/blog-post", controller.CreateBlogHandler)

	return mux
}
