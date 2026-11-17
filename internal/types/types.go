package types

import (
	"encoding/json"
)

type GitResponse struct {
	Login	string	`json:"login"`
	Id		int		`json:"id"`
}

type Actor struct {
	Id int	`json:"id"`
	Login string `json:"login"`
}

type Repo struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Url string `json:"url"`
}

type PushEvent struct {
	RepositoryId int `json:"repository_id"`
	PushId int `json:"push_id"`
	Ref string `json:"ref"`
	Head string `json"head"`
	Before string `json:"before"`
}

type GitEvent struct {
	Id	string		`json:"id"`
	Type string	`json:"type"`
	Actor Actor `json:"actor"`
	Repo Repo `json:"repo"`
	Payload json.RawMessage `json:"payload"`
	CreatedAt string `json:"created_at"`
}
