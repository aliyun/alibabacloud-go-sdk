// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEventSourceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateEventSourceShrinkRequest
	GetDescription() *string
	SetEventBusName(v string) *CreateEventSourceShrinkRequest
	GetEventBusName() *string
	SetEventSourceName(v string) *CreateEventSourceShrinkRequest
	GetEventSourceName() *string
	SetExternalSourceConfigShrink(v string) *CreateEventSourceShrinkRequest
	GetExternalSourceConfigShrink() *string
	SetExternalSourceType(v []byte) *CreateEventSourceShrinkRequest
	GetExternalSourceType() []byte
	SetLinkedExternalSource(v bool) *CreateEventSourceShrinkRequest
	GetLinkedExternalSource() *bool
	SetSourceHttpEventParametersShrink(v string) *CreateEventSourceShrinkRequest
	GetSourceHttpEventParametersShrink() *string
	SetSourceKafkaParametersShrink(v string) *CreateEventSourceShrinkRequest
	GetSourceKafkaParametersShrink() *string
	SetSourceMNSParametersShrink(v string) *CreateEventSourceShrinkRequest
	GetSourceMNSParametersShrink() *string
	SetSourceOSSEventParametersShrink(v string) *CreateEventSourceShrinkRequest
	GetSourceOSSEventParametersShrink() *string
	SetSourceRabbitMQParametersShrink(v string) *CreateEventSourceShrinkRequest
	GetSourceRabbitMQParametersShrink() *string
	SetSourceRocketMQParametersShrink(v string) *CreateEventSourceShrinkRequest
	GetSourceRocketMQParametersShrink() *string
	SetSourceSLSParametersShrink(v string) *CreateEventSourceShrinkRequest
	GetSourceSLSParametersShrink() *string
	SetSourceScheduledEventParametersShrink(v string) *CreateEventSourceShrinkRequest
	GetSourceScheduledEventParametersShrink() *string
}

type CreateEventSourceShrinkRequest struct {
	// The description of the event source.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the event bus with which the event source is associated.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-event-bus
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The name of the event source.
	//
	// This parameter is required.
	//
	// example:
	//
	// myrabbitmq.sourc
	EventSourceName *string `json:"EventSourceName,omitempty" xml:"EventSourceName,omitempty"`
	// The configurations of the external data source.
	ExternalSourceConfigShrink *string `json:"ExternalSourceConfig,omitempty" xml:"ExternalSourceConfig,omitempty"`
	// The type of the external data source.
	//
	// example:
	//
	// RabbitMQ
	ExternalSourceType []byte `json:"ExternalSourceType,omitempty" xml:"ExternalSourceType,omitempty"`
	// Specify whether to connect to an external data source.
	//
	// example:
	//
	// true
	LinkedExternalSource *bool `json:"LinkedExternalSource,omitempty" xml:"LinkedExternalSource,omitempty"`
	// The parameters that are configured if the event source is HTTP events.
	SourceHttpEventParametersShrink *string `json:"SourceHttpEventParameters,omitempty" xml:"SourceHttpEventParameters,omitempty"`
	// The parameters that are configured if the event source is Message Queue for Apache Kafka.
	SourceKafkaParametersShrink *string `json:"SourceKafkaParameters,omitempty" xml:"SourceKafkaParameters,omitempty"`
	// The parameters that are configured if you specify Simple Message Queue (formerly MNS) (SMQ) as the event source. If you specify SMQ as the event source, you must configure RegionId, IsBase64Decode, and QueueName.
	SourceMNSParametersShrink      *string `json:"SourceMNSParameters,omitempty" xml:"SourceMNSParameters,omitempty"`
	SourceOSSEventParametersShrink *string `json:"SourceOSSEventParameters,omitempty" xml:"SourceOSSEventParameters,omitempty"`
	// The parameters that are configured if the event source is Message Queue for RabbitMQ.
	SourceRabbitMQParametersShrink *string `json:"SourceRabbitMQParameters,omitempty" xml:"SourceRabbitMQParameters,omitempty"`
	// The parameters that are configured if the event source is Message Queue for Apache RocketMQ.
	SourceRocketMQParametersShrink *string `json:"SourceRocketMQParameters,omitempty" xml:"SourceRocketMQParameters,omitempty"`
	// The parameters that are configured if the event source is Log Service.
	SourceSLSParametersShrink *string `json:"SourceSLSParameters,omitempty" xml:"SourceSLSParameters,omitempty"`
	// The parameters that are configured if you specify scheduled events as the event source.
	SourceScheduledEventParametersShrink *string `json:"SourceScheduledEventParameters,omitempty" xml:"SourceScheduledEventParameters,omitempty"`
}

func (s CreateEventSourceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEventSourceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateEventSourceShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateEventSourceShrinkRequest) GetEventBusName() *string {
	return s.EventBusName
}

func (s *CreateEventSourceShrinkRequest) GetEventSourceName() *string {
	return s.EventSourceName
}

func (s *CreateEventSourceShrinkRequest) GetExternalSourceConfigShrink() *string {
	return s.ExternalSourceConfigShrink
}

func (s *CreateEventSourceShrinkRequest) GetExternalSourceType() []byte {
	return s.ExternalSourceType
}

func (s *CreateEventSourceShrinkRequest) GetLinkedExternalSource() *bool {
	return s.LinkedExternalSource
}

func (s *CreateEventSourceShrinkRequest) GetSourceHttpEventParametersShrink() *string {
	return s.SourceHttpEventParametersShrink
}

func (s *CreateEventSourceShrinkRequest) GetSourceKafkaParametersShrink() *string {
	return s.SourceKafkaParametersShrink
}

func (s *CreateEventSourceShrinkRequest) GetSourceMNSParametersShrink() *string {
	return s.SourceMNSParametersShrink
}

func (s *CreateEventSourceShrinkRequest) GetSourceOSSEventParametersShrink() *string {
	return s.SourceOSSEventParametersShrink
}

func (s *CreateEventSourceShrinkRequest) GetSourceRabbitMQParametersShrink() *string {
	return s.SourceRabbitMQParametersShrink
}

func (s *CreateEventSourceShrinkRequest) GetSourceRocketMQParametersShrink() *string {
	return s.SourceRocketMQParametersShrink
}

func (s *CreateEventSourceShrinkRequest) GetSourceSLSParametersShrink() *string {
	return s.SourceSLSParametersShrink
}

func (s *CreateEventSourceShrinkRequest) GetSourceScheduledEventParametersShrink() *string {
	return s.SourceScheduledEventParametersShrink
}

func (s *CreateEventSourceShrinkRequest) SetDescription(v string) *CreateEventSourceShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateEventSourceShrinkRequest) SetEventBusName(v string) *CreateEventSourceShrinkRequest {
	s.EventBusName = &v
	return s
}

func (s *CreateEventSourceShrinkRequest) SetEventSourceName(v string) *CreateEventSourceShrinkRequest {
	s.EventSourceName = &v
	return s
}

func (s *CreateEventSourceShrinkRequest) SetExternalSourceConfigShrink(v string) *CreateEventSourceShrinkRequest {
	s.ExternalSourceConfigShrink = &v
	return s
}

func (s *CreateEventSourceShrinkRequest) SetExternalSourceType(v []byte) *CreateEventSourceShrinkRequest {
	s.ExternalSourceType = v
	return s
}

func (s *CreateEventSourceShrinkRequest) SetLinkedExternalSource(v bool) *CreateEventSourceShrinkRequest {
	s.LinkedExternalSource = &v
	return s
}

func (s *CreateEventSourceShrinkRequest) SetSourceHttpEventParametersShrink(v string) *CreateEventSourceShrinkRequest {
	s.SourceHttpEventParametersShrink = &v
	return s
}

func (s *CreateEventSourceShrinkRequest) SetSourceKafkaParametersShrink(v string) *CreateEventSourceShrinkRequest {
	s.SourceKafkaParametersShrink = &v
	return s
}

func (s *CreateEventSourceShrinkRequest) SetSourceMNSParametersShrink(v string) *CreateEventSourceShrinkRequest {
	s.SourceMNSParametersShrink = &v
	return s
}

func (s *CreateEventSourceShrinkRequest) SetSourceOSSEventParametersShrink(v string) *CreateEventSourceShrinkRequest {
	s.SourceOSSEventParametersShrink = &v
	return s
}

func (s *CreateEventSourceShrinkRequest) SetSourceRabbitMQParametersShrink(v string) *CreateEventSourceShrinkRequest {
	s.SourceRabbitMQParametersShrink = &v
	return s
}

func (s *CreateEventSourceShrinkRequest) SetSourceRocketMQParametersShrink(v string) *CreateEventSourceShrinkRequest {
	s.SourceRocketMQParametersShrink = &v
	return s
}

func (s *CreateEventSourceShrinkRequest) SetSourceSLSParametersShrink(v string) *CreateEventSourceShrinkRequest {
	s.SourceSLSParametersShrink = &v
	return s
}

func (s *CreateEventSourceShrinkRequest) SetSourceScheduledEventParametersShrink(v string) *CreateEventSourceShrinkRequest {
	s.SourceScheduledEventParametersShrink = &v
	return s
}

func (s *CreateEventSourceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
