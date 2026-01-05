// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInstanceMetricsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDimensions(v map[string]interface{}) *QueryInstanceMetricsRequest
	GetDimensions() map[string]interface{}
	SetEndTime(v string) *QueryInstanceMetricsRequest
	GetEndTime() *string
	SetMetricType(v string) *QueryInstanceMetricsRequest
	GetMetricType() *string
	SetStartTime(v string) *QueryInstanceMetricsRequest
	GetStartTime() *string
	SetTimeStep(v string) *QueryInstanceMetricsRequest
	GetTimeStep() *string
}

type QueryInstanceMetricsRequest struct {
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

func (s QueryInstanceMetricsRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryInstanceMetricsRequest) GoString() string {
	return s.String()
}

func (s *QueryInstanceMetricsRequest) GetDimensions() map[string]interface{} {
	return s.Dimensions
}

func (s *QueryInstanceMetricsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *QueryInstanceMetricsRequest) GetMetricType() *string {
	return s.MetricType
}

func (s *QueryInstanceMetricsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *QueryInstanceMetricsRequest) GetTimeStep() *string {
	return s.TimeStep
}

func (s *QueryInstanceMetricsRequest) SetDimensions(v map[string]interface{}) *QueryInstanceMetricsRequest {
	s.Dimensions = v
	return s
}

func (s *QueryInstanceMetricsRequest) SetEndTime(v string) *QueryInstanceMetricsRequest {
	s.EndTime = &v
	return s
}

func (s *QueryInstanceMetricsRequest) SetMetricType(v string) *QueryInstanceMetricsRequest {
	s.MetricType = &v
	return s
}

func (s *QueryInstanceMetricsRequest) SetStartTime(v string) *QueryInstanceMetricsRequest {
	s.StartTime = &v
	return s
}

func (s *QueryInstanceMetricsRequest) SetTimeStep(v string) *QueryInstanceMetricsRequest {
	s.TimeStep = &v
	return s
}

func (s *QueryInstanceMetricsRequest) Validate() error {
	return dara.Validate(s)
}
