// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTrainingJobInstanceMetricsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListTrainingJobInstanceMetricsRequest
	GetEndTime() *string
	SetInstanceId(v string) *ListTrainingJobInstanceMetricsRequest
	GetInstanceId() *string
	SetMetricType(v string) *ListTrainingJobInstanceMetricsRequest
	GetMetricType() *string
	SetStartTime(v string) *ListTrainingJobInstanceMetricsRequest
	GetStartTime() *string
	SetTimeStep(v string) *ListTrainingJobInstanceMetricsRequest
	GetTimeStep() *string
}

type ListTrainingJobInstanceMetricsRequest struct {
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2020-11-08T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// trains930928remn-master-0
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// GpuCoreUsage
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2020-11-08T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 10s
	TimeStep *string `json:"TimeStep,omitempty" xml:"TimeStep,omitempty"`
}

func (s ListTrainingJobInstanceMetricsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTrainingJobInstanceMetricsRequest) GoString() string {
	return s.String()
}

func (s *ListTrainingJobInstanceMetricsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListTrainingJobInstanceMetricsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListTrainingJobInstanceMetricsRequest) GetMetricType() *string {
	return s.MetricType
}

func (s *ListTrainingJobInstanceMetricsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListTrainingJobInstanceMetricsRequest) GetTimeStep() *string {
	return s.TimeStep
}

func (s *ListTrainingJobInstanceMetricsRequest) SetEndTime(v string) *ListTrainingJobInstanceMetricsRequest {
	s.EndTime = &v
	return s
}

func (s *ListTrainingJobInstanceMetricsRequest) SetInstanceId(v string) *ListTrainingJobInstanceMetricsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListTrainingJobInstanceMetricsRequest) SetMetricType(v string) *ListTrainingJobInstanceMetricsRequest {
	s.MetricType = &v
	return s
}

func (s *ListTrainingJobInstanceMetricsRequest) SetStartTime(v string) *ListTrainingJobInstanceMetricsRequest {
	s.StartTime = &v
	return s
}

func (s *ListTrainingJobInstanceMetricsRequest) SetTimeStep(v string) *ListTrainingJobInstanceMetricsRequest {
	s.TimeStep = &v
	return s
}

func (s *ListTrainingJobInstanceMetricsRequest) Validate() error {
	return dara.Validate(s)
}
