// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeliveryOption interface {
	dara.Model
	String() string
	GoString() string
	SetConcurrency(v int64) *DeliveryOption
	GetConcurrency() *int64
	SetEventSchema(v string) *DeliveryOption
	GetEventSchema() *string
}

type DeliveryOption struct {
	// The maximum number of concurrent events that can be delivered by the upstream event source to Function Compute. This parameter is valid only when ApsaraMQ for Kafka is used as the event source.
	//
	// example:
	//
	// 2
	Concurrency *int64 `json:"concurrency,omitempty" xml:"concurrency,omitempty"`
	// The format of each data element in the event parameter of the function. CloudEvents: describes event data in a common format, including event description and event payload data. CloudEvents is designed to simplify event declaration and transmission between different services and platforms. This is the default value. RawData: delivers only the event payload data and does not include other metadata information in the CloudEvents format.
	//
	// example:
	//
	// RawData
	EventSchema *string `json:"eventSchema,omitempty" xml:"eventSchema,omitempty"`
}

func (s DeliveryOption) String() string {
	return dara.Prettify(s)
}

func (s DeliveryOption) GoString() string {
	return s.String()
}

func (s *DeliveryOption) GetConcurrency() *int64 {
	return s.Concurrency
}

func (s *DeliveryOption) GetEventSchema() *string {
	return s.EventSchema
}

func (s *DeliveryOption) SetConcurrency(v int64) *DeliveryOption {
	s.Concurrency = &v
	return s
}

func (s *DeliveryOption) SetEventSchema(v string) *DeliveryOption {
	s.EventSchema = &v
	return s
}

func (s *DeliveryOption) Validate() error {
	return dara.Validate(s)
}
