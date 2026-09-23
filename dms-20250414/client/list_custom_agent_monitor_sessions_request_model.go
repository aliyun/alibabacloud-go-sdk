// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomAgentMonitorSessionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomAgentId(v string) *ListCustomAgentMonitorSessionsRequest
	GetCustomAgentId() *string
	SetEndTime(v int64) *ListCustomAgentMonitorSessionsRequest
	GetEndTime() *int64
	SetPageNumber(v int32) *ListCustomAgentMonitorSessionsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCustomAgentMonitorSessionsRequest
	GetPageSize() *int32
	SetQueryType(v string) *ListCustomAgentMonitorSessionsRequest
	GetQueryType() *string
	SetStartTime(v int64) *ListCustomAgentMonitorSessionsRequest
	GetStartTime() *int64
	SetWorkspaceId(v string) *ListCustomAgentMonitorSessionsRequest
	GetWorkspaceId() *string
}

type ListCustomAgentMonitorSessionsRequest struct {
	// The custom agent ID.
	//
	// - Required only when QueryType is set to CustomAgent.
	//
	// example:
	//
	// ca-4y3ca4khkcu**********ysf
	CustomAgentId *string `json:"CustomAgentId,omitempty" xml:"CustomAgentId,omitempty"`
	// The end time for statistics (epoch millis).
	//
	// - Note: The maximum time range is 3 months.
	//
	// example:
	//
	// 1756742400000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The page number, starting from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 200. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The query scope. Default value: All. Valid values:
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
	// The start time for statistics (epoch millis).
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
	// 56kv1pvl9uvt9**********bb
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListCustomAgentMonitorSessionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCustomAgentMonitorSessionsRequest) GoString() string {
	return s.String()
}

func (s *ListCustomAgentMonitorSessionsRequest) GetCustomAgentId() *string {
	return s.CustomAgentId
}

func (s *ListCustomAgentMonitorSessionsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListCustomAgentMonitorSessionsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCustomAgentMonitorSessionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCustomAgentMonitorSessionsRequest) GetQueryType() *string {
	return s.QueryType
}

func (s *ListCustomAgentMonitorSessionsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListCustomAgentMonitorSessionsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListCustomAgentMonitorSessionsRequest) SetCustomAgentId(v string) *ListCustomAgentMonitorSessionsRequest {
	s.CustomAgentId = &v
	return s
}

func (s *ListCustomAgentMonitorSessionsRequest) SetEndTime(v int64) *ListCustomAgentMonitorSessionsRequest {
	s.EndTime = &v
	return s
}

func (s *ListCustomAgentMonitorSessionsRequest) SetPageNumber(v int32) *ListCustomAgentMonitorSessionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCustomAgentMonitorSessionsRequest) SetPageSize(v int32) *ListCustomAgentMonitorSessionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCustomAgentMonitorSessionsRequest) SetQueryType(v string) *ListCustomAgentMonitorSessionsRequest {
	s.QueryType = &v
	return s
}

func (s *ListCustomAgentMonitorSessionsRequest) SetStartTime(v int64) *ListCustomAgentMonitorSessionsRequest {
	s.StartTime = &v
	return s
}

func (s *ListCustomAgentMonitorSessionsRequest) SetWorkspaceId(v string) *ListCustomAgentMonitorSessionsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListCustomAgentMonitorSessionsRequest) Validate() error {
	return dara.Validate(s)
}
