package main

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/pkg/errors"
)

type Slack struct {
	Blocks []Block `json:"blocks"`
}

type Block struct {
	Type string `json:"type"`
	Text Text   `json:"text"`
}

type Text struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func run() error {
	incomingUrl := "https://hooks.slack.com/services/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"

	slackMap := Slack{
		Blocks: []Block{
			Block{
				Type: "section",
				Text: Text{
					Type: "mrkdwn",
					Text: "検証",
				},
			},
			Block{
				Type: "section",
				Text: Text{
					Type: "mrkdwn",
					Text: "検証",
				},
			},
		},
	}

	p, _ := json.Marshal(slackMap)

	resp, err := http.PostForm(
		incomingUrl,
		url.Values{"payload": {string(p)}},
	)

	if err != nil {
		return errors.WithStack(err)
	}

	defer resp.Body.Close()

	return nil
}

func main() {
	lambda.Start(run)
}
