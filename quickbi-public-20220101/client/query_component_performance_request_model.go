// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryComponentPerformanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCostTimeAvgMin(v int32) *QueryComponentPerformanceRequest
	GetCostTimeAvgMin() *int32
	SetPageNum(v int32) *QueryComponentPerformanceRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QueryComponentPerformanceRequest
	GetPageSize() *int32
	SetQueryType(v string) *QueryComponentPerformanceRequest
	GetQueryType() *string
	SetReportId(v string) *QueryComponentPerformanceRequest
	GetReportId() *string
	SetResourceType(v string) *QueryComponentPerformanceRequest
	GetResourceType() *string
	SetWorkspaceId(v string) *QueryComponentPerformanceRequest
	GetWorkspaceId() *string
}

type QueryComponentPerformanceRequest struct {
	// The average duration (minutes).
	//
	// example:
	//
	// 1
	CostTimeAvgMin *int32 `json:"CostTimeAvgMin,omitempty" xml:"CostTimeAvgMin,omitempty"`
	// The current page number of the workspace member list:
	//
	// 	- Pages start from page 1.
	//
	// 	- Default value: 1.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of rows per page in a paged query.
	//
	// 	- Default value: 10.
	//
	// 	- Maximum value: 1,000.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The query type. Valid values:
	//
	// 	- **lastDay**: Yesterday
	//
	// 	- **sevenDays**: Within seven days
	//
	// 	- **thirtyDays**: Within 30 days
	//
	// This parameter is required.
	//
	// example:
	//
	// sevenDays
	QueryType *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	// The ID of the work. The works here include BI portal, dashboards, spreadsheets, and self-service access.
	//
	// example:
	//
	// 6b407e50-e774-406b-9956-da2425c2****
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The resource types.
	//
	// example:
	//
	// report
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 89713491-cb4f-4579-b889-e82c35f1****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryComponentPerformanceRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryComponentPerformanceRequest) GoString() string {
	return s.String()
}

func (s *QueryComponentPerformanceRequest) GetCostTimeAvgMin() *int32 {
	return s.CostTimeAvgMin
}

func (s *QueryComponentPerformanceRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryComponentPerformanceRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryComponentPerformanceRequest) GetQueryType() *string {
	return s.QueryType
}

func (s *QueryComponentPerformanceRequest) GetReportId() *string {
	return s.ReportId
}

func (s *QueryComponentPerformanceRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *QueryComponentPerformanceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryComponentPerformanceRequest) SetCostTimeAvgMin(v int32) *QueryComponentPerformanceRequest {
	s.CostTimeAvgMin = &v
	return s
}

func (s *QueryComponentPerformanceRequest) SetPageNum(v int32) *QueryComponentPerformanceRequest {
	s.PageNum = &v
	return s
}

func (s *QueryComponentPerformanceRequest) SetPageSize(v int32) *QueryComponentPerformanceRequest {
	s.PageSize = &v
	return s
}

func (s *QueryComponentPerformanceRequest) SetQueryType(v string) *QueryComponentPerformanceRequest {
	s.QueryType = &v
	return s
}

func (s *QueryComponentPerformanceRequest) SetReportId(v string) *QueryComponentPerformanceRequest {
	s.ReportId = &v
	return s
}

func (s *QueryComponentPerformanceRequest) SetResourceType(v string) *QueryComponentPerformanceRequest {
	s.ResourceType = &v
	return s
}

func (s *QueryComponentPerformanceRequest) SetWorkspaceId(v string) *QueryComponentPerformanceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *QueryComponentPerformanceRequest) Validate() error {
	return dara.Validate(s)
}
