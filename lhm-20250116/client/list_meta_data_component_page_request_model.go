// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMetaDataComponentPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryType(v string) *ListMetaDataComponentPageRequest
	GetCategoryType() *string
	SetComponentType(v int32) *ListMetaDataComponentPageRequest
	GetComponentType() *int32
	SetDsName(v string) *ListMetaDataComponentPageRequest
	GetDsName() *string
	SetDsStatus(v []*int32) *ListMetaDataComponentPageRequest
	GetDsStatus() []*int32
	SetDsType(v string) *ListMetaDataComponentPageRequest
	GetDsType() *string
	SetDsTypeList(v []*string) *ListMetaDataComponentPageRequest
	GetDsTypeList() []*string
	SetGroupBy(v string) *ListMetaDataComponentPageRequest
	GetGroupBy() *string
	SetNeedTotalCount(v string) *ListMetaDataComponentPageRequest
	GetNeedTotalCount() *string
	SetOrderBy(v string) *ListMetaDataComponentPageRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ListMetaDataComponentPageRequest
	GetOrderDirection() *string
	SetPageIndex(v int32) *ListMetaDataComponentPageRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ListMetaDataComponentPageRequest
	GetPageSize() *int32
	SetSrcComponentId(v int64) *ListMetaDataComponentPageRequest
	GetSrcComponentId() *int64
}

type ListMetaDataComponentPageRequest struct {
	// The category type of the data source. Valid values: DATASET, WORKFLOW, and ENGINE. For scheduling scenarios, this parameter is set to WORKFLOW.
	//
	// example:
	//
	// WORKFLOW
	CategoryType *string `json:"categoryType,omitempty" xml:"categoryType,omitempty"`
	// The entry component type. In some operations, this parameter is used as a backward compatible field for version 1.1.0. Valid values:
	//
	// - 0: source
	//
	// - 1: destination
	//
	// example:
	//
	// 0
	ComponentType *int32 `json:"componentType,omitempty" xml:"componentType,omitempty"`
	// The data source name. Exact match and fuzzy match are supported.
	//
	// example:
	//
	// test_ds318_hangzhou_0428
	DsName *string `json:"dsName,omitempty" xml:"dsName,omitempty"`
	// The connectivity status of the data source. Valid values:
	//
	// - 0: Not tested.
	//
	// - 1: Connected.
	//
	// - 2: Connection failed.
	//
	// - -1: Connectivity test not supported.
	DsStatus []*int32 `json:"dsStatus,omitempty" xml:"dsStatus,omitempty" type:"Repeated"`
	// The data source type, such as Hive or MaxCompute.
	//
	// example:
	//
	// Hive
	DsType *string `json:"dsType,omitempty" xml:"dsType,omitempty"`
	// The list of data source types.
	DsTypeList []*string `json:"dsTypeList,omitempty" xml:"dsTypeList,omitempty" type:"Repeated"`
	// The grouping field (GROUP BY condition). Set this parameter as needed.
	//
	// example:
	//
	// order_date
	GroupBy *string `json:"groupBy,omitempty" xml:"groupBy,omitempty"`
	// Specifies whether to return the total number of records in the paginated result.
	//
	// example:
	//
	// true
	NeedTotalCount *string `json:"needTotalCount,omitempty" xml:"needTotalCount,omitempty"`
	// The sort field. Set this parameter as needed.
	//
	// example:
	//
	// gmtCreate
	OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	// The sort direction. Valid values:
	//
	// - ASC: ascending order
	//
	// - DESC: descending order
	//
	// example:
	//
	// DESC
	OrderDirection *string `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	// The page number, starting from 1.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// The page size, which is the number of records returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The source component ID, which is the primary key of the source data source component.
	//
	// example:
	//
	// 12345
	SrcComponentId *int64 `json:"srcComponentId,omitempty" xml:"srcComponentId,omitempty"`
}

func (s ListMetaDataComponentPageRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMetaDataComponentPageRequest) GoString() string {
	return s.String()
}

func (s *ListMetaDataComponentPageRequest) GetCategoryType() *string {
	return s.CategoryType
}

func (s *ListMetaDataComponentPageRequest) GetComponentType() *int32 {
	return s.ComponentType
}

func (s *ListMetaDataComponentPageRequest) GetDsName() *string {
	return s.DsName
}

func (s *ListMetaDataComponentPageRequest) GetDsStatus() []*int32 {
	return s.DsStatus
}

func (s *ListMetaDataComponentPageRequest) GetDsType() *string {
	return s.DsType
}

func (s *ListMetaDataComponentPageRequest) GetDsTypeList() []*string {
	return s.DsTypeList
}

func (s *ListMetaDataComponentPageRequest) GetGroupBy() *string {
	return s.GroupBy
}

func (s *ListMetaDataComponentPageRequest) GetNeedTotalCount() *string {
	return s.NeedTotalCount
}

func (s *ListMetaDataComponentPageRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListMetaDataComponentPageRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListMetaDataComponentPageRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListMetaDataComponentPageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMetaDataComponentPageRequest) GetSrcComponentId() *int64 {
	return s.SrcComponentId
}

func (s *ListMetaDataComponentPageRequest) SetCategoryType(v string) *ListMetaDataComponentPageRequest {
	s.CategoryType = &v
	return s
}

func (s *ListMetaDataComponentPageRequest) SetComponentType(v int32) *ListMetaDataComponentPageRequest {
	s.ComponentType = &v
	return s
}

func (s *ListMetaDataComponentPageRequest) SetDsName(v string) *ListMetaDataComponentPageRequest {
	s.DsName = &v
	return s
}

func (s *ListMetaDataComponentPageRequest) SetDsStatus(v []*int32) *ListMetaDataComponentPageRequest {
	s.DsStatus = v
	return s
}

func (s *ListMetaDataComponentPageRequest) SetDsType(v string) *ListMetaDataComponentPageRequest {
	s.DsType = &v
	return s
}

func (s *ListMetaDataComponentPageRequest) SetDsTypeList(v []*string) *ListMetaDataComponentPageRequest {
	s.DsTypeList = v
	return s
}

func (s *ListMetaDataComponentPageRequest) SetGroupBy(v string) *ListMetaDataComponentPageRequest {
	s.GroupBy = &v
	return s
}

func (s *ListMetaDataComponentPageRequest) SetNeedTotalCount(v string) *ListMetaDataComponentPageRequest {
	s.NeedTotalCount = &v
	return s
}

func (s *ListMetaDataComponentPageRequest) SetOrderBy(v string) *ListMetaDataComponentPageRequest {
	s.OrderBy = &v
	return s
}

func (s *ListMetaDataComponentPageRequest) SetOrderDirection(v string) *ListMetaDataComponentPageRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListMetaDataComponentPageRequest) SetPageIndex(v int32) *ListMetaDataComponentPageRequest {
	s.PageIndex = &v
	return s
}

func (s *ListMetaDataComponentPageRequest) SetPageSize(v int32) *ListMetaDataComponentPageRequest {
	s.PageSize = &v
	return s
}

func (s *ListMetaDataComponentPageRequest) SetSrcComponentId(v int64) *ListMetaDataComponentPageRequest {
	s.SrcComponentId = &v
	return s
}

func (s *ListMetaDataComponentPageRequest) Validate() error {
	return dara.Validate(s)
}
