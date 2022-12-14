package types

type Quiz struct {
	Question string   `json:"question"`
	Answers  []Answer `json:"answers"`
}

type Answer struct {
	IsCorrect bool   `json:"-"`
	Key       string `json:"key"`
	Value     string `json:"value"`
}
