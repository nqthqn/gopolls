package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// Choice ...
type Choice struct {
	ID    int    `json:"id"`
	Text  string `json:"text"`
	Votes int    `json:"votes"`
}

// Poll ...
type Poll struct {
	Code     string   `json:"code"`
	Question string   `json:"question"`
	Choices  []Choice `json:"choices"`
}

var poll = Poll{
	Code:     "x := 3;",
	Question: "Is the semi colon required?",
	Choices:  []Choice{{0, "yes", 0}, {1, "no", 0}},
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./pub")))
	http.HandleFunc("/q", qHandler)
	http.HandleFunc("/r", rHandler)

	http.ListenAndServe(":3000", nil)
}

// Return the polls
func qHandler(w http.ResponseWriter, r *http.Request) {

	bs, _ := json.Marshal(poll)

	w.Write(bs)
}

// Take a response
func rHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("choiceID"))
	for i := range poll.Choices {
		if poll.Choices[i].ID == id {
			poll.Choices[i].Votes++
		}
	}

	http.Redirect(w, r, "/", 301)
}
