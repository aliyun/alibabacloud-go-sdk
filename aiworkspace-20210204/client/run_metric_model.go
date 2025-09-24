// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunMetric interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *RunMetric
	GetKey() *string
	SetStep(v int64) *RunMetric
	GetStep() *int64
	SetTimestamp(v int64) *RunMetric
	GetTimestamp() *int64
	SetValue(v float32) *RunMetric
	GetValue() *float32
}

type RunMetric struct {
	// This parameter is required.
	Key       *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Step      *int64  `json:"Step,omitempty" xml:"Step,omitempty"`
	Timestamp *int64  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// This parameter is required.
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s RunMetric) String() string {
	return dara.Prettify(s)
}

func (s RunMetric) GoString() string {
	return s.String()
}

func (s *RunMetric) GetKey() *string {
	return s.Key
}

func (s *RunMetric) GetStep() *int64 {
	return s.Step
}

func (s *RunMetric) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *RunMetric) GetValue() *float32 {
	return s.Value
}

func (s *RunMetric) SetKey(v string) *RunMetric {
	s.Key = &v
	return s
}

func (s *RunMetric) SetStep(v int64) *RunMetric {
	s.Step = &v
	return s
}

func (s *RunMetric) SetTimestamp(v int64) *RunMetric {
	s.Timestamp = &v
	return s
}

func (s *RunMetric) SetValue(v float32) *RunMetric {
	s.Value = &v
	return s
}

func (s *RunMetric) Validate() error {
	return dara.Validate(s)
}
