// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInstanceMetricsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDimensionsShrink(v string) *QueryInstanceMetricsShrinkRequest
	GetDimensionsShrink() *string
	SetEndTime(v string) *QueryInstanceMetricsShrinkRequest
	GetEndTime() *string
	SetMetricType(v string) *QueryInstanceMetricsShrinkRequest
	GetMetricType() *string
	SetStartTime(v string) *QueryInstanceMetricsShrinkRequest
	GetStartTime() *string
	SetTimeStep(v string) *QueryInstanceMetricsShrinkRequest
	GetTimeStep() *string
}

type QueryInstanceMetricsShrinkRequest struct {
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
	// StorageUsage
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

func (s QueryInstanceMetricsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryInstanceMetricsShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryInstanceMetricsShrinkRequest) GetDimensionsShrink() *string {
	return s.DimensionsShrink
}

func (s *QueryInstanceMetricsShrinkRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *QueryInstanceMetricsShrinkRequest) GetMetricType() *string {
	return s.MetricType
}

func (s *QueryInstanceMetricsShrinkRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *QueryInstanceMetricsShrinkRequest) GetTimeStep() *string {
	return s.TimeStep
}

func (s *QueryInstanceMetricsShrinkRequest) SetDimensionsShrink(v string) *QueryInstanceMetricsShrinkRequest {
	s.DimensionsShrink = &v
	return s
}

func (s *QueryInstanceMetricsShrinkRequest) SetEndTime(v string) *QueryInstanceMetricsShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *QueryInstanceMetricsShrinkRequest) SetMetricType(v string) *QueryInstanceMetricsShrinkRequest {
	s.MetricType = &v
	return s
}

func (s *QueryInstanceMetricsShrinkRequest) SetStartTime(v string) *QueryInstanceMetricsShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *QueryInstanceMetricsShrinkRequest) SetTimeStep(v string) *QueryInstanceMetricsShrinkRequest {
	s.TimeStep = &v
	return s
}

func (s *QueryInstanceMetricsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
