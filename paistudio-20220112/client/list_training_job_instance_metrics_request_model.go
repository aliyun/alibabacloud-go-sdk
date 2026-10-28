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
	// The end UTC time in ISO 8601 format. If left empty, the current time is used.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2020-11-08T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// trains930928remn-master-0
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The metric type. Valid values:
	//
	// - GpuCoreUsage: POD GPU utilization.
	//
	// - GpuMemoryUsage: POD GPU memory utilization.
	//
	// - CpuCoreUsage: POD CPU utilization.
	//
	// - MemoryUsage: POD memory utilization.
	//
	// - NetworkInputRate: POD network write rate (TCP/IP) (MB/s).
	//
	// - NetworkOutputRate: POD network read rate (TCP/IP) (MB/s).
	//
	// - DiskReadRate: POD disk read rate (MB/s).
	//
	// - DiskWriteRate: POD disk write rate (MB/s).
	//
	// This parameter is required.
	//
	// example:
	//
	// GpuCoreUsage
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// The start UTC time in ISO 8601 format. If left empty, the job start time is used.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2020-11-08T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The metric collection interval. Valid values: 1h, 30m, 5m, and 10s. h indicates hours, m indicates minutes, and s indicates seconds.
	//
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
