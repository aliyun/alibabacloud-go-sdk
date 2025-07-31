// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataServiceApiImpactsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQuery(v *ListDataServiceApiImpactsRequestListQuery) *ListDataServiceApiImpactsRequest
	GetListQuery() *ListDataServiceApiImpactsRequestListQuery
	SetOpTenantId(v int64) *ListDataServiceApiImpactsRequest
	GetOpTenantId() *int64
	SetProjectId(v int32) *ListDataServiceApiImpactsRequest
	GetProjectId() *int32
}

type ListDataServiceApiImpactsRequest struct {
	// This parameter is required.
	ListQuery *ListDataServiceApiImpactsRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 102102
	ProjectId *int32 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListDataServiceApiImpactsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceApiImpactsRequest) GoString() string {
	return s.String()
}

func (s *ListDataServiceApiImpactsRequest) GetListQuery() *ListDataServiceApiImpactsRequestListQuery {
	return s.ListQuery
}

func (s *ListDataServiceApiImpactsRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListDataServiceApiImpactsRequest) GetProjectId() *int32 {
	return s.ProjectId
}

func (s *ListDataServiceApiImpactsRequest) SetListQuery(v *ListDataServiceApiImpactsRequestListQuery) *ListDataServiceApiImpactsRequest {
	s.ListQuery = v
	return s
}

func (s *ListDataServiceApiImpactsRequest) SetOpTenantId(v int64) *ListDataServiceApiImpactsRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListDataServiceApiImpactsRequest) SetProjectId(v int32) *ListDataServiceApiImpactsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDataServiceApiImpactsRequest) Validate() error {
	return dara.Validate(s)
}

type ListDataServiceApiImpactsRequestListQuery struct {
	// This parameter is required.
	//
	// example:
	//
	// 1021
	ApiId *int64 `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// example:
	//
	// test
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2025-06-30 20:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// CALL_COUNT
	OrderColumn *string `json:"OrderColumn,omitempty" xml:"OrderColumn,omitempty"`
	// example:
	//
	// 1
	OrderType *int32 `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2025-06-30 08:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListDataServiceApiImpactsRequestListQuery) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceApiImpactsRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListDataServiceApiImpactsRequestListQuery) GetApiId() *int64 {
	return s.ApiId
}

func (s *ListDataServiceApiImpactsRequestListQuery) GetAppName() *string {
	return s.AppName
}

func (s *ListDataServiceApiImpactsRequestListQuery) GetEndTime() *string {
	return s.EndTime
}

func (s *ListDataServiceApiImpactsRequestListQuery) GetOrderColumn() *string {
	return s.OrderColumn
}

func (s *ListDataServiceApiImpactsRequestListQuery) GetOrderType() *int32 {
	return s.OrderType
}

func (s *ListDataServiceApiImpactsRequestListQuery) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListDataServiceApiImpactsRequestListQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataServiceApiImpactsRequestListQuery) GetStartTime() *string {
	return s.StartTime
}

func (s *ListDataServiceApiImpactsRequestListQuery) SetApiId(v int64) *ListDataServiceApiImpactsRequestListQuery {
	s.ApiId = &v
	return s
}

func (s *ListDataServiceApiImpactsRequestListQuery) SetAppName(v string) *ListDataServiceApiImpactsRequestListQuery {
	s.AppName = &v
	return s
}

func (s *ListDataServiceApiImpactsRequestListQuery) SetEndTime(v string) *ListDataServiceApiImpactsRequestListQuery {
	s.EndTime = &v
	return s
}

func (s *ListDataServiceApiImpactsRequestListQuery) SetOrderColumn(v string) *ListDataServiceApiImpactsRequestListQuery {
	s.OrderColumn = &v
	return s
}

func (s *ListDataServiceApiImpactsRequestListQuery) SetOrderType(v int32) *ListDataServiceApiImpactsRequestListQuery {
	s.OrderType = &v
	return s
}

func (s *ListDataServiceApiImpactsRequestListQuery) SetPageNo(v int32) *ListDataServiceApiImpactsRequestListQuery {
	s.PageNo = &v
	return s
}

func (s *ListDataServiceApiImpactsRequestListQuery) SetPageSize(v int32) *ListDataServiceApiImpactsRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListDataServiceApiImpactsRequestListQuery) SetStartTime(v string) *ListDataServiceApiImpactsRequestListQuery {
	s.StartTime = &v
	return s
}

func (s *ListDataServiceApiImpactsRequestListQuery) Validate() error {
	return dara.Validate(s)
}
