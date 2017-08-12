package bot

import (
	"factsbot/config"
	"os"
	"os/signal"
	"syscall"

	"github.com/nlopes/slack"
	facts "github.com/sheremetat/randfacts-lib"
	log "github.com/sirupsen/logrus"
)

// Run listener deamon
func Run(cfg *config.Config, facts *facts.FactsLib) error {
	// Initialization of Slack Bot
	api := slack.New(cfg.SlackToken)
	api.SetDebug(true)

	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		log.Print("Event Received: ")
		switch ev := msg.Data.(type) {
		case *slack.HelloEvent:
			// Ignore hello
		case *slack.ConnectedEvent:
			log.Println("Infos:", ev.Info)
			log.Println("Connection counter:", ev.ConnectionCount)

		case *slack.MessageEvent:
			log.Printf("Message: %v\n", ev.Msg)
			fact, err := facts.FindFact(ev.Msg.Text)
			if err == nil && len(fact) > 0 {
				outMsg := rtm.NewOutgoingMessage(fact, cfg.SlackBouUserID)
				outMsg.Channel = ev.Channel
				rtm.SendMessage(outMsg)
			}

		case *slack.PresenceChangeEvent:
			log.Printf("Presence Change: %v\n", ev)

		case *slack.LatencyReport:
			log.Printf("Current latency: %v\n", ev.Value)

		case *slack.RTMError:
			log.Printf("Error: %s\n", ev.Error())

		case *slack.InvalidAuthEvent:
			log.Printf("Invalid credentials")

		default:
			// Ignore other events..
			// fmt.Printf("Unexpected: %v\n", msg.Data)
		}
	}

	waitForSignal()
	return nil
}

func waitForSignal() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	s := <-ch
	log.Infof("Got signal: %v, exiting.", s)
}
