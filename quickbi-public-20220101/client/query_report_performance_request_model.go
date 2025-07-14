// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryReportPerformanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCostTimeAvgMin(v int32) *QueryReportPerformanceRequest
	GetCostTimeAvgMin() *int32
	SetPageNum(v int32) *QueryReportPerformanceRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QueryReportPerformanceRequest
	GetPageSize() *int32
	SetQueryType(v string) *QueryReportPerformanceRequest
	GetQueryType() *string
	SetReportId(v string) *QueryReportPerformanceRequest
	GetReportId() *string
	SetResourceType(v string) *QueryReportPerformanceRequest
	GetResourceType() *string
	SetWorkspaceId(v string) *QueryReportPerformanceRequest
	GetWorkspaceId() *string
}

type QueryReportPerformanceRequest struct {
	// The average duration (minutes).
	//
	// example:
	//
	// 1
	CostTimeAvgMin *int32 `json:"CostTimeAvgMin,omitempty" xml:"CostTimeAvgMin,omitempty"`
	// Current page number for organization member list:
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
	// 10
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
	// The ID of the security report.
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
	// 95296e95-ca89-4c7d-8af9-dedf0ad0****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryReportPerformanceRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryReportPerformanceRequest) GoString() string {
	return s.String()
}

func (s *QueryReportPerformanceRequest) GetCostTimeAvgMin() *int32 {
	return s.CostTimeAvgMin
}

func (s *QueryReportPerformanceRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryReportPerformanceRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryReportPerformanceRequest) GetQueryType() *string {
	return s.QueryType
}

func (s *QueryReportPerformanceRequest) GetReportId() *string {
	return s.ReportId
}

func (s *QueryReportPerformanceRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *QueryReportPerformanceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryReportPerformanceRequest) SetCostTimeAvgMin(v int32) *QueryReportPerformanceRequest {
	s.CostTimeAvgMin = &v
	return s
}

func (s *QueryReportPerformanceRequest) SetPageNum(v int32) *QueryReportPerformanceRequest {
	s.PageNum = &v
	return s
}

func (s *QueryReportPerformanceRequest) SetPageSize(v int32) *QueryReportPerformanceRequest {
	s.PageSize = &v
	return s
}

func (s *QueryReportPerformanceRequest) SetQueryType(v string) *QueryReportPerformanceRequest {
	s.QueryType = &v
	return s
}

func (s *QueryReportPerformanceRequest) SetReportId(v string) *QueryReportPerformanceRequest {
	s.ReportId = &v
	return s
}

func (s *QueryReportPerformanceRequest) SetResourceType(v string) *QueryReportPerformanceRequest {
	s.ResourceType = &v
	return s
}

func (s *QueryReportPerformanceRequest) SetWorkspaceId(v string) *QueryReportPerformanceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *QueryReportPerformanceRequest) Validate() error {
	return dara.Validate(s)
}
