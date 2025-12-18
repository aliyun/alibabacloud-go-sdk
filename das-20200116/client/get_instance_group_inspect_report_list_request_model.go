// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceGroupInspectReportListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *GetInstanceGroupInspectReportListRequest
	GetAgentId() *string
	SetEndTime(v int64) *GetInstanceGroupInspectReportListRequest
	GetEndTime() *int64
	SetGroupId(v string) *GetInstanceGroupInspectReportListRequest
	GetGroupId() *string
	SetStartTime(v int64) *GetInstanceGroupInspectReportListRequest
	GetStartTime() *int64
}

type GetInstanceGroupInspectReportListRequest struct {
	// example:
	//
	// ag-472T0DxtmjIxxxxx
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1655427625000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// null
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1596177993000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetInstanceGroupInspectReportListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceGroupInspectReportListRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceGroupInspectReportListRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *GetInstanceGroupInspectReportListRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetInstanceGroupInspectReportListRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *GetInstanceGroupInspectReportListRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetInstanceGroupInspectReportListRequest) SetAgentId(v string) *GetInstanceGroupInspectReportListRequest {
	s.AgentId = &v
	return s
}

func (s *GetInstanceGroupInspectReportListRequest) SetEndTime(v int64) *GetInstanceGroupInspectReportListRequest {
	s.EndTime = &v
	return s
}

func (s *GetInstanceGroupInspectReportListRequest) SetGroupId(v string) *GetInstanceGroupInspectReportListRequest {
	s.GroupId = &v
	return s
}

func (s *GetInstanceGroupInspectReportListRequest) SetStartTime(v int64) *GetInstanceGroupInspectReportListRequest {
	s.StartTime = &v
	return s
}

func (s *GetInstanceGroupInspectReportListRequest) Validate() error {
	return dara.Validate(s)
}
