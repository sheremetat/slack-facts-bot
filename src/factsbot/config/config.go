package config

// Config fills from file config.xml
type Config struct {
	FactsDir       string `xml:"facts_dir"`
	SlackToken     string `xml:"slack_token"`
	SlackBouUserID string `xml:"slack_bot_user_id"`
}
