package main

import "time"

type CreateResponseJson struct {
	Content struct {
		Name        string `json:"name"`
		Path        string `json:"path"`
		Sha         string `json:"sha"`
		Size        int    `json:"size"`
		URL         string `json:"url"`
		HTMLURL     string `json:"html_url"`
		GitURL      string `json:"git_url"`
		DownloadURL string `json:"download_url"`
		Type        string `json:"type"`
		Links       struct {
			Self string `json:"self"`
			Git  string `json:"git"`
			HTML string `json:"html"`
		} `json:"_links"`
	} `json:"content"`
	Commit struct {
		Sha     string `json:"sha"`
		NodeID  string `json:"node_id"`
		URL     string `json:"url"`
		HTMLURL string `json:"html_url"`
		Author  struct {
			Date  time.Time `json:"date"`
			Name  string    `json:"name"`
			Email string    `json:"email"`
		} `json:"author"`
		Committer struct {
			Date  time.Time `json:"date"`
			Name  string    `json:"name"`
			Email string    `json:"email"`
		} `json:"committer"`
		Message string `json:"message"`
		Tree    struct {
			URL string `json:"url"`
			Sha string `json:"sha"`
		} `json:"tree"`
		Parents []struct {
			URL     string `json:"url"`
			HTMLURL string `json:"html_url"`
			Sha     string `json:"sha"`
		} `json:"parents"`
		Verification struct {
			Verified  bool        `json:"verified"`
			Reason    string      `json:"reason"`
			Signature interface{} `json:"signature"`
			Payload   interface{} `json:"payload"`
		} `json:"verification"`
	} `json:"commit"`
}
