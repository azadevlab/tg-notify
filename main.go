package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/urfave/cli/v2"
)

func sendTextToTelegramChat(chatID int, text, botToken, parseMode string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)

	message := map[string]interface{}{
		"chat_id":    chatID,
		"text":       text,
		"parse_mode": parseMode,
	}
	messageBody, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %v", err)
	}

	log.Printf("Sending message to Telegram chat ID %d: %s", chatID, text)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(messageBody))
	if err != nil {
		log.Printf("Failed to send message: %v", err)
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Unexpected status code: %d", resp.StatusCode)
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	log.Printf("Message successfully sent to chat ID %d", chatID)
	return nil
}

func createApp() *cli.App {
	var (
		telegramBotToken, telegramParseMode, ciProjectUrl, ciPipelineId, ciCommitName, ciCommitAuthor, text string
		telegramChatId int
	)

	return &cli.App{
		Name:    "notifyer",
		Version: "v1.0.1",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "telegram_bot_token",
				Aliases:     []string{"tbt"},
				Usage:       "Telegram Bot Token",
				EnvVars:     []string{"NOTIFYER_TELEGRAM_BOT_TOKEN"},
				Destination: &telegramBotToken,
			},
			&cli.IntFlag{
				Name:        "telegram_chat_id",
				Aliases:     []string{"tci"},
				Usage:       "Telegram Chat ID",
				EnvVars:     []string{"NOTIFYER_TELEGRAM_CHAT_ID"},
				Destination: &telegramChatId,
			},
			&cli.StringFlag{
				Name:        "telegram_parse_mode",
				Aliases:     []string{"tpm"},
				Usage:       "Telegram Parse Mode",
				EnvVars:     []string{"NOTIFYER_TELEGRAM_PARSE_MODE"},
				Value:       "HTML",
				Destination: &telegramParseMode,
			},
			&cli.StringFlag{
				Name:        "ci_commit_author",
				Aliases:     []string{"author"},
				Usage:       "CI Commit Author - GitLab CI",
				EnvVars:     []string{"CI_COMMIT_AUTHOR"},
				Destination: &ciCommitAuthor,
			},
			&cli.StringFlag{
				Name:        "ci_project_url",
				Aliases:     []string{"cpu"},
				Usage:       "CI Project URL - GitLab CI",
				EnvVars:     []string{"CI_PROJECT_URL"},
				Destination: &ciProjectUrl,
			},
			&cli.StringFlag{
				Name:        "ci_pipeline_id",
				Aliases:     []string{"cpi"},
				Usage:       "CI Pipeline ID - GitLab CI",
				EnvVars:     []string{"CI_PIPELINE_ID"},
				Destination: &ciPipelineId,
			},
			&cli.StringFlag{
				Name:        "ci_commit_ref_name",
				Aliases:     []string{"branch"},
				Usage:       "CI Commit Branch or Tag name - GitLab CI",
				EnvVars:     []string{"CI_COMMIT_REF_NAME"},
				Destination: &ciCommitName,
			},
			&cli.StringFlag{
				Name:        "text",
				Aliases:     []string{"t"},
				Usage:       "Text to send",
				EnvVars:     []string{"NOTIFYER_TEXT"},
				Destination: &text,
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "release",
				Aliases: []string{"r"},
				Usage:   "Release [ready]",
				Subcommands: []*cli.Command{
					{
						Name:  "ready",
						Usage: "Release ready",
						Action: func(c *cli.Context) error {
							fmt.Println("[Notifyer] Release ready")
							err := sendTextToTelegramChat(telegramChatId, text, telegramBotToken, telegramParseMode)
							if err != nil {
								log.Printf("Error sending release ready message: %v", err)
							}
							return err
						},
					},
				},
			},
			{
				Name:    "deploy",
				Aliases: []string{"d"},
				Usage:   "Deploy [done]",
				Subcommands: []*cli.Command{
					{
						Name:  "done",
						Usage: "Deploy done",
						Action: func(c *cli.Context) error {
							fmt.Println("[Notifyer] Deploy done")
							err := sendTextToTelegramChat(telegramChatId, text, telegramBotToken, telegramParseMode)
							if err != nil {
								log.Printf("Error sending deploy done message: %v", err)
							}
							return err
						},
					},
				},
			},
		},
	}
}

func main() {
	app := createApp()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
