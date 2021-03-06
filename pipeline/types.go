package pipeline

import (
	"github.com/redhatinsights/insights-ingress-go/announcers"
	"github.com/redhatinsights/insights-ingress-go/interactions/inventory"
	"github.com/redhatinsights/insights-ingress-go/stage"
	"github.com/redhatinsights/insights-ingress-go/validators"
)

// Pipeline defines the descrete processing steps for ingress
type Pipeline struct {
	Stager      stage.Stager
	Validator   validators.Validator
	Announcer   announcers.Announcer
	ValidChan   chan *validators.Response
	InvalidChan chan *validators.Response
	Inventory   inventory.Inventory
	Tracker     announcers.Announcer
}
