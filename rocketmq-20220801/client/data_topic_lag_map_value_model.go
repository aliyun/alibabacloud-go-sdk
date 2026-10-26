// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataTopicLagMapValue interface {
	dara.Model
	String() string
	GoString() string
	SetReadyCount(v int64) *DataTopicLagMapValue
	GetReadyCount() *int64
	SetInflightCount(v int64) *DataTopicLagMapValue
	GetInflightCount() *int64
	SetDeliveryDuration(v int64) *DataTopicLagMapValue
	GetDeliveryDuration() *int64
	SetLastConsumeTimestamp(v int64) *DataTopicLagMapValue
	GetLastConsumeTimestamp() *int64
}

type DataTopicLagMapValue struct {
	// The number of ready messages.
	//
	// example:
	//
	// 1
	ReadyCount *int64 `json:"readyCount,omitempty" xml:"readyCount,omitempty"`
	// The number of inflight messages.
	//
	// example:
	//
	// 1
	InflightCount *int64 `json:"inflightCount,omitempty" xml:"inflightCount,omitempty"`
	// The delivery latency. Unit: seconds.
	//
	// example:
	//
	// 12
	DeliveryDuration *int64 `json:"deliveryDuration,omitempty" xml:"deliveryDuration,omitempty"`
	// The timestamp of the last consumption.
	//
	// example:
	//
	// 1735629607846
	LastConsumeTimestamp *int64 `json:"lastConsumeTimestamp,omitempty" xml:"lastConsumeTimestamp,omitempty"`
}

func (s DataTopicLagMapValue) String() string {
	return dara.Prettify(s)
}

func (s DataTopicLagMapValue) GoString() string {
	return s.String()
}

func (s *DataTopicLagMapValue) GetReadyCount() *int64 {
	return s.ReadyCount
}

func (s *DataTopicLagMapValue) GetInflightCount() *int64 {
	return s.InflightCount
}

func (s *DataTopicLagMapValue) GetDeliveryDuration() *int64 {
	return s.DeliveryDuration
}

func (s *DataTopicLagMapValue) GetLastConsumeTimestamp() *int64 {
	return s.LastConsumeTimestamp
}

func (s *DataTopicLagMapValue) SetReadyCount(v int64) *DataTopicLagMapValue {
	s.ReadyCount = &v
	return s
}

func (s *DataTopicLagMapValue) SetInflightCount(v int64) *DataTopicLagMapValue {
	s.InflightCount = &v
	return s
}

func (s *DataTopicLagMapValue) SetDeliveryDuration(v int64) *DataTopicLagMapValue {
	s.DeliveryDuration = &v
	return s
}

func (s *DataTopicLagMapValue) SetLastConsumeTimestamp(v int64) *DataTopicLagMapValue {
	s.LastConsumeTimestamp = &v
	return s
}

func (s *DataTopicLagMapValue) Validate() error {
	return dara.Validate(s)
}
