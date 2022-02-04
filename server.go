package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
)

func InitServe(addr string) error {
	if config.routes != "" {
		defaultPath = "./" + config.routes + "/"
	} else {
		defaultPath = ""
	}
	http.HandleFunc("/", HandlerRoutes)
	return http.ListenAndServe(addr, nil)
}

func Logger(r *http.Request, statusCode int) {
	log.Println(fmt.Sprintf("%v %v %v", r.Method, strconv.Itoa(statusCode), r.URL.Path))
}

func HandlerRoutes(w http.ResponseWriter, req *http.Request) {

	filename := defaultPath + req.URL.Path[1:]
	if last := len(filename) - 1; last >= 0 && filename[last] == '/' && len(filename) != 1 {
		filename = filename[:last]
	}
	if filename == "" {
		filename = "./"
	}
	file, err := os.Stat(filename)
	if os.IsNotExist(err) {
		_, _ = io.WriteString(w, "404 Not Found")
		Logger(req, http.StatusNotFound)
		return
	}
	if file.IsDir() {
		slash := ""
		files, err := ioutil.ReadDir(filename)
		if err != nil {
			http.Redirect(w, req, "", http.StatusInternalServerError)
			Logger(req, http.StatusInternalServerError)
		}
		if filename != "./" {
			if filename[len(filename)-1] != '/' {
				slash = "/"
			}
		}
		pwd := ""
		if req.URL.Path[1:] == "" {
			pwd = "/"
		} else {
			pwd = req.URL.Path[1:]
		}
		response := fmt.Sprintf(`<html><body><h4>Pwd: %v</h4><hr><ul>`, pwd)
		p := req.URL.Path
		if len(p) > 1 {
			base := path.Base(p)
			slice := len(p) - len(base) - 1
			url := "/"
			if slice > 1 {
				url = req.URL.Path[:slice]
				url = strings.TrimRight(url, "/")
			}
			response += fmt.Sprintf(`<li><a href="%v">..</a></li>`, url)
		}
		for _, f := range files {
			if f.Name()[0] != '.' {
				if f.IsDir() {
					response += fmt.Sprintf(`<li><a href="%v">%v/</a></li>`, req.URL.Path[0:]+slash+f.Name(), f.Name())
				} else {
					response += fmt.Sprintf(`<li><a href="%v">%v</a></li>`, req.URL.Path[0:]+slash+f.Name(), f.Name())
				}
			}
		}
		response = response + `</ul></body></html>`
		_, err = io.WriteString(w, response)
		if err != nil {
			http.Redirect(w, req, "", http.StatusInternalServerError)
			Logger(req, http.StatusInternalServerError)
		} else {
			Logger(req, http.StatusOK)
		}

		return
	}
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		http.Redirect(w, req, "", http.StatusInternalServerError)
		Logger(req, http.StatusInternalServerError)
		return
	}
	str := string(b)
	extension := path.Ext(filename)
	if extension == ".css" {
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
	} else if extension == ".js" {
		w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
	}
	_, err = io.WriteString(w, str)
	if err != nil {
		http.Redirect(w, req, "", http.StatusInternalServerError)
	} else {
		Logger(req, http.StatusOK)
	}
}
