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
	// The data arrival delay in minutes. After a window ends, the system waits for this duration before creating a run to allow data to arrive completely. Default value: 0.
	//
	// example:
	//
	// 5
	DataDelayMinutes *int32 `json:"dataDelayMinutes,omitempty" xml:"dataDelayMinutes,omitempty"`
	// Specifies whether to enable continuous evaluation. If this parameter is not specified or is set to true, continuous evaluation is enabled. If this parameter is set to false, continuous evaluation is disabled but the configuration is retained.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The unit of the continuous evaluation window interval. This field is required for the current polling implementation.
	//
	// example:
	//
	// HOUR
	IntervalUnit *string `json:"intervalUnit,omitempty" xml:"intervalUnit,omitempty"`
	// The size of the continuous evaluation window interval. This parameter is used together with intervalUnit. The value must be greater than 0.
	//
	// example:
	//
	// 1
	IntervalValue *int32 `json:"intervalValue,omitempty" xml:"intervalValue,omitempty"`
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
