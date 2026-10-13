// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMetaProfilingResultUdfPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *ListMetaProfilingResultUdfPageRequest
	GetCatalogName() *string
	SetComponentType(v int32) *ListMetaProfilingResultUdfPageRequest
	GetComponentType() *int32
	SetDbName(v string) *ListMetaProfilingResultUdfPageRequest
	GetDbName() *string
	SetFuncName(v string) *ListMetaProfilingResultUdfPageRequest
	GetFuncName() *string
	SetGroupBy(v string) *ListMetaProfilingResultUdfPageRequest
	GetGroupBy() *string
	SetIsSummary(v bool) *ListMetaProfilingResultUdfPageRequest
	GetIsSummary() *bool
	SetNeedTotalCount(v bool) *ListMetaProfilingResultUdfPageRequest
	GetNeedTotalCount() *bool
	SetOrderBy(v string) *ListMetaProfilingResultUdfPageRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ListMetaProfilingResultUdfPageRequest
	GetOrderDirection() *string
	SetPageIndex(v int32) *ListMetaProfilingResultUdfPageRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ListMetaProfilingResultUdfPageRequest
	GetPageSize() *int32
	SetTaskId(v int64) *ListMetaProfilingResultUdfPageRequest
	GetTaskId() *int64
}

type ListMetaProfilingResultUdfPageRequest struct {
	// example:
	//
	// catalog_hive_prod
	CatalogName *string `json:"catalogName,omitempty" xml:"catalogName,omitempty"`
	// example:
	//
	// 1
	ComponentType *int32 `json:"componentType,omitempty" xml:"componentType,omitempty"`
	// example:
	//
	// ods
	DbName *string `json:"dbName,omitempty" xml:"dbName,omitempty"`
	// example:
	//
	// udf_parse_json
	FuncName *string `json:"funcName,omitempty" xml:"funcName,omitempty"`
	// example:
	//
	// dbName
	GroupBy        *string `json:"groupBy,omitempty" xml:"groupBy,omitempty"`
	IsSummary      *bool   `json:"isSummary,omitempty" xml:"isSummary,omitempty"`
	NeedTotalCount *bool   `json:"needTotalCount,omitempty" xml:"needTotalCount,omitempty"`
	// example:
	//
	// gmtCreate
	OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	// example:
	//
	// DESC
	OrderDirection *string `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100256
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s ListMetaProfilingResultUdfPageRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMetaProfilingResultUdfPageRequest) GoString() string {
	return s.String()
}

func (s *ListMetaProfilingResultUdfPageRequest) GetCatalogName() *string {
	return s.CatalogName
}

func (s *ListMetaProfilingResultUdfPageRequest) GetComponentType() *int32 {
	return s.ComponentType
}

func (s *ListMetaProfilingResultUdfPageRequest) GetDbName() *string {
	return s.DbName
}

func (s *ListMetaProfilingResultUdfPageRequest) GetFuncName() *string {
	return s.FuncName
}

func (s *ListMetaProfilingResultUdfPageRequest) GetGroupBy() *string {
	return s.GroupBy
}

func (s *ListMetaProfilingResultUdfPageRequest) GetIsSummary() *bool {
	return s.IsSummary
}

func (s *ListMetaProfilingResultUdfPageRequest) GetNeedTotalCount() *bool {
	return s.NeedTotalCount
}

func (s *ListMetaProfilingResultUdfPageRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListMetaProfilingResultUdfPageRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListMetaProfilingResultUdfPageRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListMetaProfilingResultUdfPageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMetaProfilingResultUdfPageRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *ListMetaProfilingResultUdfPageRequest) SetCatalogName(v string) *ListMetaProfilingResultUdfPageRequest {
	s.CatalogName = &v
	return s
}

func (s *ListMetaProfilingResultUdfPageRequest) SetComponentType(v int32) *ListMetaProfilingResultUdfPageRequest {
	s.ComponentType = &v
	return s
}

func (s *ListMetaProfilingResultUdfPageRequest) SetDbName(v string) *ListMetaProfilingResultUdfPageRequest {
	s.DbName = &v
	return s
}

func (s *ListMetaProfilingResultUdfPageRequest) SetFuncName(v string) *ListMetaProfilingResultUdfPageRequest {
	s.FuncName = &v
	return s
}

func (s *ListMetaProfilingResultUdfPageRequest) SetGroupBy(v string) *ListMetaProfilingResultUdfPageRequest {
	s.GroupBy = &v
	return s
}

func (s *ListMetaProfilingResultUdfPageRequest) SetIsSummary(v bool) *ListMetaProfilingResultUdfPageRequest {
	s.IsSummary = &v
	return s
}

func (s *ListMetaProfilingResultUdfPageRequest) SetNeedTotalCount(v bool) *ListMetaProfilingResultUdfPageRequest {
	s.NeedTotalCount = &v
	return s
}

func (s *ListMetaProfilingResultUdfPageRequest) SetOrderBy(v string) *ListMetaProfilingResultUdfPageRequest {
	s.OrderBy = &v
	return s
}

func (s *ListMetaProfilingResultUdfPageRequest) SetOrderDirection(v string) *ListMetaProfilingResultUdfPageRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListMetaProfilingResultUdfPageRequest) SetPageIndex(v int32) *ListMetaProfilingResultUdfPageRequest {
	s.PageIndex = &v
	return s
}

func (s *ListMetaProfilingResultUdfPageRequest) SetPageSize(v int32) *ListMetaProfilingResultUdfPageRequest {
	s.PageSize = &v
	return s
}

func (s *ListMetaProfilingResultUdfPageRequest) SetTaskId(v int64) *ListMetaProfilingResultUdfPageRequest {
	s.TaskId = &v
	return s
}

func (s *ListMetaProfilingResultUdfPageRequest) Validate() error {
	return dara.Validate(s)
}
