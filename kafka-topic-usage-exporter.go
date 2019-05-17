package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/Shopify/sarama"
	"github.com/spf13/viper"
)

func main() {
	// command line flags
	verbose := flag.Bool("verbose", false, "display output to stdio")
	configName := flag.String("configname", "config", "location of config file")
	configPath := flag.String("config", ".", "location of config file")
	flag.Parse()
	viper.SetConfigType("yaml")
	viper.SetConfigName(*configName)
	viper.AddConfigPath(*configPath)

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	dataDirs := viper.GetStringSlice("data_dirs")
	var promFile = viper.GetString("prom_file")
	var cluster = viper.GetString("cluster")
	var delay = viper.GetInt("delay")
	brokers := viper.GetStringSlice("brokers")

	for {
		topics := GetKafkaTopics(brokers)

		hostname, err := os.Hostname()
		if err != nil {
			panic(err)
		}
		var lines []string
		for _, logDir := range dataDirs {

			dirListing, err := ioutil.ReadDir(logDir)
			if err != nil {
				log.Fatal(err)
			}

			for _, topic := range topics {
				var topicSize int64
				topicSize = 0
				for _, file := range dirListing {
					topic := topic + "-"
					if strings.HasPrefix(file.Name(), topic) {
						dirname := path.Join(logDir, file.Name())
						topicSize = topicSize + GetDirSizeBytes(dirname)
					}
				}
				var message = fmt.Sprintf("kafka_topic_disk_usage_bytes{kafka_log_dir=\"%v\", kafka_topic=\"%v\", kafka_node=\"%v\", kafka_cluster=\"%v\"} %v\n", logDir, topic, hostname, cluster, topicSize)
				if *verbose {
					fmt.Printf(message)
				}
				lines = append(lines, message)
			}

		}
		if len(lines) > 0 {
			err = ioutil.WriteFile(promFile, []byte(strings.Join(lines, "")), 0644)
		}
		time.Sleep(time.Duration(delay) * time.Second)
	}
}

/*
 get the directory size in bytes
*/
func GetDirSizeBytes(path string) int64 {

	var dirSize int64

	readSize := func(path string, file os.FileInfo, err error) error {
		if !file.IsDir() {
			dirSize += file.Size()
		}

		return nil
	}

	filepath.Walk(path, readSize)
	return dirSize
}

func GetKafkaTopics(brokers []string) []string {
	var topics []string
	kafkaClient := sarama.NewConfig()
	kafkaClient.Consumer.Return.Errors = true
	// setup consumer
	consumer, err := sarama.NewConsumer(brokers, kafkaClient)

	if err == nil {
		topics, _ = consumer.Topics()
		consumer.Close()
	}

	return topics
}
