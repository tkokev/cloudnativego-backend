package configuration

import (
	"encoding/json"
	"os"

	"bitbucket.org/minamartinteam/myevents/src/lib/persistence/dblayer"
	"strings"
)

var (
	DBTypeDefault              = dblayer.DBTYPE("mongodb")
	DBConnectionDefault        = "mongodb://127.0.0.1"
	RestfulEPDefault           = "localhost:8181"
	MessageBrokerTypeDefault   = "amqp"
	AMQPMessageBrokerDefault   = "amqp://guest:guest@localhost:5672"
	KafkaMessageBrokersDefault = []string{"localhost:9092"}
)

type EventServiceConfig struct {
	Databasetype        dblayer.DBTYPE `json:"databasetype"`
	DBConnection        string         `json:"dbconnection"`
	RestfulEndpoint     string         `json:"restfulapi_endpoint"`
	MessageBrokerType   string         `json:"message_broker_type"`
	AMQPMessageBroker   string         `json:"amqp_message_broker"`
	KafkaMessageBrokers []string       `json:"kafka_message_brokers"`
}

func ExtractConfiguration(filename string) EventServiceConfig {
	conf := EventServiceConfig{
		DBTypeDefault,
		DBConnectionDefault,
		RestfulEPDefault,
		MessageBrokerTypeDefault,
		AMQPMessageBrokerDefault,
		KafkaMessageBrokersDefault,
	}

	file, err := os.Open(filename)
	if err != nil {
		return conf
	}

	json.NewDecoder(file).Decode(&conf)

	if v := os.Getenv("MONGO_URL"); v != "" {
		conf.Databasetype = "mongodb"
		conf.DBConnection = v
	}

	if v := os.Getenv("AMQP_BROKER_URL"); v != "" {
		conf.MessageBrokerType = "amqp"
		conf.AMQPMessageBroker = v
	} else if v := os.Getenv("KAFKA_BROKER_URLS"); v != "" {
		conf.MessageBrokerType = "kafka"
		conf.KafkaMessageBrokers = strings.Split(v, ",")
	}

	return conf
}
