package main // import "github.com/cblecker/action-tyn"

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"encoding/json"
	"github.com/google/go-github/v21/github"
	"golang.org/x/oauth2"
)

func main() {
	// pull in the event data file path from the environment
	jsonFilePath := os.Getenv("GITHUB_EVENT_PATH")
	// open the json event file
	jsonFile, err := os.Open(jsonFilePath)
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	jsonByte, _ := ioutil.ReadAll(jsonFile)

	// unmarshal the json into a release event
	var prEvent github.PullRequestEvent
	err = json.Unmarshal(jsonByte, &prEvent)
	if err != nil {
		fmt.Println(err)
	}

	// only comment if PR is closed, but not merged
	if *prEvent.PullRequest.State == "closed" && *prEvent.PullRequest.Merged == false {

		// pull in token from the environment
		ghToken := os.Getenv("GITHUB_TOKEN")

		// set up connection to GitHub
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: ghToken},
		)
		tc := oauth2.NewClient(ctx, ts)
		client := github.NewClient(tc)

		// format and set up the new comment
		var comment github.IssueComment
		commentBody := "thank u, next"
		comment.Body = &commentBody

		// post the comment to GitHub
		_, _, err = client.Issues.CreateComment(ctx, *prEvent.Repo.Owner.Login, *prEvent.Repo.Name, *prEvent.Number, &comment)
		if err != nil {
			fmt.Println(err)
		}
	}
}
