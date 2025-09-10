// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEventSourceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateEventSourceShrinkRequest
	GetDescription() *string
	SetEventBusName(v string) *UpdateEventSourceShrinkRequest
	GetEventBusName() *string
	SetEventSourceName(v string) *UpdateEventSourceShrinkRequest
	GetEventSourceName() *string
	SetExternalSourceConfigShrink(v string) *UpdateEventSourceShrinkRequest
	GetExternalSourceConfigShrink() *string
	SetExternalSourceType(v string) *UpdateEventSourceShrinkRequest
	GetExternalSourceType() *string
	SetLinkedExternalSource(v bool) *UpdateEventSourceShrinkRequest
	GetLinkedExternalSource() *bool
	SetSourceHttpEventParametersShrink(v string) *UpdateEventSourceShrinkRequest
	GetSourceHttpEventParametersShrink() *string
	SetSourceKafkaParametersShrink(v string) *UpdateEventSourceShrinkRequest
	GetSourceKafkaParametersShrink() *string
	SetSourceMNSParametersShrink(v string) *UpdateEventSourceShrinkRequest
	GetSourceMNSParametersShrink() *string
	SetSourceOSSEventParametersShrink(v string) *UpdateEventSourceShrinkRequest
	GetSourceOSSEventParametersShrink() *string
	SetSourceRabbitMQParametersShrink(v string) *UpdateEventSourceShrinkRequest
	GetSourceRabbitMQParametersShrink() *string
	SetSourceRocketMQParametersShrink(v string) *UpdateEventSourceShrinkRequest
	GetSourceRocketMQParametersShrink() *string
	SetSourceSLSParametersShrink(v string) *UpdateEventSourceShrinkRequest
	GetSourceSLSParametersShrink() *string
	SetSourceScheduledEventParametersShrink(v string) *UpdateEventSourceShrinkRequest
	GetSourceScheduledEventParametersShrink() *string
}

type UpdateEventSourceShrinkRequest struct {
	// The description of the event source.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The event bus with which the event source is associated.
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
	// myrabbitmq.source
	EventSourceName *string `json:"EventSourceName,omitempty" xml:"EventSourceName,omitempty"`
	// The configurations of the external data source.
	//
	// example:
	//
	// {\\"ConsumePosition\\":\\"end\\",\\"LogStore\\":\\"oss_log\\",\\"Project\\":\\"slsaudit-center-5795350335281001-cn-beijing\\",\\"RoleName\\":\\"sls-beijing-tf\\"}
	ExternalSourceConfigShrink *string `json:"ExternalSourceConfig,omitempty" xml:"ExternalSourceConfig,omitempty"`
	// The type of the external data source.
	//
	// example:
	//
	// SLS
	ExternalSourceType *string `json:"ExternalSourceType,omitempty" xml:"ExternalSourceType,omitempty"`
	// Specifies whether to connect to an external data source.
	//
	// example:
	//
	// true
	LinkedExternalSource *bool `json:"LinkedExternalSource,omitempty" xml:"LinkedExternalSource,omitempty"`
	// The parameters that are configured if the event source is HTTP events.
	SourceHttpEventParametersShrink *string `json:"SourceHttpEventParameters,omitempty" xml:"SourceHttpEventParameters,omitempty"`
	// The parameters that are configured if the event source is Message Queue for Apache Kafka.
	SourceKafkaParametersShrink *string `json:"SourceKafkaParameters,omitempty" xml:"SourceKafkaParameters,omitempty"`
	// The parameters that are configured if the event source is Message Service (MNS).
	SourceMNSParametersShrink      *string `json:"SourceMNSParameters,omitempty" xml:"SourceMNSParameters,omitempty"`
	SourceOSSEventParametersShrink *string `json:"SourceOSSEventParameters,omitempty" xml:"SourceOSSEventParameters,omitempty"`
	// The parameters that are configured if the event source is Message Queue for RabbitMQ.
	SourceRabbitMQParametersShrink *string `json:"SourceRabbitMQParameters,omitempty" xml:"SourceRabbitMQParameters,omitempty"`
	// The parameters that are configured if the event source is Message Queue for Apache RocketMQ.
	SourceRocketMQParametersShrink *string `json:"SourceRocketMQParameters,omitempty" xml:"SourceRocketMQParameters,omitempty"`
	// SourceSLSParameters
	SourceSLSParametersShrink *string `json:"SourceSLSParameters,omitempty" xml:"SourceSLSParameters,omitempty"`
	// The parameters that are configured if you specify scheduled events as the event source.
	SourceScheduledEventParametersShrink *string `json:"SourceScheduledEventParameters,omitempty" xml:"SourceScheduledEventParameters,omitempty"`
}

func (s UpdateEventSourceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventSourceShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateEventSourceShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateEventSourceShrinkRequest) GetEventBusName() *string {
	return s.EventBusName
}

func (s *UpdateEventSourceShrinkRequest) GetEventSourceName() *string {
	return s.EventSourceName
}

func (s *UpdateEventSourceShrinkRequest) GetExternalSourceConfigShrink() *string {
	return s.ExternalSourceConfigShrink
}

func (s *UpdateEventSourceShrinkRequest) GetExternalSourceType() *string {
	return s.ExternalSourceType
}

func (s *UpdateEventSourceShrinkRequest) GetLinkedExternalSource() *bool {
	return s.LinkedExternalSource
}

func (s *UpdateEventSourceShrinkRequest) GetSourceHttpEventParametersShrink() *string {
	return s.SourceHttpEventParametersShrink
}

func (s *UpdateEventSourceShrinkRequest) GetSourceKafkaParametersShrink() *string {
	return s.SourceKafkaParametersShrink
}

func (s *UpdateEventSourceShrinkRequest) GetSourceMNSParametersShrink() *string {
	return s.SourceMNSParametersShrink
}

func (s *UpdateEventSourceShrinkRequest) GetSourceOSSEventParametersShrink() *string {
	return s.SourceOSSEventParametersShrink
}

func (s *UpdateEventSourceShrinkRequest) GetSourceRabbitMQParametersShrink() *string {
	return s.SourceRabbitMQParametersShrink
}

func (s *UpdateEventSourceShrinkRequest) GetSourceRocketMQParametersShrink() *string {
	return s.SourceRocketMQParametersShrink
}

func (s *UpdateEventSourceShrinkRequest) GetSourceSLSParametersShrink() *string {
	return s.SourceSLSParametersShrink
}

func (s *UpdateEventSourceShrinkRequest) GetSourceScheduledEventParametersShrink() *string {
	return s.SourceScheduledEventParametersShrink
}

func (s *UpdateEventSourceShrinkRequest) SetDescription(v string) *UpdateEventSourceShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateEventSourceShrinkRequest) SetEventBusName(v string) *UpdateEventSourceShrinkRequest {
	s.EventBusName = &v
	return s
}

func (s *UpdateEventSourceShrinkRequest) SetEventSourceName(v string) *UpdateEventSourceShrinkRequest {
	s.EventSourceName = &v
	return s
}

func (s *UpdateEventSourceShrinkRequest) SetExternalSourceConfigShrink(v string) *UpdateEventSourceShrinkRequest {
	s.ExternalSourceConfigShrink = &v
	return s
}

func (s *UpdateEventSourceShrinkRequest) SetExternalSourceType(v string) *UpdateEventSourceShrinkRequest {
	s.ExternalSourceType = &v
	return s
}

func (s *UpdateEventSourceShrinkRequest) SetLinkedExternalSource(v bool) *UpdateEventSourceShrinkRequest {
	s.LinkedExternalSource = &v
	return s
}

func (s *UpdateEventSourceShrinkRequest) SetSourceHttpEventParametersShrink(v string) *UpdateEventSourceShrinkRequest {
	s.SourceHttpEventParametersShrink = &v
	return s
}

func (s *UpdateEventSourceShrinkRequest) SetSourceKafkaParametersShrink(v string) *UpdateEventSourceShrinkRequest {
	s.SourceKafkaParametersShrink = &v
	return s
}

func (s *UpdateEventSourceShrinkRequest) SetSourceMNSParametersShrink(v string) *UpdateEventSourceShrinkRequest {
	s.SourceMNSParametersShrink = &v
	return s
}

func (s *UpdateEventSourceShrinkRequest) SetSourceOSSEventParametersShrink(v string) *UpdateEventSourceShrinkRequest {
	s.SourceOSSEventParametersShrink = &v
	return s
}

func (s *UpdateEventSourceShrinkRequest) SetSourceRabbitMQParametersShrink(v string) *UpdateEventSourceShrinkRequest {
	s.SourceRabbitMQParametersShrink = &v
	return s
}

func (s *UpdateEventSourceShrinkRequest) SetSourceRocketMQParametersShrink(v string) *UpdateEventSourceShrinkRequest {
	s.SourceRocketMQParametersShrink = &v
	return s
}

func (s *UpdateEventSourceShrinkRequest) SetSourceSLSParametersShrink(v string) *UpdateEventSourceShrinkRequest {
	s.SourceSLSParametersShrink = &v
	return s
}

func (s *UpdateEventSourceShrinkRequest) SetSourceScheduledEventParametersShrink(v string) *UpdateEventSourceShrinkRequest {
	s.SourceScheduledEventParametersShrink = &v
	return s
}

func (s *UpdateEventSourceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
