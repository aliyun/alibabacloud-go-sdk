// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContinuousStrategy interface {
	dara.Model
	String() string
	GoString() string
	SetDataDelayMinutes(v int32) *ContinuousStrategy
	GetDataDelayMinutes() *int32
	SetEnabled(v bool) *ContinuousStrategy
	GetEnabled() *bool
	SetIntervalUnit(v string) *ContinuousStrategy
	GetIntervalUnit() *string
	SetIntervalValue(v int32) *ContinuousStrategy
	GetIntervalValue() *int32
}

type ContinuousStrategy struct {
	DataDelayMinutes *int32  `json:"dataDelayMinutes,omitempty" xml:"dataDelayMinutes,omitempty"`
	Enabled          *bool   `json:"enabled,omitempty" xml:"enabled,omitempty"`
	IntervalUnit     *string `json:"intervalUnit,omitempty" xml:"intervalUnit,omitempty"`
	IntervalValue    *int32  `json:"intervalValue,omitempty" xml:"intervalValue,omitempty"`
}

func (s ContinuousStrategy) String() string {
	return dara.Prettify(s)
}

func (s ContinuousStrategy) GoString() string {
	return s.String()
}

func (s *ContinuousStrategy) GetDataDelayMinutes() *int32 {
	return s.DataDelayMinutes
}

func (s *ContinuousStrategy) GetEnabled() *bool {
	return s.Enabled
}

func (s *ContinuousStrategy) GetIntervalUnit() *string {
	return s.IntervalUnit
}

func (s *ContinuousStrategy) GetIntervalValue() *int32 {
	return s.IntervalValue
}

func (s *ContinuousStrategy) SetDataDelayMinutes(v int32) *ContinuousStrategy {
	s.DataDelayMinutes = &v
	return s
}

func (s *ContinuousStrategy) SetEnabled(v bool) *ContinuousStrategy {
	s.Enabled = &v
	return s
}

func (s *ContinuousStrategy) SetIntervalUnit(v string) *ContinuousStrategy {
	s.IntervalUnit = &v
	return s
}

func (s *ContinuousStrategy) SetIntervalValue(v int32) *ContinuousStrategy {
	s.IntervalValue = &v
	return s
}

func (s *ContinuousStrategy) Validate() error {
	return dara.Validate(s)
}
