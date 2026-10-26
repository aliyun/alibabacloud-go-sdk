// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataLiteTopicLagMapValue interface {
	dara.Model
	String() string
	GoString() string
	SetReadyCount(v int64) *DataLiteTopicLagMapValue
	GetReadyCount() *int64
	SetDeliveryDuration(v int64) *DataLiteTopicLagMapValue
	GetDeliveryDuration() *int64
}

type DataLiteTopicLagMapValue struct {
	// The number of accumulated messages.
	//
	// example:
	//
	// 300
	ReadyCount *int64 `json:"readyCount,omitempty" xml:"readyCount,omitempty"`
	// The delivery latency. Unit: seconds.
	//
	// example:
	//
	// 30
	DeliveryDuration *int64 `json:"deliveryDuration,omitempty" xml:"deliveryDuration,omitempty"`
}

func (s DataLiteTopicLagMapValue) String() string {
	return dara.Prettify(s)
}

func (s DataLiteTopicLagMapValue) GoString() string {
	return s.String()
}

func (s *DataLiteTopicLagMapValue) GetReadyCount() *int64 {
	return s.ReadyCount
}

func (s *DataLiteTopicLagMapValue) GetDeliveryDuration() *int64 {
	return s.DeliveryDuration
}

func (s *DataLiteTopicLagMapValue) SetReadyCount(v int64) *DataLiteTopicLagMapValue {
	s.ReadyCount = &v
	return s
}

func (s *DataLiteTopicLagMapValue) SetDeliveryDuration(v int64) *DataLiteTopicLagMapValue {
	s.DeliveryDuration = &v
	return s
}

func (s *DataLiteTopicLagMapValue) Validate() error {
	return dara.Validate(s)
}
