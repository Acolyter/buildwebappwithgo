package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"
    "crypto/md5"
    "io"
    "strconv"
    "os"
)

func main() {
	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/login", login)
    http.HandleFunc("/upload", upload)
    
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("*******")
	fmt.Println("path", r.URL.Scheme)
	fmt.Println("*******")
	fmt.Println(r.Form["url_long"])
	fmt.Println("*******")
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Println("*******")
	fmt.Fprintf(w, "Hello there!")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("methon: ", r.Method) //获取请求的方法
	if r.Method == "GET" {            // GET 获取当前页内容
		crutime := time.Now().Unix()
        h := md5.New()
        io.WriteString(h, strconv.FormatInt(crutime, 10))
        token := fmt.Sprintf("%x", h.Sum(nil))
        
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, token)
	} else {//此例此处为 POST 当前页账密
		r.ParseForm()
        token := r.Form.Get("token")
        if token != "" {
            //验证token
        }else {
            //不存在的token
        }
        fmt.Println("username length:", len(r.Form["username"][0]))//输出到服务端
        fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username")))//输出到服务端
        fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))//输出到服务端
        template.HTMLEscape(w, []byte(r.Form.Get("username")))//输出到客户端
        // template.HTMLEscape(w, []byte(r.Form.Get("token")))
        
		// fmt.Println("username: ", r.Form["username"])
		// fmt.Println("password: ", r.Form["password"])
	}
}

func upload(w http.ResponseWriter, r *http.Request)  {
    fmt.Println("method: ", r.Method)
    if r.Method == "GET"{
        crutime := time.Now().Unix()
        h := md5.New()
        io.WriteString(h, strconv.FormatInt(crutime, 10))
        token := fmt.Sprintf("%x", h.Sum(nil))

        t, _ := template.ParseFiles("upload.gtpl")
        t.Execute(w, token)
    }else {
        r.ParseMultipartForm(32 << 20)
        file, handler, err := r.FormFile("uploadfile")
        if err != nil{
            fmt.Println(err)
            return
        }
        defer file.Close()
        f, err := os.OpenFile("./test/" +handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)//注意需要提前手动创建./test/目录
        if err != nil {
            fmt.Println(err)
            return
        }
        defer f.Close()
        io.Copy(f, file)
    }
}
