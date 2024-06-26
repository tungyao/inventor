package main

import (
	"database/sql"
	_ "embed"
	_ "github.com/mattn/go-sqlite3"
	uc "github.com/tungyao/ultimate-cedar"
	"html/template"
	"log"
	"math"
)

//go:embed index.html
var indexHtml string

//go:embed css/bootstrap.min.css
var css []byte

//go:embed js/bootstrap.js
var js []byte

var db *sql.DB

func init() {
	initDb()
}
func initDb() {
	d, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatalln(err)
	}
	db = d
}

type data struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Model     string `json:"model"`
	Footprint string `json:"footprint"`
	Number    string `json:"number"`
}

func Index(w uc.ResponseWriter, r uc.Request) {
	page, limit := pagination(r)
	s := indexHtml

	var search = r.Query.Get("search")

	t, err := template.New("index.html").Parse(string(s))
	if err != nil {
		log.Println()
	}
	var rows *sql.Rows

	if search != "" {
		rows, err = db.Query("select * from main.data where name like ? or model like ? or footprint like ? or number like ? limit ? offset ?", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", limit, page*limit)
	} else {
		rows, err = db.Query("select * from main.data limit ? offset ?", limit, page*limit)
	}
	if err != nil {
		log.Println(err)
	}
	outData := struct {
		Data     []*data
		Count    int
		Page     int
		PrePage  int
		NextPage int
		NowPage  int
		Limit    int
	}{}
	outData.Data = make([]*data, 0)
	for rows.Next() {
		d := &data{}
		err = rows.Scan(&d.Id, &d.Name, &d.Model, &d.Footprint, &d.Number)
		if err != nil {
			log.Println(err)
		}
		outData.Data = append(outData.Data, d)
	}

	var row *sql.Row

	if search != "" {
		row = db.QueryRow("select count(*) from main.data where name like ? or model like ? or footprint like ? or number like ?", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	} else {
		row = db.QueryRow("select count(*) from main.data")
	}
	row.Scan(&outData.Count)

	outData.Page = int(math.Ceil(float64(outData.Count) / float64(limit)))
	outData.PrePage = page
	outData.NextPage = page + 2
	outData.Limit = limit
	if outData.PrePage < 0 {
		outData.PrePage = 0
	}
	outData.NowPage = page + 1
	err = t.Execute(w.ResponseWriter, outData)
}
