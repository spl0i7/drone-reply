package main

import (
	"context"
	"github.com/BurntSushi/toml"
	"github.com/eclipse/paho.golang/paho"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	if len(os.Args) == 1 {
		log.Fatalln("Please specify configuration file ", os.Args[0], "<path to config>")
	}

	var config Config
	if _, err := toml.DecodeFile(os.Args[1], &config); err != nil {
		log.Fatalln(err)
	}

	mqttClient, err := config.NewMQTTClient()
	if err != nil {
		log.Fatalln(err)
	}

	handler := NewMessageHandler(mqttClient)

	router := paho.NewStandardRouter()
	router.RegisterHandler("sys/product/+/status", handler.FlightStatusHandler)
	// uncomment to debug reply messages
	// router.RegisterHandler("sys/product/+/status_reply", handler.DebugMessage)
	mqttClient.Router = router

	if err := config.ConnectMQTT(mqttClient); err != nil {
		log.Fatalln(err)
	}
	log.Println("Connected to MQTT")

	if _, err := mqttClient.Subscribe(context.Background(), &paho.Subscribe{
		Subscriptions: map[string]paho.SubscribeOptions{
			"sys/product/+/status": {QoS: 0, NoLocal: false},
			// uncomment to debug reply messages
			// "sys/product/+/status_reply": {QoS: 0, NoLocal: false},
		},
	}); err != nil {
		log.Fatalln(err)
	}
	log.Println("Subscribed to Topics")

	// Uncomment following to publish dummy data
	//go func() {
	//	time.Sleep(time.Second * 5)
	//	_, _ = mqttClient.Publish(context.Background(), &paho.Publish{
	//		QoS:     0,
	//		Topic:   "sys/product/4LFCK4M0024Z3X/status",
	//		Payload: []byte("{\"tid\":\"cbcd8ce3-8636-f3e2-a400-e40049f070a4\",\"bid\":\"522f7bae-e47d-6d0e-a8f0-ae0a08127558\",\"timestamp\":1654620644746,\"method\":\"update_topo\",\"data\":{\"type\":119,\"sub_type\":0,\"device_secret\":\"289c5119bc5993d32997e26e34059ac7\",\"nonce\":\"5fe11aa247f709481fbb8351b458426f\",\"version\":1,\"sub_devices\":[{\"sn\":\"1581F5BKD225200BGP82\",\"type\":67,\"sub_type\":1,\"index\":\"A\",\"device_secret\":\"42a5c36f13ab9e825f960d0d23d7533b\",\"nonce\":\"53efda445346abe79d8713b6779998a4\",\"version\":1}]}}"),
	//	})
	//}()

	signalChanel := make(chan os.Signal, 1)
	signal.Notify(signalChanel,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	<-signalChanel
	log.Println("Received exit signal")
}
