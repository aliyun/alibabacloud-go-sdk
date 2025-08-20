// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetViewObjectsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *GetViewObjectsRequest
	GetDBClusterId() *string
	SetFilterOwner(v string) *GetViewObjectsRequest
	GetFilterOwner() *string
	SetFilterViewName(v string) *GetViewObjectsRequest
	GetFilterViewName() *string
	SetFilterViewType(v string) *GetViewObjectsRequest
	GetFilterViewType() *string
	SetOrderBy(v string) *GetViewObjectsRequest
	GetOrderBy() *string
	SetPageNumber(v int64) *GetViewObjectsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *GetViewObjectsRequest
	GetPageSize() *int64
	SetRegionId(v string) *GetViewObjectsRequest
	GetRegionId() *string
	SetSchemaName(v string) *GetViewObjectsRequest
	GetSchemaName() *string
	SetShowMvBaseTable(v bool) *GetViewObjectsRequest
	GetShowMvBaseTable() *bool
}

type GetViewObjectsRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp1xxxxxxxx47
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The owner of the view.
	//
	// example:
	//
	// admin
	FilterOwner *string `json:"FilterOwner,omitempty" xml:"FilterOwner,omitempty"`
	// The name of the view.
	//
	// example:
	//
	// test_filter
	FilterViewName *string `json:"FilterViewName,omitempty" xml:"FilterViewName,omitempty"`
	// The type of the view.
	//
	// Valid values:
	//
	// \\-VIRTUAL_VIEW
	//
	// \\-MATERIALIZED_VIEW
	//
	// Default value: null.
	//
	// example:
	//
	// VIRTUAL_VIEW
	FilterViewType *string `json:"FilterViewType,omitempty" xml:"FilterViewType,omitempty"`
	// The order in which you want to sort the query results. Valid values for Type:
	//
	// 	- Asc
	//
	// 	- Desc
	//
	// Valid values for Field: -ViewName
	//
	// \\-CreateTime
	//
	// \\-UpdateTime
	//
	// Default value: {"Type": "Desc","Field": "ViewName"}.
	//
	// example:
	//
	// {"Type": "Desc","Field": "ViewName"}
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// 	- **30*	- (default)
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// example:
	//
	// 30
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// adb_demo
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// example:
	//
	// true
	ShowMvBaseTable *bool `json:"ShowMvBaseTable,omitempty" xml:"ShowMvBaseTable,omitempty"`
}

func (s GetViewObjectsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetViewObjectsRequest) GoString() string {
	return s.String()
}

func (s *GetViewObjectsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *GetViewObjectsRequest) GetFilterOwner() *string {
	return s.FilterOwner
}

func (s *GetViewObjectsRequest) GetFilterViewName() *string {
	return s.FilterViewName
}

func (s *GetViewObjectsRequest) GetFilterViewType() *string {
	return s.FilterViewType
}

func (s *GetViewObjectsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *GetViewObjectsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *GetViewObjectsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetViewObjectsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetViewObjectsRequest) GetSchemaName() *string {
	return s.SchemaName
}

func (s *GetViewObjectsRequest) GetShowMvBaseTable() *bool {
	return s.ShowMvBaseTable
}

func (s *GetViewObjectsRequest) SetDBClusterId(v string) *GetViewObjectsRequest {
	s.DBClusterId = &v
	return s
}

func (s *GetViewObjectsRequest) SetFilterOwner(v string) *GetViewObjectsRequest {
	s.FilterOwner = &v
	return s
}

func (s *GetViewObjectsRequest) SetFilterViewName(v string) *GetViewObjectsRequest {
	s.FilterViewName = &v
	return s
}

func (s *GetViewObjectsRequest) SetFilterViewType(v string) *GetViewObjectsRequest {
	s.FilterViewType = &v
	return s
}

func (s *GetViewObjectsRequest) SetOrderBy(v string) *GetViewObjectsRequest {
	s.OrderBy = &v
	return s
}

func (s *GetViewObjectsRequest) SetPageNumber(v int64) *GetViewObjectsRequest {
	s.PageNumber = &v
	return s
}

func (s *GetViewObjectsRequest) SetPageSize(v int64) *GetViewObjectsRequest {
	s.PageSize = &v
	return s
}

func (s *GetViewObjectsRequest) SetRegionId(v string) *GetViewObjectsRequest {
	s.RegionId = &v
	return s
}

func (s *GetViewObjectsRequest) SetSchemaName(v string) *GetViewObjectsRequest {
	s.SchemaName = &v
	return s
}

func (s *GetViewObjectsRequest) SetShowMvBaseTable(v bool) *GetViewObjectsRequest {
	s.ShowMvBaseTable = &v
	return s
}

func (s *GetViewObjectsRequest) Validate() error {
	return dara.Validate(s)
}
