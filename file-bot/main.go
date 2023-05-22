package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"INTRO TO INFO SYSTEM.pdf"} //we can have 100 of files here

	for i := 0; i < len(fileArr); i++ { //we are looping pver file array which cna be multiple file but in these video we are using a file but the file array is represented in file arr
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URL)
	}
}
