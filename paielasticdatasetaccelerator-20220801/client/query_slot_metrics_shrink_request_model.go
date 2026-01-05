// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySlotMetricsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDimensionsShrink(v string) *QuerySlotMetricsShrinkRequest
	GetDimensionsShrink() *string
	SetEndTime(v string) *QuerySlotMetricsShrinkRequest
	GetEndTime() *string
	SetMetricType(v string) *QuerySlotMetricsShrinkRequest
	GetMetricType() *string
	SetStartTime(v string) *QuerySlotMetricsShrinkRequest
	GetStartTime() *string
	SetTimeStep(v string) *QuerySlotMetricsShrinkRequest
	GetTimeStep() *string
}

type QuerySlotMetricsShrinkRequest struct {
	// example:
	//
	// SlotIDs: xxx
	DimensionsShrink *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	// example:
	//
	// 2020-11-08T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// NetworkReceiveBytesPerSecond
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// example:
	//
	// 2020-11-08T15:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 5m
	TimeStep *string `json:"TimeStep,omitempty" xml:"TimeStep,omitempty"`
}

func (s QuerySlotMetricsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySlotMetricsShrinkRequest) GoString() string {
	return s.String()
}

func (s *QuerySlotMetricsShrinkRequest) GetDimensionsShrink() *string {
	return s.DimensionsShrink
}

func (s *QuerySlotMetricsShrinkRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *QuerySlotMetricsShrinkRequest) GetMetricType() *string {
	return s.MetricType
}

func (s *QuerySlotMetricsShrinkRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *QuerySlotMetricsShrinkRequest) GetTimeStep() *string {
	return s.TimeStep
}

func (s *QuerySlotMetricsShrinkRequest) SetDimensionsShrink(v string) *QuerySlotMetricsShrinkRequest {
	s.DimensionsShrink = &v
	return s
}

func (s *QuerySlotMetricsShrinkRequest) SetEndTime(v string) *QuerySlotMetricsShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *QuerySlotMetricsShrinkRequest) SetMetricType(v string) *QuerySlotMetricsShrinkRequest {
	s.MetricType = &v
	return s
}

func (s *QuerySlotMetricsShrinkRequest) SetStartTime(v string) *QuerySlotMetricsShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *QuerySlotMetricsShrinkRequest) SetTimeStep(v string) *QuerySlotMetricsShrinkRequest {
	s.TimeStep = &v
	return s
}

func (s *QuerySlotMetricsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
