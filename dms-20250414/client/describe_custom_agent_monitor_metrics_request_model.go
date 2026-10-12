// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomAgentMonitorMetricsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomAgentId(v string) *DescribeCustomAgentMonitorMetricsRequest
	GetCustomAgentId() *string
	SetEndTime(v int64) *DescribeCustomAgentMonitorMetricsRequest
	GetEndTime() *int64
	SetGranularity(v string) *DescribeCustomAgentMonitorMetricsRequest
	GetGranularity() *string
	SetQueryType(v string) *DescribeCustomAgentMonitorMetricsRequest
	GetQueryType() *string
	SetStartTime(v int64) *DescribeCustomAgentMonitorMetricsRequest
	GetStartTime() *int64
	SetWorkspaceId(v string) *DescribeCustomAgentMonitorMetricsRequest
	GetWorkspaceId() *string
}

type DescribeCustomAgentMonitorMetricsRequest struct {
	// The custom agent ID.
	//
	// - Required only when QueryType is set to CustomAgent.
	//
	// example:
	//
	// ca-a9fd******0lnq4g6c
	CustomAgentId *string `json:"CustomAgentId,omitempty" xml:"CustomAgentId,omitempty"`
	// The end time of the statistical period (epoch millis).
	//
	// - Note: The maximum time range is 3 months.
	//
	// example:
	//
	// 1756742400000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The aggregation granularity. Valid values:
	//
	// - DAY: daily. The maximum supported time range is 3 months.
	//
	// - HOUR: hourly. The maximum supported time range is 72 hours.
	//
	// example:
	//
	// DAY
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	// The statistical scope. Default value: All. Valid values:
	//
	// - Default: default DataAgent sessions.
	//
	// - CustomAgent: specified custom agent sessions.
	//
	// - All: all sessions in the workspace.
	//
	// example:
	//
	// All
	QueryType *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	// The start time of the statistical period (epoch millis).
	//
	// - Note: The maximum time range is 3 months.
	//
	// example:
	//
	// 1756656000000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 99fad******qg6c0l4nlacu
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DescribeCustomAgentMonitorMetricsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomAgentMonitorMetricsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomAgentMonitorMetricsRequest) GetCustomAgentId() *string {
	return s.CustomAgentId
}

func (s *DescribeCustomAgentMonitorMetricsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeCustomAgentMonitorMetricsRequest) GetGranularity() *string {
	return s.Granularity
}

func (s *DescribeCustomAgentMonitorMetricsRequest) GetQueryType() *string {
	return s.QueryType
}

func (s *DescribeCustomAgentMonitorMetricsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeCustomAgentMonitorMetricsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DescribeCustomAgentMonitorMetricsRequest) SetCustomAgentId(v string) *DescribeCustomAgentMonitorMetricsRequest {
	s.CustomAgentId = &v
	return s
}

func (s *DescribeCustomAgentMonitorMetricsRequest) SetEndTime(v int64) *DescribeCustomAgentMonitorMetricsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeCustomAgentMonitorMetricsRequest) SetGranularity(v string) *DescribeCustomAgentMonitorMetricsRequest {
	s.Granularity = &v
	return s
}

func (s *DescribeCustomAgentMonitorMetricsRequest) SetQueryType(v string) *DescribeCustomAgentMonitorMetricsRequest {
	s.QueryType = &v
	return s
}

func (s *DescribeCustomAgentMonitorMetricsRequest) SetStartTime(v int64) *DescribeCustomAgentMonitorMetricsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeCustomAgentMonitorMetricsRequest) SetWorkspaceId(v string) *DescribeCustomAgentMonitorMetricsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DescribeCustomAgentMonitorMetricsRequest) Validate() error {
	return dara.Validate(s)
}
