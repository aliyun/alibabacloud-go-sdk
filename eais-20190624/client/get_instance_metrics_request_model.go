// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceMetricsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *GetInstanceMetricsRequest
	GetEndTime() *string
	SetInstanceId(v string) *GetInstanceMetricsRequest
	GetInstanceId() *string
	SetMetricType(v string) *GetInstanceMetricsRequest
	GetMetricType() *string
	SetRegionId(v string) *GetInstanceMetricsRequest
	GetRegionId() *string
	SetStartTime(v string) *GetInstanceMetricsRequest
	GetStartTime() *string
	SetTimeStep(v string) *GetInstanceMetricsRequest
	GetTimeStep() *string
}

type GetInstanceMetricsRequest struct {
	// example:
	//
	// 2022-11-22T16:30:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// eais-hznzre6ffmz9num4****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MemoryUsage
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 2022-11-22T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 5m
	TimeStep *string `json:"TimeStep,omitempty" xml:"TimeStep,omitempty"`
}

func (s GetInstanceMetricsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceMetricsRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceMetricsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetInstanceMetricsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceMetricsRequest) GetMetricType() *string {
	return s.MetricType
}

func (s *GetInstanceMetricsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetInstanceMetricsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetInstanceMetricsRequest) GetTimeStep() *string {
	return s.TimeStep
}

func (s *GetInstanceMetricsRequest) SetEndTime(v string) *GetInstanceMetricsRequest {
	s.EndTime = &v
	return s
}

func (s *GetInstanceMetricsRequest) SetInstanceId(v string) *GetInstanceMetricsRequest {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceMetricsRequest) SetMetricType(v string) *GetInstanceMetricsRequest {
	s.MetricType = &v
	return s
}

func (s *GetInstanceMetricsRequest) SetRegionId(v string) *GetInstanceMetricsRequest {
	s.RegionId = &v
	return s
}

func (s *GetInstanceMetricsRequest) SetStartTime(v string) *GetInstanceMetricsRequest {
	s.StartTime = &v
	return s
}

func (s *GetInstanceMetricsRequest) SetTimeStep(v string) *GetInstanceMetricsRequest {
	s.TimeStep = &v
	return s
}

func (s *GetInstanceMetricsRequest) Validate() error {
	return dara.Validate(s)
}
