package main

import (
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/veremchukvv/render"
)

var cfg = struct {
	Port     string
	UserAddr string
	WebAddr  string
}{
	Port:     "8083",
	UserAddr: "http://localhost:8080",
	WebAddr:  "http://localhost:8082",
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/checkout", checkoutFormHandler).Methods("GET")
	r.HandleFunc("/checkout", checkoutHandler).Methods("POST")
	render.SetTemplateDir(".")
	render.AddTemplate("payform", "payform.html")
	render.AddTemplate("msg", "msg.html")
	err := render.ParseTemplates()
	if err != nil {
		log.Fatalf("Init template error %v", err)
	}
	log.Printf("Starting on port %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, r))
}

func checkoutFormHandler(w http.ResponseWriter, r *http.Request) {
	uid := r.FormValue("uid")
	if uid == "" {
		render.RenderTemplate(w, "msg", Msg{
			"Не указан идентификатор пользователя",
			cfg.WebAddr,
		})
		return
	}
	render.RenderTemplate(w, "payform", struct{ Uid string }{uid})
}

type Msg struct {
	Msg     string
	BackURL string
}

func checkoutHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	uid := r.FormValue("uid")

	if !makePayment(
		r.FormValue("pan"),
		r.FormValue("date"),
		r.FormValue("cvc"),
	) {
		render.RenderTemplate(w, "msg", Msg{
			"неверные платёжные данные",
			"/checkout?uid=" + uid,
		})
	}
	err := requester.PatchJSON(
		cfg.UserAddr+"/user",
		url.Values{
			"id":      []string{uid},
			"is_paid": []string{"true"},
		},
		nil,
	)
	if err != nil {
		log.Printf("Payment error %v", err)
		render.RenderTemplate(w, "msg", Msg{
			"Во время проведения платежа произошла ошибка",
			"/checkout?uid=" + uid,
		})
		return
	}
	render.RenderTemplate(w, "msg", Msg{
		"Платёж успешно совершён",
		cfg.WebAddr,
	})
	return
}

func makePayment(pan, date, cvc string) bool {
	if pan != "4444444444444444" && date != "12/12" && cvc != "123" {
		return false
	}
	return true
}
