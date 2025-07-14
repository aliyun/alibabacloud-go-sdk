// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCubePerformanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCostTimeAvgMin(v int32) *QueryCubePerformanceRequest
	GetCostTimeAvgMin() *int32
	SetCubeId(v string) *QueryCubePerformanceRequest
	GetCubeId() *string
	SetPageNum(v int32) *QueryCubePerformanceRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QueryCubePerformanceRequest
	GetPageSize() *int32
	SetQueryType(v string) *QueryCubePerformanceRequest
	GetQueryType() *string
	SetWorkspaceId(v string) *QueryCubePerformanceRequest
	GetWorkspaceId() *string
}

type QueryCubePerformanceRequest struct {
	// The average duration (minutes).
	//
	// example:
	//
	// 1
	CostTimeAvgMin *int32 `json:"CostTimeAvgMin,omitempty" xml:"CostTimeAvgMin,omitempty"`
	// The dataset ID.
	//
	// example:
	//
	// 7c7223ae-****-3c744528014b
	CubeId *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
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
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 95296e95-ca89-4c7d-8af9-dedf0ad0****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryCubePerformanceRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCubePerformanceRequest) GoString() string {
	return s.String()
}

func (s *QueryCubePerformanceRequest) GetCostTimeAvgMin() *int32 {
	return s.CostTimeAvgMin
}

func (s *QueryCubePerformanceRequest) GetCubeId() *string {
	return s.CubeId
}

func (s *QueryCubePerformanceRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryCubePerformanceRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryCubePerformanceRequest) GetQueryType() *string {
	return s.QueryType
}

func (s *QueryCubePerformanceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryCubePerformanceRequest) SetCostTimeAvgMin(v int32) *QueryCubePerformanceRequest {
	s.CostTimeAvgMin = &v
	return s
}

func (s *QueryCubePerformanceRequest) SetCubeId(v string) *QueryCubePerformanceRequest {
	s.CubeId = &v
	return s
}

func (s *QueryCubePerformanceRequest) SetPageNum(v int32) *QueryCubePerformanceRequest {
	s.PageNum = &v
	return s
}

func (s *QueryCubePerformanceRequest) SetPageSize(v int32) *QueryCubePerformanceRequest {
	s.PageSize = &v
	return s
}

func (s *QueryCubePerformanceRequest) SetQueryType(v string) *QueryCubePerformanceRequest {
	s.QueryType = &v
	return s
}

func (s *QueryCubePerformanceRequest) SetWorkspaceId(v string) *QueryCubePerformanceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *QueryCubePerformanceRequest) Validate() error {
	return dara.Validate(s)
}
