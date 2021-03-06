package validators

import (
	"time"
)

// Request is sent to the validation topic for each new payload
type Request struct {
	Account     string    `json:"account"`
	Category    string    `json:"category"`
	Metadata    Metadata  `json:"metadata"`
	RequestID   string    `json:"request_id"`
	Principal   string    `json:"principal"`
	Service     string    `json:"service"`
	Size        int64     `json:"size"`
	URL         string    `json:"url"`
	ID          string    `json:"id,omitempty"`
	B64Identity string    `json:"b64_identity"`
	Timestamp   time.Time `json:"timestamp"`
}

// Response is returned by validators and sent via the announcement
type Response struct {
	Account          string                 `json:"account"`
	Validation       string                 `json:"validation"`
	RequestID        string                 `json:"request_id"`
	Principal        string                 `json:"principal"`
	Service          string                 `json:"service"`
	URL              string                 `json:"url"`
	B64Identity      string                 `json:"b64_identity"`
	ID               string                 `json:"id,omitempty"`
	SatelliteManaged *bool                  `json:"satellite_managed,omitempty"`
	Extras           map[string]interface{} `json:"extras"`
	Timestamp        time.Time              `json:"timestamp"`
}

// Status is the message sent to the payload tracker
type Status struct {
	Service     string    `json:"service"`
	Source      string    `json:"source,omitempty"`
	Account     string    `json:"account"`
	RequestID   string    `json:"request_id"`
	InventoryID string    `json:"inventory_id,omitempty"`
	SystemID    string    `json:"system_id,omitempty"`
	Status      string    `json:"status"`
	StatusMsg   string    `json:"status_msg"`
	Date        time.Time `json:"date"`
}

// Metadata is the expected data from a client
type Metadata struct {
	IPAddresses  []string `json:"ip_addresses,omitempty"`
	Account      string   `json:"account,omitempty"`
	InsightsID   string   `json:"insights_id,omitempty"`
	MachineID    string   `json:"machine_id,omitempty"`
	SubManID     string   `json:"subscription_manager_id,omitempty"`
	MacAddresses []string `json:"mac_addresses,omitempty"`
	FQDN         string   `json:"fqdn,omitempty"`
}

// ServiceDescriptor is used to select a message topic
type ServiceDescriptor struct {
	Service  string
	Category string
}

// Validator validates requests
type Validator interface {
	Validate(req *Request)
	ValidateService(service *ServiceDescriptor) error
}
