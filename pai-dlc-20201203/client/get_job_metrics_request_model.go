// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobMetricsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *GetJobMetricsRequest
	GetEndTime() *string
	SetMetricType(v string) *GetJobMetricsRequest
	GetMetricType() *string
	SetStartTime(v string) *GetJobMetricsRequest
	GetStartTime() *string
	SetTimeStep(v string) *GetJobMetricsRequest
	GetTimeStep() *string
	SetToken(v string) *GetJobMetricsRequest
	GetToken() *string
}

type GetJobMetricsRequest struct {
	// The end time of the time range for querying monitoring data, in UTC. Default value: the current time.
	//
	// example:
	//
	// 2020-11-09T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The metric type of the monitoring data to query. Valid values:
	//
	// - GpuCoreUsage: the GPU usage
	//
	// - GpuMemoryUsage: the GPU memory usage
	//
	// - CpuCoreUsage: the CPU usage
	//
	// - MemoryUsage: the memory usage
	//
	// - NetworkInputRate: the network input rate
	//
	// - NetworkOutputRate: the network output rate
	//
	// - DiskReadRate: the disk read rate
	//
	// - DiskWriteRate: the disk write rate
	//
	// This parameter is required.
	//
	// example:
	//
	// GpuMemoryUsage
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// The start time of the time range for querying monitoring data, in UTC. Default value: 1 hour ago.
	//
	// example:
	//
	// 2020-11-08T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The time interval at which monitoring data is returned. Default value: 5 minutes.
	//
	// example:
	//
	// 5m
	TimeStep *string `json:"TimeStep,omitempty" xml:"TimeStep,omitempty"`
	// The temporary token used for authentication.
	//
	// example:
	//
	// eyXXXX-XXXX.XXXXX
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetJobMetricsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetJobMetricsRequest) GoString() string {
	return s.String()
}

func (s *GetJobMetricsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetJobMetricsRequest) GetMetricType() *string {
	return s.MetricType
}

func (s *GetJobMetricsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetJobMetricsRequest) GetTimeStep() *string {
	return s.TimeStep
}

func (s *GetJobMetricsRequest) GetToken() *string {
	return s.Token
}

func (s *GetJobMetricsRequest) SetEndTime(v string) *GetJobMetricsRequest {
	s.EndTime = &v
	return s
}

func (s *GetJobMetricsRequest) SetMetricType(v string) *GetJobMetricsRequest {
	s.MetricType = &v
	return s
}

func (s *GetJobMetricsRequest) SetStartTime(v string) *GetJobMetricsRequest {
	s.StartTime = &v
	return s
}

func (s *GetJobMetricsRequest) SetTimeStep(v string) *GetJobMetricsRequest {
	s.TimeStep = &v
	return s
}

func (s *GetJobMetricsRequest) SetToken(v string) *GetJobMetricsRequest {
	s.Token = &v
	return s
}

func (s *GetJobMetricsRequest) Validate() error {
	return dara.Validate(s)
}
