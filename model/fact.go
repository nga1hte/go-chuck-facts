package model

import (
	"strings"
)

type Fact struct {
	Categories []string `json:"categories"`
	URL        string   `json:"url"`
	Value      string   `json:"value"`
}

type categories struct {
	Values []string
}

var Categories = categories{
	Values: []string{"animal", "career", "celebrity", "dev", "explicit", "fashion", "food", "history", "money", "movie", "music", "political", "religion", "science", "sport", "travel"},
}

func (c categories) PrintCategories() string {
	return strings.Join(c.Values, ", ")
}

func (c categories) MapValues() map[string]bool {
	cat := make(map[string]bool)
	for _, val := range c.Values {
		cat[val] = true
	}
	return cat
}
