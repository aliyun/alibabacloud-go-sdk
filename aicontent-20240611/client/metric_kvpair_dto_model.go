// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMetricKVPairDTO interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *MetricKVPairDTO
	GetKey() *string
	SetValue(v float32) *MetricKVPairDTO
	GetValue() *float32
}

type MetricKVPairDTO struct {
	// example:
	//
	// total_calls
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// 100
	Value *float32 `json:"value,omitempty" xml:"value,omitempty"`
}

func (s MetricKVPairDTO) String() string {
	return dara.Prettify(s)
}

func (s MetricKVPairDTO) GoString() string {
	return s.String()
}

func (s *MetricKVPairDTO) GetKey() *string {
	return s.Key
}

func (s *MetricKVPairDTO) GetValue() *float32 {
	return s.Value
}

func (s *MetricKVPairDTO) SetKey(v string) *MetricKVPairDTO {
	s.Key = &v
	return s
}

func (s *MetricKVPairDTO) SetValue(v float32) *MetricKVPairDTO {
	s.Value = &v
	return s
}

func (s *MetricKVPairDTO) Validate() error {
	return dara.Validate(s)
}
