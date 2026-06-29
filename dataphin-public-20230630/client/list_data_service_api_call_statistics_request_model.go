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
	SetProjectId(v int32) *ListDataServiceApiCallStatisticsRequest
	GetProjectId() *int32
}

type ListDataServiceApiCallStatisticsRequest struct {
	// Search conditions.
	//
	// This parameter is required.
	ListQuery *ListDataServiceApiCallStatisticsRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// Tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// Data service project ID.
	//
	// This parameter is required.
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
	// End time. Format: yyyy-MM-dd HH:mm:ss.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2025-06-30 20:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Search keyword.
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// Sort column. Valid values: CALL_COUNT (call count), ERROR_COUNT (error count), ERROR_RATE (error rate), AVG_RESPONSE_TIME (average response time), and OFFLINE_RATE (offline percentage). Default value: CALL_COUNT.
	//
	// example:
	//
	// CALL_COUNT
	OrderColumn *string `json:"OrderColumn,omitempty" xml:"OrderColumn,omitempty"`
	// Sort order. Valid values: 1 (ascending) and 2 (descending). Default value: 1.
	//
	// example:
	//
	// 1
	OrderType *int32 `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// Page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// Number of entries per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Start time. Format: yyyy-MM-dd HH:mm:ss.
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
