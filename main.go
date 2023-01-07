package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析參數，預設是不會解析的
	// fmt.Println(r.Form) //這些資訊是輸出到伺服器端的列印資訊
	// fmt.Println("path", r.URL.Path)
	// fmt.Println("scheme", r.URL.Scheme)
	// fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	str := ""
	if len(r.URL.Path) > 1 {
		str = r.URL.Path[1:]
	} else {
		str = "World"
	}
	fmt.Fprintf(w, "Hello %v!", str) //這個寫入到 w 的是輸出到客戶端的
}

func main() {
	http.HandleFunc("/", sayhelloName)       //設定存取的路由
	err := http.ListenAndServe(":9090", nil) //設定監聽的埠
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
