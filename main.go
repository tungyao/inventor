package main

import (
	"flag"
	"fmt"
	uc "github.com/tungyao/ultimate-cedar"
	"log"
	"net/http"
	"os"
)

var (
	_port string
)

func main() {
	flag.StringVar(&_port, "port", "9002", "port")
	flag.Parse()
	if env := os.Getenv("PORT"); env != "" {
		_port = env
	}
	r := uc.NewRouter()
	var port string = _port
	if len(os.Args) > 1 {
		port = os.Args[1]
	}
	log.Println("listen on " + port)
	r.Get("/", Index)
	r.Get("/static/app.css", func(writer uc.ResponseWriter, request uc.Request) {
		writer.Header().Set("content-type", "text/css")
		writer.Header().Set("Cache-Control", "public, max-age=31536000")
		writer.Header().Set("Content-Length", fmt.Sprintf("%d", len(css)))
		writer.Write(css)
	})
	r.Get("/static/app.js", func(writer uc.ResponseWriter, request uc.Request) {
		writer.Header().Set("Cache-Control", "public, max-age=31536000")
		writer.Header().Set("content-type", "application/javascript")
		writer.Header().Set("Content-Length", fmt.Sprintf("%d", len(js)))
		writer.Write(js)
	})
	r.Get("/del", func(writer uc.ResponseWriter, request uc.Request) {
		id := request.URL.Query().Get("id")
		writer.Header().Set("content-type", "application/json")
		if id != "" {
			_, err := db.Exec("delete from main.data where id=?", id)
			if err != nil {
				log.Println(err)
				writer.Data(`{"ok":"no","msg":"删除失败"}`)
				return
			}
			writer.Data(`{"ok":"yes"}`)
		}
		writer.Data(`{"ok":"no","msg":"id不能为空"}`)
	})
	r.Post("/save", func(writer uc.ResponseWriter, request uc.Request) {
		dt := &data{}
		writer.Header().Set("content-type", "application/json")
		if err := IoRead(request, dt); err != nil {
			log.Println(err)
			writer.Data(`{"ok":"no","msg":"数据格式错误"}`)
			return
		}
		n, err := db.Exec("insert into main.data (`name`,`model`,`footprint`,`number`)values(?,?,?,?)", dt.Name, dt.Model, dt.Footprint, dt.Number)
		if err != nil {
			log.Println(err)
			writer.Data(`{"ok":"no","msg":"保存失败"}`)
		}
		nm, _ := n.LastInsertId()
		writer.Data(fmt.Sprintf(`{"ok":"yes","msg":"保存成功","id":%d}`, nm))
	})
	r.Post("/edit", func(writer uc.ResponseWriter, request uc.Request) {
		dt := &data{}
		writer.Header().Set("content-type", "application/json")
		if err := IoRead(request, dt); err != nil {
			log.Println(err)
			writer.Data(`{"ok":"no","msg":"数据格式错误"}`)
			return
		}
		if _, err := db.Exec("update main.data set `name`=?,`model`=?,`footprint`=?,`number`=? where id=?", dt.Name, dt.Model, dt.Footprint, dt.Number, dt.Id); err != nil {
			log.Println(err)
			writer.Data(`{"ok":"no","msg":"保存失败"}`)
		} else {
			writer.Data(`{"ok":"yes","msg":"保存成功"}`)
		}
		return
	})
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalln(err)
	}
}
