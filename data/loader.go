// PopSenzawa Echo
// (c) 2023 SuperSonic (https://github.com/supersonictw).

package data

import (
	"time"
)

const (
	unackedLimit = 1000
)

var (
	uploader *Uploader
	broker   *Broker
	ticker   *time.Ticker
)

func init() {
	uploader = NewUploader()
	broker = NewBroker()
	ticker = time.NewTicker(1 * time.Second)
}

func init() {
	MessageQueue.StartConsuming(unackedLimit, 5*time.Second)
	MessageQueue.AddConsumer("uploader", uploader)
	MessageQueue.AddConsumer("broker", broker)
}

func init() {
	go func() {
		for range ticker.C {
			uploader.Wave()
		}
	}()
}