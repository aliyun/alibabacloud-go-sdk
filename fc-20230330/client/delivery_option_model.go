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
	// example:
	//
	// 2
	Concurrency *int64 `json:"concurrency,omitempty" xml:"concurrency,omitempty"`
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
