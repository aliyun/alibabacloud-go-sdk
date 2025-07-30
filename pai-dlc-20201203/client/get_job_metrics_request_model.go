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
	// The end time of the time range to query monitoring data. The time is displayed in UTC. The default value is the current time.
	//
	// example:
	//
	// 2020-11-09T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The type of the monitoring metrics. Valid values:
	//
	// 	- GpuCoreUsage: GPU utilization
	//
	// 	- GpuMemoryUsage: GPU memory utilization
	//
	// 	- CpuCoreUsage: CPU utilization
	//
	// 	- MemoryUsage: memory utilization
	//
	// 	- NetworkInputRate: the network write in rate.
	//
	// 	- NetworkOutputRate: the network write out rate
	//
	// 	- DiskReadRate: the disk read rate
	//
	// 	- DiskWriteRate: the disk write rate
	//
	// This parameter is required.
	//
	// example:
	//
	// GpuMemoryUsage
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// The beginning of the time range to query monitoring data. The time is displayed in UTC. The default value is the time 1 hour before the current time.
	//
	// example:
	//
	// 2020-11-08T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The interval at which monitoring data is returned. Default value: 5. Unit: minutes.
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
