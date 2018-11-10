package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	//解析参数，默认是不会解析的
	r.ParseForm()
	log.Println("one request, Method:", r.Method)
	// 请按照公众平台官网\基本配置中信息填写
	// token := "webtest"
	if r.Method == "GET" {
		for k, v := range r.Form {
			log.Printf("key=%s, val=%s", k, strings.Join(v, " "))

		}
	} else if r.Method == "POST" {

	}

	fmt.Fprintf(w, "Hello there!\n")
}

func main() {
	http.HandleFunc("/wx", myHandler) //	设置访问路由
	log.Fatal(http.ListenAndServe(":80", nil))
}
