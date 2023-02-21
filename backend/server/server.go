package asciiartweb

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"os"
	"strings"
	"text/template"
)

const (
	b1 string = "81802daa5d0000076c70ba728c2deb73" // banner 1 "standard.txt"
	b2 string = "d44671e556d138171774efbababfc135" // banner 2 "shadow.txt"
	b3 string = "0021f26ad06f2f73a0cfa7b7d38d1434" // banner 3 "thinkertoy.txt"
)

// Структура где хранится вводимый текст, его результат, описание ошибки и код ошибки
type Data struct {
	InputText        string
	Result           string
	ErrorDescription string
	Error            int
}

// функция которая запускает сервер по порту 4040
func Server() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home_page)
	mux.HandleFunc("/ascii-art/", result_page)
	fmt.Printf("Listening server at port : http://localhost:4040\n")
	err := http.ListenAndServe(":4040", mux)
	if err != nil {
		return err
	}
	return nil
}

// функция обработки главной страницы аски арта
func home_page(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		Errors(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}
	if r.Method != http.MethodGet {
		Errors(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}
	tmpl, err := template.ParseFiles("frontend/index.html")
	if err != nil {
		Errors(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	err = tmpl.Execute(w, nil)
}

// функция обработки результата страницы аски арта
func result_page(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art/" {
		Errors(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}
	if r.Method != http.MethodPost {
		Errors(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}
	text := r.FormValue("inputText")
	font := r.FormValue("selectFont")
	for _, v := range text {
		if (v >= ' ' && v <= '~') || v == '\r' || v == '\n' {
			continue
		} else {
			Errors(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			return
		}
	}
	art, errorNum, errorDescript := bannerPrint(text, font)
	if errorNum != 0 {
		Errors(w, errorNum, errorDescript)
		return
	}
	Data := Data{InputText: r.FormValue("inputText"), Result: art}
	tmpl, err := template.ParseFiles("frontend/index.html")
	if err != nil {
		Errors(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	err = tmpl.Execute(w, Data)
}

// функция обработки ошибки страницы
func Errors(w http.ResponseWriter, errorNum int, errorDescript string) {
	tmpl, err := template.ParseFiles("frontend/error.html")
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(errorNum)
	Data := Data{Error: errorNum, ErrorDescription: errorDescript}
	err = tmpl.Execute(w, Data)
}

// функция обработки текста в аски арт, возвращает строку аски, код ошибки и описание ошибки
func bannerPrint(s string, bannerName string) (string, int, string) {
	var res string
	path := "./backend/banners/" + bannerName + ".txt"
	banner, err := os.ReadFile(path)
	if err != nil {
		return "", http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)
	}
	switch bannerName {
	case "standard":
		if checkHash(string(banner)) != b1 {
			return "", http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)
		}
	case "shadow":
		if checkHash(string(banner)) != b2 {
			return "", http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)
		}
	case "thinkertoy":
		if checkHash(string(banner)) != b3 {
			return "", http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)
		}
	default:
		return "", http.StatusBadRequest, http.StatusText(http.StatusBadRequest)

	}
	symbols := strings.Split(strings.ReplaceAll(string(banner), "\r", ""), "\n\n")
	for _, word := range strings.Split(s, "\r\n") {
		for lines := 0; lines < 8; lines++ {
			for _, v := range word {
				res += strings.Split(symbols[v-32], "\n")[lines]
			}
			res += "\n"
		}
	}
	return res, 0, ""
}

// функция проверки хэша баннеров на изменение
func checkHash(s string) string {
	hash := md5.New()
	hash.Write([]byte(s))
	return hex.EncodeToString(hash.Sum(nil))
}
