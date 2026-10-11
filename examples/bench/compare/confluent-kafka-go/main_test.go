package main

import (
	"testing"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func TestProducerConfigAccepted(t *testing.T) {
	cfg := producerConfig()
	// 不设置 broker 地址, 仅验证生产配置能被真实的 librdkafka 接受.
	(*cfg)["bootstrap.servers"] = ""
	p, err := kafka.NewProducer(cfg)
	if err != nil {
		t.Fatal(err)
	}
	p.Close()
}
