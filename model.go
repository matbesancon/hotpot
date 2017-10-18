package hotpot

import (
	"strings"

	"github.com/mbesancon/hotpot/detector"
)

// Potato is the whole structure received from the client
type Potato struct {
	Text    string  `json:"text"`
	History []Entry `json:"history"`
}

// Entry is an identified text chunk
type Entry struct {
	Node string `json:"node"`
	Text string `json:"text"`
	Desc string `json:"desc"`
}

// LabelLanguages from a slice of entries
func LabelLanguages(entries []Entry) string {
	result := []string{}
	for _, entry := range entries {
		if language := detector.GuessLanguage(entry.Text); language != "undefined" {
			result = append(result, language+": "+entry.Text)
		}
	}
	if len(result) > 0 {
		return strings.Join(result, ";")
	}
	return "No language could be detected :("
}
