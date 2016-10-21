package hnsq

import (
	"github.com/nsqio/go-nsq"
)

type Engine struct {
	consumers []*nsq.Consumer
	cfg *nsq.Config
}

func New(cfg *nsq.Config) *Engine {
	engine := new(Engine)
	engine.cfg = cfg
	return engine
}

func (e *Engine) Handler(topic, channel string, handler nsq.Handler) {
	c, err := nsq.NewConsumer(topic, channel, e.cfg)
	if err != nil {
		panic(err)
		return
	}

	c.AddHandler(handler)

	e.consumers = append(e.consumers, c)
}

func (e *Engine) Run(addr string) {
	for i, _ := range e.consumers {
		if err := e.consumers[i].ConnectToNSQD(addr); err != nil {
			panic(err)
			return
		}
	}

	for {
		select {
		}
	}
}
