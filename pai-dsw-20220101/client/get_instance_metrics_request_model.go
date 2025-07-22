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
	SetMetricType(v string) *GetInstanceMetricsRequest
	GetMetricType() *string
	SetStartTime(v string) *GetInstanceMetricsRequest
	GetStartTime() *string
	SetTimeStep(v string) *GetInstanceMetricsRequest
	GetTimeStep() *string
}

type GetInstanceMetricsRequest struct {
	// The end of the time range to query.
	//
	// example:
	//
	// 2020-11-08T15:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The metric type. Valid values:
	//
	// 	- GpuCoreUsage: the GPU utilization.
	//
	// 	- GpuMemoryUsage: the GPU memory utilization.
	//
	// 	- CpuCoreUsage: the CPU utilization.
	//
	// 	- MemoryUsage: the memory utilization.
	//
	// 	- NetworkInputRate: the network ingress rate.
	//
	// 	- NetworkOutputRate: the network egress rate.
	//
	// 	- DiskReadRate: the disk read rate.
	//
	// 	- DiskWriteRate: the disk write rate.
	//
	// This parameter is required.
	//
	// example:
	//
	// GpuCoreUsage
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// The beginning of the time range to query.
	//
	// example:
	//
	// 2020-11-08T15:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The interval at which metrics are returned. Unit: minutes.
	//
	// example:
	//
	// 15m
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

func (s *GetInstanceMetricsRequest) GetMetricType() *string {
	return s.MetricType
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

func (s *GetInstanceMetricsRequest) SetMetricType(v string) *GetInstanceMetricsRequest {
	s.MetricType = &v
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
