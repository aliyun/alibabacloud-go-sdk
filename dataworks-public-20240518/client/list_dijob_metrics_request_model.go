// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDIJobMetricsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDIJobId(v int64) *ListDIJobMetricsRequest
	GetDIJobId() *int64
	SetEndTime(v int64) *ListDIJobMetricsRequest
	GetEndTime() *int64
	SetMetricName(v []*string) *ListDIJobMetricsRequest
	GetMetricName() []*string
	SetStartTime(v int64) *ListDIJobMetricsRequest
	GetStartTime() *int64
}

type ListDIJobMetricsRequest struct {
	// The ID of the synchronization task.
	//
	// example:
	//
	// 11265
	DIJobId *int64 `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	// The end of the time range to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1712205941
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The metrics that you want to query.
	//
	// This parameter is required.
	MetricName []*string `json:"MetricName,omitempty" xml:"MetricName,omitempty" type:"Repeated"`
	// The beginning of the time range to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1586509407
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListDIJobMetricsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDIJobMetricsRequest) GoString() string {
	return s.String()
}

func (s *ListDIJobMetricsRequest) GetDIJobId() *int64 {
	return s.DIJobId
}

func (s *ListDIJobMetricsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListDIJobMetricsRequest) GetMetricName() []*string {
	return s.MetricName
}

func (s *ListDIJobMetricsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListDIJobMetricsRequest) SetDIJobId(v int64) *ListDIJobMetricsRequest {
	s.DIJobId = &v
	return s
}

func (s *ListDIJobMetricsRequest) SetEndTime(v int64) *ListDIJobMetricsRequest {
	s.EndTime = &v
	return s
}

func (s *ListDIJobMetricsRequest) SetMetricName(v []*string) *ListDIJobMetricsRequest {
	s.MetricName = v
	return s
}

func (s *ListDIJobMetricsRequest) SetStartTime(v int64) *ListDIJobMetricsRequest {
	s.StartTime = &v
	return s
}

func (s *ListDIJobMetricsRequest) Validate() error {
	return dara.Validate(s)
}
