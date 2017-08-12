package main

import (
	"encoding/xml"
	"factsbot/bot"
	"factsbot/config"
	"io/ioutil"
	"os"

	facts "github.com/sheremetat/randfacts-lib"
	log "github.com/sirupsen/logrus"
)

func createConfig(filePath string) *config.Config {
	cfg := &config.Config{}

	cfgFile, err := os.Open(filePath)
	if err != nil {
		log.Error("Error opening config file: ", err)
		panic("Error opening config file")
	}
	defer cfgFile.Close()

	b, _ := ioutil.ReadAll(cfgFile)
	err = xml.Unmarshal(b, cfg)
	if err != nil {
		panic("Error parse congig file: " + err.Error())
	}

	return cfg
}

func main() {
	cfg := createConfig("./config.xml")
	log.Info("Have got parameners:")
	log.Infof("    -facts_dir: %s", cfg.FactsDir)
	log.Infof("    -slack_token: %s", cfg.SlackToken)
	log.Infof("    -slack_bot_user_id: %s", cfg.SlackBouUserID)

	factLib, err := facts.Init(cfg.FactsDir)
	if err != nil {
		log.Errorf("Error in main():  %v", err)
		panic("Error in main(): " + err.Error())
	}
	if err := bot.Run(cfg, factLib); err != nil {
		log.Errorf("Error in main():  %v", err)
	}
}
