// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iThresholdList interface {
	dara.Model
	String() string
	GoString() string
	SetMax(v float64) *ThresholdList
	GetMax() *float64
	SetMin(v float64) *ThresholdList
	GetMin() *float64
	SetSeverity(v string) *ThresholdList
	GetSeverity() *string
	SetThreshold(v float32) *ThresholdList
	GetThreshold() *float32
}

type ThresholdList struct {
	Max *float64 `json:"max,omitempty" xml:"max,omitempty"`
	Min *float64 `json:"min,omitempty" xml:"min,omitempty"`
	// This parameter is required.
	Severity  *string  `json:"severity,omitempty" xml:"severity,omitempty"`
	Threshold *float32 `json:"threshold,omitempty" xml:"threshold,omitempty"`
}

func (s ThresholdList) String() string {
	return dara.Prettify(s)
}

func (s ThresholdList) GoString() string {
	return s.String()
}

func (s *ThresholdList) GetMax() *float64 {
	return s.Max
}

func (s *ThresholdList) GetMin() *float64 {
	return s.Min
}

func (s *ThresholdList) GetSeverity() *string {
	return s.Severity
}

func (s *ThresholdList) GetThreshold() *float32 {
	return s.Threshold
}

func (s *ThresholdList) SetMax(v float64) *ThresholdList {
	s.Max = &v
	return s
}

func (s *ThresholdList) SetMin(v float64) *ThresholdList {
	s.Min = &v
	return s
}

func (s *ThresholdList) SetSeverity(v string) *ThresholdList {
	s.Severity = &v
	return s
}

func (s *ThresholdList) SetThreshold(v float32) *ThresholdList {
	s.Threshold = &v
	return s
}

func (s *ThresholdList) Validate() error {
	return dara.Validate(s)
}
