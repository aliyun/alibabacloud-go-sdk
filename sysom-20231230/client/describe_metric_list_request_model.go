// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetricListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXDebugId(v string) *DescribeMetricListRequest
	GetXDebugId() *string
	SetEndTime(v int64) *DescribeMetricListRequest
	GetEndTime() *int64
	SetInstance(v string) *DescribeMetricListRequest
	GetInstance() *string
	SetMetricName(v string) *DescribeMetricListRequest
	GetMetricName() *string
	SetStartTime(v int64) *DescribeMetricListRequest
	GetStartTime() *int64
	SetXSysomInvokeSource(v string) *DescribeMetricListRequest
	GetXSysomInvokeSource() *string
}

type DescribeMetricListRequest struct {
	XDebugId *string `json:"X-Debug-Id,omitempty" xml:"X-Debug-Id,omitempty"`
	// The end time, in seconds (UNIX timestamp).
	//
	// example:
	//
	// 1683618245000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// i-wz9d00ut2ska3mlyhn6j
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// The metric name.
	//
	// example:
	//
	// sysom_cpu_graph
	MetricName *string `json:"metricName,omitempty" xml:"metricName,omitempty"`
	// The start time, in seconds (UNIX timestamp).
	//
	// example:
	//
	// 1709740800000
	StartTime          *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	XSysomInvokeSource *string `json:"x-sysom-invoke-source,omitempty" xml:"x-sysom-invoke-source,omitempty"`
}

func (s DescribeMetricListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricListRequest) GoString() string {
	return s.String()
}

func (s *DescribeMetricListRequest) GetXDebugId() *string {
	return s.XDebugId
}

func (s *DescribeMetricListRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeMetricListRequest) GetInstance() *string {
	return s.Instance
}

func (s *DescribeMetricListRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeMetricListRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeMetricListRequest) GetXSysomInvokeSource() *string {
	return s.XSysomInvokeSource
}

func (s *DescribeMetricListRequest) SetXDebugId(v string) *DescribeMetricListRequest {
	s.XDebugId = &v
	return s
}

func (s *DescribeMetricListRequest) SetEndTime(v int64) *DescribeMetricListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeMetricListRequest) SetInstance(v string) *DescribeMetricListRequest {
	s.Instance = &v
	return s
}

func (s *DescribeMetricListRequest) SetMetricName(v string) *DescribeMetricListRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeMetricListRequest) SetStartTime(v int64) *DescribeMetricListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeMetricListRequest) SetXSysomInvokeSource(v string) *DescribeMetricListRequest {
	s.XSysomInvokeSource = &v
	return s
}

func (s *DescribeMetricListRequest) Validate() error {
	return dara.Validate(s)
}
