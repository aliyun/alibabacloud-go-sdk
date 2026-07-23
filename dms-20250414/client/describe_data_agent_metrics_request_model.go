// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataAgentMetricsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DescribeDataAgentMetricsRequest
	GetEndTime() *int64
	SetMetricNames(v string) *DescribeDataAgentMetricsRequest
	GetMetricNames() *string
	SetMetricType(v string) *DescribeDataAgentMetricsRequest
	GetMetricType() *string
	SetStartTime(v int64) *DescribeDataAgentMetricsRequest
	GetStartTime() *int64
}

type DescribeDataAgentMetricsRequest struct {
	// The end time of the query range.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1782836200000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The metric names. Separate multiple names with commas (,).
	//
	// example:
	//
	// data_agent_session_per_user
	MetricNames *string `json:"MetricNames,omitempty" xml:"MetricNames,omitempty"`
	// The metric type. Valid values:
	//
	// - **basic**: basic metrics.
	//
	// - **high_level**: advanced metrics.
	//
	// This parameter is required.
	//
	// example:
	//
	// basic
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// The start time of the query range.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1782835200000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDataAgentMetricsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataAgentMetricsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataAgentMetricsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDataAgentMetricsRequest) GetMetricNames() *string {
	return s.MetricNames
}

func (s *DescribeDataAgentMetricsRequest) GetMetricType() *string {
	return s.MetricType
}

func (s *DescribeDataAgentMetricsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDataAgentMetricsRequest) SetEndTime(v int64) *DescribeDataAgentMetricsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDataAgentMetricsRequest) SetMetricNames(v string) *DescribeDataAgentMetricsRequest {
	s.MetricNames = &v
	return s
}

func (s *DescribeDataAgentMetricsRequest) SetMetricType(v string) *DescribeDataAgentMetricsRequest {
	s.MetricType = &v
	return s
}

func (s *DescribeDataAgentMetricsRequest) SetStartTime(v int64) *DescribeDataAgentMetricsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDataAgentMetricsRequest) Validate() error {
	return dara.Validate(s)
}
