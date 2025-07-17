// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrUpdateEventBridgeIntegrationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessKey(v string) *CreateOrUpdateEventBridgeIntegrationRequest
	GetAccessKey() *string
	SetAccessSecret(v string) *CreateOrUpdateEventBridgeIntegrationRequest
	GetAccessSecret() *string
	SetDescription(v string) *CreateOrUpdateEventBridgeIntegrationRequest
	GetDescription() *string
	SetEndpoint(v string) *CreateOrUpdateEventBridgeIntegrationRequest
	GetEndpoint() *string
	SetEventBusName(v string) *CreateOrUpdateEventBridgeIntegrationRequest
	GetEventBusName() *string
	SetEventBusRegionId(v string) *CreateOrUpdateEventBridgeIntegrationRequest
	GetEventBusRegionId() *string
	SetId(v int64) *CreateOrUpdateEventBridgeIntegrationRequest
	GetId() *int64
	SetName(v string) *CreateOrUpdateEventBridgeIntegrationRequest
	GetName() *string
	SetSource(v string) *CreateOrUpdateEventBridgeIntegrationRequest
	GetSource() *string
}

type CreateOrUpdateEventBridgeIntegrationRequest struct {
	// The AccessKey ID that is used to connect to EventBridge.
	//
	// example:
	//
	// abc******************
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	// The AccessKey secret that is used to connect to EventBridge.
	//
	// example:
	//
	// abc******************
	AccessSecret *string `json:"AccessSecret,omitempty" xml:"AccessSecret,omitempty"`
	// The description of the EventBridge integration.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The public endpoint of EventBridge.
	//
	// example:
	//
	// http://xxxxx
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The name of the event bus.
	//
	// example:
	//
	// EventBus_Test
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The region ID of the event bus.
	//
	// example:
	//
	// cn-hangzhou
	EventBusRegionId *string `json:"EventBusRegionId,omitempty" xml:"EventBusRegionId,omitempty"`
	// The ID of the EventBridge integration.
	//
	// 	- If you do not specify this parameter, an EventBridge integration is created.
	//
	// 	- If you specify this parameter, the specified EventBridge integration is modified.
	//
	// example:
	//
	// 123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the EventBridge integration.
	//
	// This parameter is required.
	//
	// example:
	//
	// EventBridge_Test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The event source.
	//
	// example:
	//
	// arms
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s CreateOrUpdateEventBridgeIntegrationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateEventBridgeIntegrationRequest) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateEventBridgeIntegrationRequest) GetAccessKey() *string {
	return s.AccessKey
}

func (s *CreateOrUpdateEventBridgeIntegrationRequest) GetAccessSecret() *string {
	return s.AccessSecret
}

func (s *CreateOrUpdateEventBridgeIntegrationRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateOrUpdateEventBridgeIntegrationRequest) GetEndpoint() *string {
	return s.Endpoint
}

func (s *CreateOrUpdateEventBridgeIntegrationRequest) GetEventBusName() *string {
	return s.EventBusName
}

func (s *CreateOrUpdateEventBridgeIntegrationRequest) GetEventBusRegionId() *string {
	return s.EventBusRegionId
}

func (s *CreateOrUpdateEventBridgeIntegrationRequest) GetId() *int64 {
	return s.Id
}

func (s *CreateOrUpdateEventBridgeIntegrationRequest) GetName() *string {
	return s.Name
}

func (s *CreateOrUpdateEventBridgeIntegrationRequest) GetSource() *string {
	return s.Source
}

func (s *CreateOrUpdateEventBridgeIntegrationRequest) SetAccessKey(v string) *CreateOrUpdateEventBridgeIntegrationRequest {
	s.AccessKey = &v
	return s
}

func (s *CreateOrUpdateEventBridgeIntegrationRequest) SetAccessSecret(v string) *CreateOrUpdateEventBridgeIntegrationRequest {
	s.AccessSecret = &v
	return s
}

func (s *CreateOrUpdateEventBridgeIntegrationRequest) SetDescription(v string) *CreateOrUpdateEventBridgeIntegrationRequest {
	s.Description = &v
	return s
}

func (s *CreateOrUpdateEventBridgeIntegrationRequest) SetEndpoint(v string) *CreateOrUpdateEventBridgeIntegrationRequest {
	s.Endpoint = &v
	return s
}

func (s *CreateOrUpdateEventBridgeIntegrationRequest) SetEventBusName(v string) *CreateOrUpdateEventBridgeIntegrationRequest {
	s.EventBusName = &v
	return s
}

func (s *CreateOrUpdateEventBridgeIntegrationRequest) SetEventBusRegionId(v string) *CreateOrUpdateEventBridgeIntegrationRequest {
	s.EventBusRegionId = &v
	return s
}

func (s *CreateOrUpdateEventBridgeIntegrationRequest) SetId(v int64) *CreateOrUpdateEventBridgeIntegrationRequest {
	s.Id = &v
	return s
}

func (s *CreateOrUpdateEventBridgeIntegrationRequest) SetName(v string) *CreateOrUpdateEventBridgeIntegrationRequest {
	s.Name = &v
	return s
}

func (s *CreateOrUpdateEventBridgeIntegrationRequest) SetSource(v string) *CreateOrUpdateEventBridgeIntegrationRequest {
	s.Source = &v
	return s
}

func (s *CreateOrUpdateEventBridgeIntegrationRequest) Validate() error {
	return dara.Validate(s)
}
