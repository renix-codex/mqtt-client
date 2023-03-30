package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	// create MQTT client options
	opts := MQTT.NewClientOptions().AddBroker("0.0.0.0:1883")
	opts.SetClientID("go-client")

	// create MQTT client
	client := MQTT.NewClient(opts)

	// connect to MQTT broker
	for {
		if token := client.Connect(); token.Wait() && token.Error() != nil {
			log.Printf("MQTT connect error: %v", token.Error())
			time.Sleep(5 * time.Second)
			continue
		}
		break
	}

	// subscribe to MQTT topic
	topic := "my/topic"
	qos := 0
	token := client.Subscribe(topic, byte(qos), func(client MQTT.Client, msg MQTT.Message) {
		fmt.Printf("received message: %s\n", string(msg.Payload()))
	})
	token.Wait()

	// publish MQTT message
	message := "hello, world!"
	token = client.Publish(topic, byte(qos), false, message)
	token.Wait()
	fmt.Printf("message published: %s\n", message)

	// handle OS signals
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	<-sigchan

	// disconnect from MQTT broker
	client.Disconnect(250)
}
