package  main

import ("log"
        "net/http"
        "html/template"
)

type logs struct {
  Logs    [][4]string
}

type nil_logs struct {
  Logs    [][1]string
}

func (l *logs) get_logs (id string) {
  l.Logs = get_logs(id)
}

func home_page(w http.ResponseWriter, r *http.Request) {
  l := logs{ Logs: [][4]string {{"", "", "", ""}}}

  if r.FormValue("find") != "" { l.get_logs(r.FormValue("find")) }
  if l.Logs[0][3] == "" {
    nl := nil_logs{Logs: [][1]string {{"Unfortunately your track isn`t in the database 😔"}} }
    t, _ := template.ParseFiles("templace/index.html")
    t.Execute(w, nl)
  } else {
    t, _ := template.ParseFiles("templace/index.html")
    t.Execute(w, l)
  }
}

func pageHeaders() {
  http.Handle("/styles/",
             http.StripPrefix("/styles/",
                              http.FileServer(http.Dir("./styles/"))))
http.HandleFunc("/", home_page)
http.ListenAndServe(":8080", nil)
}

func main() {
  log.Println("==== Start ====")
  pageHeaders()
}
