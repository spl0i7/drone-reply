package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/eclipse/paho.golang/paho"
	"log"
	"strings"
	"time"
)

type MQTTMessageHandler struct {
	mqttClient *paho.Client
}

func NewMessageHandler(client *paho.Client) *MQTTMessageHandler {
	return &MQTTMessageHandler{
		mqttClient: client,
	}
}

func (p *MQTTMessageHandler) FlightStatusHandler(m *paho.Publish) {

	var message MQTTFlightStatus
	if err := json.Unmarshal(m.Payload, &message); err != nil {
		log.Println(err)
		return
	}

	response := MQTTFlightStatusResponse{
		Tid: message.Tid,
		Bid: message.Bid,
		Data: struct {
			Result int `json:"result"`
		}{
			Result: 0,
		},
	}

	pid := m.Topic
	splitTopic := strings.Split(m.Topic, "/")
	if len(splitTopic) == 4 {
		pid = splitTopic[2]
	}
	responseTopic := fmt.Sprintf("sys/product/%s/status_reply", pid)

	responseBytes, err := json.Marshal(&response)
	if err != nil {
		log.Println(err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if _, err := p.mqttClient.Publish(ctx, &paho.Publish{
		Topic:   responseTopic,
		Payload: responseBytes,
	}); err != nil {
		log.Println(err)
	}
}

func (p *MQTTMessageHandler) DebugMessage(m *paho.Publish) {
	fmt.Println("Topic", m.Topic)
	fmt.Println("Payload", string(m.Payload))
}
