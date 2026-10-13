// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMetaProfilingResultCatalogPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *ListMetaProfilingResultCatalogPageRequest
	GetCatalogName() *string
	SetDbName(v string) *ListMetaProfilingResultCatalogPageRequest
	GetDbName() *string
	SetGroupBy(v string) *ListMetaProfilingResultCatalogPageRequest
	GetGroupBy() *string
	SetNeedTotalCount(v bool) *ListMetaProfilingResultCatalogPageRequest
	GetNeedTotalCount() *bool
	SetOrderBy(v string) *ListMetaProfilingResultCatalogPageRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ListMetaProfilingResultCatalogPageRequest
	GetOrderDirection() *string
	SetPageIndex(v int32) *ListMetaProfilingResultCatalogPageRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ListMetaProfilingResultCatalogPageRequest
	GetPageSize() *int32
	SetTaskId(v int64) *ListMetaProfilingResultCatalogPageRequest
	GetTaskId() *int64
}

type ListMetaProfilingResultCatalogPageRequest struct {
	// example:
	//
	// catalog_hive_prod
	CatalogName *string `json:"catalogName,omitempty" xml:"catalogName,omitempty"`
	// example:
	//
	// ods
	DbName *string `json:"dbName,omitempty" xml:"dbName,omitempty"`
	// example:
	//
	// dbName
	GroupBy        *string `json:"groupBy,omitempty" xml:"groupBy,omitempty"`
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
	// example:
	//
	// 100256
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s ListMetaProfilingResultCatalogPageRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMetaProfilingResultCatalogPageRequest) GoString() string {
	return s.String()
}

func (s *ListMetaProfilingResultCatalogPageRequest) GetCatalogName() *string {
	return s.CatalogName
}

func (s *ListMetaProfilingResultCatalogPageRequest) GetDbName() *string {
	return s.DbName
}

func (s *ListMetaProfilingResultCatalogPageRequest) GetGroupBy() *string {
	return s.GroupBy
}

func (s *ListMetaProfilingResultCatalogPageRequest) GetNeedTotalCount() *bool {
	return s.NeedTotalCount
}

func (s *ListMetaProfilingResultCatalogPageRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListMetaProfilingResultCatalogPageRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListMetaProfilingResultCatalogPageRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListMetaProfilingResultCatalogPageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMetaProfilingResultCatalogPageRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *ListMetaProfilingResultCatalogPageRequest) SetCatalogName(v string) *ListMetaProfilingResultCatalogPageRequest {
	s.CatalogName = &v
	return s
}

func (s *ListMetaProfilingResultCatalogPageRequest) SetDbName(v string) *ListMetaProfilingResultCatalogPageRequest {
	s.DbName = &v
	return s
}

func (s *ListMetaProfilingResultCatalogPageRequest) SetGroupBy(v string) *ListMetaProfilingResultCatalogPageRequest {
	s.GroupBy = &v
	return s
}

func (s *ListMetaProfilingResultCatalogPageRequest) SetNeedTotalCount(v bool) *ListMetaProfilingResultCatalogPageRequest {
	s.NeedTotalCount = &v
	return s
}

func (s *ListMetaProfilingResultCatalogPageRequest) SetOrderBy(v string) *ListMetaProfilingResultCatalogPageRequest {
	s.OrderBy = &v
	return s
}

func (s *ListMetaProfilingResultCatalogPageRequest) SetOrderDirection(v string) *ListMetaProfilingResultCatalogPageRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListMetaProfilingResultCatalogPageRequest) SetPageIndex(v int32) *ListMetaProfilingResultCatalogPageRequest {
	s.PageIndex = &v
	return s
}

func (s *ListMetaProfilingResultCatalogPageRequest) SetPageSize(v int32) *ListMetaProfilingResultCatalogPageRequest {
	s.PageSize = &v
	return s
}

func (s *ListMetaProfilingResultCatalogPageRequest) SetTaskId(v int64) *ListMetaProfilingResultCatalogPageRequest {
	s.TaskId = &v
	return s
}

func (s *ListMetaProfilingResultCatalogPageRequest) Validate() error {
	return dara.Validate(s)
}
