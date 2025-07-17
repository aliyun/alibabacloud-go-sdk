// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrUpdateEventBridgeIntegrationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEventBridgeIntegration(v *CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration) *CreateOrUpdateEventBridgeIntegrationResponseBody
	GetEventBridgeIntegration() *CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration
	SetRequestId(v string) *CreateOrUpdateEventBridgeIntegrationResponseBody
	GetRequestId() *string
}

type CreateOrUpdateEventBridgeIntegrationResponseBody struct {
	// The information about the EventBridge integration.
	EventBridgeIntegration *CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration `json:"EventBridgeIntegration,omitempty" xml:"EventBridgeIntegration,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 2B289756-E791-5842-BCBD-AD414C******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateOrUpdateEventBridgeIntegrationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateEventBridgeIntegrationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateEventBridgeIntegrationResponseBody) GetEventBridgeIntegration() *CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration {
	return s.EventBridgeIntegration
}

func (s *CreateOrUpdateEventBridgeIntegrationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOrUpdateEventBridgeIntegrationResponseBody) SetEventBridgeIntegration(v *CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration) *CreateOrUpdateEventBridgeIntegrationResponseBody {
	s.EventBridgeIntegration = v
	return s
}

func (s *CreateOrUpdateEventBridgeIntegrationResponseBody) SetRequestId(v string) *CreateOrUpdateEventBridgeIntegrationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOrUpdateEventBridgeIntegrationResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration struct {
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
	// example:
	//
	// 123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the EventBridge integration.
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

func (s CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration) GetAccessKey() *string {
	return s.AccessKey
}

func (s *CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration) GetAccessSecret() *string {
	return s.AccessSecret
}

func (s *CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration) GetDescription() *string {
	return s.Description
}

func (s *CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration) GetEndpoint() *string {
	return s.Endpoint
}

func (s *CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration) GetEventBusName() *string {
	return s.EventBusName
}

func (s *CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration) GetEventBusRegionId() *string {
	return s.EventBusRegionId
}

func (s *CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration) GetId() *int64 {
	return s.Id
}

func (s *CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration) GetName() *string {
	return s.Name
}

func (s *CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration) GetSource() *string {
	return s.Source
}

func (s *CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration) SetAccessKey(v string) *CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration {
	s.AccessKey = &v
	return s
}

func (s *CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration) SetAccessSecret(v string) *CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration {
	s.AccessSecret = &v
	return s
}

func (s *CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration) SetDescription(v string) *CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration {
	s.Description = &v
	return s
}

func (s *CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration) SetEndpoint(v string) *CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration {
	s.Endpoint = &v
	return s
}

func (s *CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration) SetEventBusName(v string) *CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration {
	s.EventBusName = &v
	return s
}

func (s *CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration) SetEventBusRegionId(v string) *CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration {
	s.EventBusRegionId = &v
	return s
}

func (s *CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration) SetId(v int64) *CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration {
	s.Id = &v
	return s
}

func (s *CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration) SetName(v string) *CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration {
	s.Name = &v
	return s
}

func (s *CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration) SetSource(v string) *CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration {
	s.Source = &v
	return s
}

func (s *CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration) Validate() error {
	return dara.Validate(s)
}
