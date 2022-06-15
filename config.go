package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/eclipse/paho.golang/paho"
	"net"
	"time"
)

type Config struct {
	MQTTConfig   string `toml:"mqtt_config"`
	MQTTClientID string `toml:"mqtt_client_id"`
	MQTTUsername string `toml:"mqtt_username"`
	MQTTPassword string `toml:"mqtt_password"`
}

func (c *Config) NewMQTTClient() (*paho.Client, error) {
	conn, err := net.Dial("tcp", c.MQTTConfig)
	if err != nil {
		return nil, err
	}

	client := paho.NewClient(paho.ClientConfig{
		Conn: conn,
	})

	return client, nil
}

func (c *Config) ConnectMQTT(client *paho.Client) error {

	cp := &paho.Connect{
		KeepAlive:  30,
		ClientID:   c.MQTTClientID,
		CleanStart: true,
		Username:   c.MQTTUsername,
		Password:   []byte(c.MQTTPassword),
	}
	if c.MQTTUsername != "" {
		cp.UsernameFlag = true
	}
	if c.MQTTPassword != "" {
		cp.PasswordFlag = true
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	ca, err := client.Connect(ctx, cp)
	if err != nil {
		return err
	}
	if ca.ReasonCode != 0 {
		return errors.New(fmt.Sprintf("Failed to connect to %s : %d - %s\n", c.MQTTConfig, ca.ReasonCode, ca.Properties.ReasonString))
	}

	return nil
}
