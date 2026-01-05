// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMetric interface {
	dara.Model
	String() string
	GoString() string
	SetTimestamp(v string) *Metric
	GetTimestamp() *string
	SetValue(v float64) *Metric
	GetValue() *float64
}

type Metric struct {
	// example:
	//
	// 1655897743
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// example:
	//
	// 98.35
	Value *float64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s Metric) String() string {
	return dara.Prettify(s)
}

func (s Metric) GoString() string {
	return s.String()
}

func (s *Metric) GetTimestamp() *string {
	return s.Timestamp
}

func (s *Metric) GetValue() *float64 {
	return s.Value
}

func (s *Metric) SetTimestamp(v string) *Metric {
	s.Timestamp = &v
	return s
}

func (s *Metric) SetValue(v float64) *Metric {
	s.Value = &v
	return s
}

func (s *Metric) Validate() error {
	return dara.Validate(s)
}
