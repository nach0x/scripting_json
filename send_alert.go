package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/newrelic/go-agent/v3/newrelic"
)

func main() {

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("secops"),
		newrelic.ConfigLicense("AAAXX5554432DDDFDFDS"),
		func(config *newrelic.Config) {
			config.CustomInsightsEvents.Enabled = true
		},
	)
	if err != nil {
		panic(err)
	}

	Alert := Read_alert("phpmyadmin")

	if Alert == true {
		if err := Send_alert("AABBAA333AS"); err != nil {
			panic(err)
		}
		http.HandleFunc(newrelic.WrapHandleFunc(app, "/custom_event", customEvent))
	}

}

func Send_alert(token string) error {
	data := url.Values{}
	data.Add("token", token)
	data.Add("channel", "#secops")
	data.Add("username", "Defender_bot")
	data.Add("text", "Someone tried with phpmyadmin >>>> possible attack is happening!!!")

	resp, err := http.PostForm("https://slack.com/api/chat.postMessage", data)
	if err != nil {
		return err
	}
	return nil
}

func Read_alert(log string) bool {
	b, err := ioutil.ReadFile("/var/log/apache2/access.log")
	if err != nil {
		panic(err)
	}
	s := string(b)
	// //check whether s contains substring text
	return strings.Contains(s, log)
}

func customEvent(w http.ResponseWriter, r *http.Request) {
	txn := newrelic.FromContext(r.Context())

	io.WriteString(w, "recording a custom event")

	txn.Application().RecordCustomEvent("alert_event", map[string]interface{}{
		"myString": "a possible attack is happening",
		"myFloat":  0.603,
		"myInt":    123,
		"myBool":   true,
	})
}
