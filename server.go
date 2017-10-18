package hotpot

import (
	"encoding/json"
	"log"
	"net/http"
)

// Handler is the main handler fetching the history and adding a text
func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(404)
		w.Write([]byte("Method not allowed on route\n"))
		return
	}
	p := Potato{}
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		log.Printf("Error parsing potatoe structure: %#v\n", err.Error())
		w.WriteHeader(400)
		w.Write([]byte("Error parsing JSON"))
		return
	}
	payload := Entry{
		Node: "Mat's translation",
		Text: LabelLanguages(p.History),
		Desc: desc,
	}
	p.History = append(p.History, payload)
	p.Text = payload.Text
	b, err := json.Marshal(p)
	if err != nil {
		log.Printf("Error parsing potatoe structure: %#v\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error marshaling response"))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("content", "application/json")
	w.Write(b)
}
