package main

import (
	"time"

	"github.com/olivere/elastic/v7"
	log "github.com/sirupsen/logrus"
	"gopkg.in/sohlich/elogrus.v7"
)

func main() {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://192.168.1.230:9200"))
	if err != nil {
		log.Panic(err)
	}
	hook, err := elogrus.NewAsyncElasticHook(client, "localhost", log.DebugLevel, "mylog")
	if err != nil {
		log.Panic(err)
	}
	log.AddHook(hook)

	log.WithFields(log.Fields{
		"name": "joe",
		"age":  42,
	}).Error("Hello logrus!")
	time.Sleep(time.Second)
}
