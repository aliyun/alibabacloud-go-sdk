// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataServiceApiCallStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQuery(v *ListDataServiceApiCallStatisticsRequestListQuery) *ListDataServiceApiCallStatisticsRequest
	GetListQuery() *ListDataServiceApiCallStatisticsRequestListQuery
	SetOpTenantId(v int64) *ListDataServiceApiCallStatisticsRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ListDataServiceApiCallStatisticsRequest
	GetOpUserId() *string
	SetProjectId(v int32) *ListDataServiceApiCallStatisticsRequest
	GetProjectId() *int32
}

type ListDataServiceApiCallStatisticsRequest struct {
	// The search conditions.
	//
	// This parameter is required.
	ListQuery *ListDataServiceApiCallStatisticsRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// The data service project ID.
	//
	// > If this parameter is not specified:
	//
	// > - Super administrators can query data of all service projects.
	//
	// > - Users other than super administrators can query data of service projects where they have the service project administrator or O&M user permissions.
	//
	// example:
	//
	// 102102
	ProjectId *int32 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListDataServiceApiCallStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceApiCallStatisticsRequest) GoString() string {
	return s.String()
}

func (s *ListDataServiceApiCallStatisticsRequest) GetListQuery() *ListDataServiceApiCallStatisticsRequestListQuery {
	return s.ListQuery
}

func (s *ListDataServiceApiCallStatisticsRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListDataServiceApiCallStatisticsRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListDataServiceApiCallStatisticsRequest) GetProjectId() *int32 {
	return s.ProjectId
}

func (s *ListDataServiceApiCallStatisticsRequest) SetListQuery(v *ListDataServiceApiCallStatisticsRequestListQuery) *ListDataServiceApiCallStatisticsRequest {
	s.ListQuery = v
	return s
}

func (s *ListDataServiceApiCallStatisticsRequest) SetOpTenantId(v int64) *ListDataServiceApiCallStatisticsRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListDataServiceApiCallStatisticsRequest) SetOpUserId(v string) *ListDataServiceApiCallStatisticsRequest {
	s.OpUserId = &v
	return s
}

func (s *ListDataServiceApiCallStatisticsRequest) SetProjectId(v int32) *ListDataServiceApiCallStatisticsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDataServiceApiCallStatisticsRequest) Validate() error {
	if s.ListQuery != nil {
		if err := s.ListQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDataServiceApiCallStatisticsRequestListQuery struct {
	// The end time. Format: yyyy-MM-dd HH:mm:ss.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2025-06-30 20:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The search keyword.
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The column by which to sort the results. Valid values:
	//
	// - CALL_COUNT: the number of calls.
	//
	// - ERROR_COUNT: the number of exceptions.
	//
	// - ERROR_RATE: the error rate.
	//
	// - AVG_RESPONSE_TIME: the average response time.
	//
	// - OFFLINE_RATE: the offline percentage.
	//
	// Default value: CALL_COUNT.
	//
	// example:
	//
	// CALL_COUNT
	OrderColumn *string `json:"OrderColumn,omitempty" xml:"OrderColumn,omitempty"`
	// The sort order. Valid values:
	//
	// - 1: ascending order.
	//
	// - 2: descending order.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	OrderType *int32 `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The start time. Format: yyyy-MM-dd HH:mm:ss.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2025-06-30 08:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListDataServiceApiCallStatisticsRequestListQuery) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceApiCallStatisticsRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListDataServiceApiCallStatisticsRequestListQuery) GetEndTime() *string {
	return s.EndTime
}

func (s *ListDataServiceApiCallStatisticsRequestListQuery) GetKeyword() *string {
	return s.Keyword
}

func (s *ListDataServiceApiCallStatisticsRequestListQuery) GetOrderColumn() *string {
	return s.OrderColumn
}

func (s *ListDataServiceApiCallStatisticsRequestListQuery) GetOrderType() *int32 {
	return s.OrderType
}

func (s *ListDataServiceApiCallStatisticsRequestListQuery) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListDataServiceApiCallStatisticsRequestListQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataServiceApiCallStatisticsRequestListQuery) GetStartTime() *string {
	return s.StartTime
}

func (s *ListDataServiceApiCallStatisticsRequestListQuery) SetEndTime(v string) *ListDataServiceApiCallStatisticsRequestListQuery {
	s.EndTime = &v
	return s
}

func (s *ListDataServiceApiCallStatisticsRequestListQuery) SetKeyword(v string) *ListDataServiceApiCallStatisticsRequestListQuery {
	s.Keyword = &v
	return s
}

func (s *ListDataServiceApiCallStatisticsRequestListQuery) SetOrderColumn(v string) *ListDataServiceApiCallStatisticsRequestListQuery {
	s.OrderColumn = &v
	return s
}

func (s *ListDataServiceApiCallStatisticsRequestListQuery) SetOrderType(v int32) *ListDataServiceApiCallStatisticsRequestListQuery {
	s.OrderType = &v
	return s
}

func (s *ListDataServiceApiCallStatisticsRequestListQuery) SetPageNo(v int32) *ListDataServiceApiCallStatisticsRequestListQuery {
	s.PageNo = &v
	return s
}

func (s *ListDataServiceApiCallStatisticsRequestListQuery) SetPageSize(v int32) *ListDataServiceApiCallStatisticsRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListDataServiceApiCallStatisticsRequestListQuery) SetStartTime(v string) *ListDataServiceApiCallStatisticsRequestListQuery {
	s.StartTime = &v
	return s
}

func (s *ListDataServiceApiCallStatisticsRequestListQuery) Validate() error {
	return dara.Validate(s)
}
