// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableObjectsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *GetTableObjectsRequest
	GetDBClusterId() *string
	SetFilterDescription(v string) *GetTableObjectsRequest
	GetFilterDescription() *string
	SetFilterOwner(v string) *GetTableObjectsRequest
	GetFilterOwner() *string
	SetFilterTblName(v string) *GetTableObjectsRequest
	GetFilterTblName() *string
	SetFilterTblType(v string) *GetTableObjectsRequest
	GetFilterTblType() *string
	SetOrderBy(v string) *GetTableObjectsRequest
	GetOrderBy() *string
	SetPageNumber(v int64) *GetTableObjectsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *GetTableObjectsRequest
	GetPageSize() *int64
	SetRegionId(v string) *GetTableObjectsRequest
	GetRegionId() *string
	SetSchemaName(v string) *GetTableObjectsRequest
	GetSchemaName() *string
}

type GetTableObjectsRequest struct {
	// The ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp1565u55p32****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The description of the table.
	//
	// example:
	//
	// description
	FilterDescription *string `json:"FilterDescription,omitempty" xml:"FilterDescription,omitempty"`
	// The owner of the table.
	//
	// example:
	//
	// admin
	FilterOwner *string `json:"FilterOwner,omitempty" xml:"FilterOwner,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// test_tbl
	FilterTblName *string `json:"FilterTblName,omitempty" xml:"FilterTblName,omitempty"`
	// The type of the table.
	//
	// Valid values:
	//
	// DIMENSION_TABLE
	//
	// FACT_TABLE
	//
	// EXTERNAL_TABLE
	//
	// Default value: null.
	//
	// example:
	//
	// FACT_TABLE
	FilterTblType *string `json:"FilterTblType,omitempty" xml:"FilterTblType,omitempty"`
	// The order in which the fields to be returned are sorted.
	//
	// Valid values:
	//
	// 	- Asc
	//
	// 	- Desc
	//
	// Values for fields:
	//
	// TableName
	//
	// TableSize
	//
	// CreateTime
	//
	// UpdateTime
	//
	// Default value: {"Type": "Desc","Field": "TableName"};
	//
	// example:
	//
	// {"Type": "Desc","Field": "TableName"}
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The number of the page to return. The value is an integer that is greater than 0. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values:
	//
	// 	- 30
	//
	// 	- 50
	//
	// 	- 100
	//
	// Default value: 30.
	//
	// example:
	//
	// 30
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region in which the cluster resides.
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
}

func (s GetTableObjectsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTableObjectsRequest) GoString() string {
	return s.String()
}

func (s *GetTableObjectsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *GetTableObjectsRequest) GetFilterDescription() *string {
	return s.FilterDescription
}

func (s *GetTableObjectsRequest) GetFilterOwner() *string {
	return s.FilterOwner
}

func (s *GetTableObjectsRequest) GetFilterTblName() *string {
	return s.FilterTblName
}

func (s *GetTableObjectsRequest) GetFilterTblType() *string {
	return s.FilterTblType
}

func (s *GetTableObjectsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *GetTableObjectsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *GetTableObjectsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetTableObjectsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetTableObjectsRequest) GetSchemaName() *string {
	return s.SchemaName
}

func (s *GetTableObjectsRequest) SetDBClusterId(v string) *GetTableObjectsRequest {
	s.DBClusterId = &v
	return s
}

func (s *GetTableObjectsRequest) SetFilterDescription(v string) *GetTableObjectsRequest {
	s.FilterDescription = &v
	return s
}

func (s *GetTableObjectsRequest) SetFilterOwner(v string) *GetTableObjectsRequest {
	s.FilterOwner = &v
	return s
}

func (s *GetTableObjectsRequest) SetFilterTblName(v string) *GetTableObjectsRequest {
	s.FilterTblName = &v
	return s
}

func (s *GetTableObjectsRequest) SetFilterTblType(v string) *GetTableObjectsRequest {
	s.FilterTblType = &v
	return s
}

func (s *GetTableObjectsRequest) SetOrderBy(v string) *GetTableObjectsRequest {
	s.OrderBy = &v
	return s
}

func (s *GetTableObjectsRequest) SetPageNumber(v int64) *GetTableObjectsRequest {
	s.PageNumber = &v
	return s
}

func (s *GetTableObjectsRequest) SetPageSize(v int64) *GetTableObjectsRequest {
	s.PageSize = &v
	return s
}

func (s *GetTableObjectsRequest) SetRegionId(v string) *GetTableObjectsRequest {
	s.RegionId = &v
	return s
}

func (s *GetTableObjectsRequest) SetSchemaName(v string) *GetTableObjectsRequest {
	s.SchemaName = &v
	return s
}

func (s *GetTableObjectsRequest) Validate() error {
	return dara.Validate(s)
}
