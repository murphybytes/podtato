package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/user"
	"strings"
)

//go:embed img/*
var contents embed.FS

var page string = `<html>
<head>
<title>Podtato</title>
</head>
<body>
<h1>Podtato</h1>
<img src="img/podtato2.png" >
<p>
<h3>User</h3>
{{ .User }}
</p>
<p>
<h3>Working Directory</h3>
{{ .WorkDir }}
</p>
<p>
<h3>Environment</h3>
<table>

{{range .Environment}}
<tr><td>{{.}}</td></tr>
{{end}}
</table>
</p>
</body>
</html>
`

func main() {
	http.HandleFunc("/img/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "image/png")

		data, err := contents.ReadFile(strings.TrimLeft(r.URL.Path, "/"))
		if err != nil {

			log.Printf("could not read %q error %s", r.URL.Path, err)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if _, err := w.Write(data); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.New("the page").Parse(page)
		if err != nil {
			log.Printf("could not parse template %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		data := struct {
			WorkDir     string
			User        string
			Environment []string
		}{
			WorkDir: func() string {
				s, err := os.Getwd()
				if err != nil {
					return "N/A"
				}
				return s
			}(),
			Environment: os.Environ(),
			User: func() string {
				u, err := user.Current()
				if err != nil {
					return "N/A"
				}
				return fmt.Sprintf("%s:%s", u.Username, u.Uid)
			}(),
		}
		if err := tmpl.Execute(w, data); err != nil {
			log.Printf("error executing tempate %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	})
	s := &http.Server{
		Addr: ":8080",
	}

	log.Fatal(s.ListenAndServe())

}
