// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySlotMetricsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDimensions(v map[string]interface{}) *QuerySlotMetricsRequest
	GetDimensions() map[string]interface{}
	SetEndTime(v string) *QuerySlotMetricsRequest
	GetEndTime() *string
	SetMetricType(v string) *QuerySlotMetricsRequest
	GetMetricType() *string
	SetStartTime(v string) *QuerySlotMetricsRequest
	GetStartTime() *string
	SetTimeStep(v string) *QuerySlotMetricsRequest
	GetTimeStep() *string
}

type QuerySlotMetricsRequest struct {
	// example:
	//
	// SlotIDs: xxx
	Dimensions map[string]interface{} `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
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

func (s QuerySlotMetricsRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySlotMetricsRequest) GoString() string {
	return s.String()
}

func (s *QuerySlotMetricsRequest) GetDimensions() map[string]interface{} {
	return s.Dimensions
}

func (s *QuerySlotMetricsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *QuerySlotMetricsRequest) GetMetricType() *string {
	return s.MetricType
}

func (s *QuerySlotMetricsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *QuerySlotMetricsRequest) GetTimeStep() *string {
	return s.TimeStep
}

func (s *QuerySlotMetricsRequest) SetDimensions(v map[string]interface{}) *QuerySlotMetricsRequest {
	s.Dimensions = v
	return s
}

func (s *QuerySlotMetricsRequest) SetEndTime(v string) *QuerySlotMetricsRequest {
	s.EndTime = &v
	return s
}

func (s *QuerySlotMetricsRequest) SetMetricType(v string) *QuerySlotMetricsRequest {
	s.MetricType = &v
	return s
}

func (s *QuerySlotMetricsRequest) SetStartTime(v string) *QuerySlotMetricsRequest {
	s.StartTime = &v
	return s
}

func (s *QuerySlotMetricsRequest) SetTimeStep(v string) *QuerySlotMetricsRequest {
	s.TimeStep = &v
	return s
}

func (s *QuerySlotMetricsRequest) Validate() error {
	return dara.Validate(s)
}
