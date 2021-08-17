package modules

import (
	"context"
	"time"

	"github.com/erda-project/erda-infra/base/logs"
	"github.com/erda-project/erda-infra/base/servicehub"
)

type config struct {
	// some fields of config for this provider
	Message string `file:"message" flag:"msg" default:"hi" desc:"message to print"` 
}

// +provider
type provider struct {
	Cfg *config
	Log logs.Logger
}

// Run this is optional
func (p *provider) Init(ctx servicehub.Context) error {
	p.Log.Info("message: ", p.Cfg.Message)
	return nil
}

// Run this is optional
func (p *provider) Run(ctx context.Context) error {
	tick := time.NewTicker(3 * time.Second)
	defer tick.Stop()
	for {
		select {
		case <-tick.C:
			p.Log.Info("do something...")
		case <-ctx.Done():
			return nil
		}
	}
}

func init() {
	servicehub.Register("erda.collector.exporter.kafka", &servicehub.Spec{
		Services:    []string{
			"erda.collector.exporter.kafka-service",
		},
		Description: "here is description of erda.collector.exporter.kafka",
		ConfigFunc: func() interface{} {
			return &config{}
		},
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
