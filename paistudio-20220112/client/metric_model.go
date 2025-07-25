// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMetric interface {
	dara.Model
	String() string
	GoString() string
	SetTime(v int64) *Metric
	GetTime() *int64
	SetValue(v string) *Metric
	GetValue() *string
}

type Metric struct {
	// example:
	//
	// rg17tmvwiokhzaxg
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
	// example:
	//
	// 23000
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s Metric) String() string {
	return dara.Prettify(s)
}

func (s Metric) GoString() string {
	return s.String()
}

func (s *Metric) GetTime() *int64 {
	return s.Time
}

func (s *Metric) GetValue() *string {
	return s.Value
}

func (s *Metric) SetTime(v int64) *Metric {
	s.Time = &v
	return s
}

func (s *Metric) SetValue(v string) *Metric {
	s.Value = &v
	return s
}

func (s *Metric) Validate() error {
	return dara.Validate(s)
}
