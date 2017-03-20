package main

import (
	"net/http"
	"fmt"
	"log"
	"html/template"
)

// cookieの設定を行う
func setCookies(w http.ResponseWriter, r *http.Request) {
	// ステータスコードのみ設定
	cookie := &http.Cookie{
		Name: "hoge",
		Value: "bar",
	}
	http.SetCookie(w, cookie)

	fmt.Fprintf(w, "Cookieの設定ができたよ")
}

func showCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("hoge")

	if err != nil {
		fmt.Println("cookieにお探しの値は存在しませんでした")
	}

	// 存在する場合は、モードセレクトを表示
	// モード選択のテンプレートを表示
	tmpl := template.Must(template.ParseFiles("./cookie.html"))
	tmpl.Execute(w, cookie)



}

func main() {
	http.HandleFunc("/", setCookies)
	http.HandleFunc("/cookie", showCookie)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

