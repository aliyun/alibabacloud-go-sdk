// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTableListByTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeTableListByTypeRequest
	GetCurrentPage() *int32
	SetDbName(v string) *DescribeTableListByTypeRequest
	GetDbName() *string
	SetDrdsInstanceId(v string) *DescribeTableListByTypeRequest
	GetDrdsInstanceId() *string
	SetPageSize(v int32) *DescribeTableListByTypeRequest
	GetPageSize() *int32
	SetQuery(v string) *DescribeTableListByTypeRequest
	GetQuery() *string
	SetRegionId(v string) *DescribeTableListByTypeRequest
	GetRegionId() *string
	SetTableType(v string) *DescribeTableListByTypeRequest
	GetTableType() *string
}

type DescribeTableListByTypeRequest struct {
	// The number of the page to return.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The name of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds_flash****
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drdshbga76p6****
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The field that you specify for your query.
	//
	// example:
	//
	// drdshbga76p61861
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of tables. Valid values:
	//
	// This parameter is required.
	//
	// example:
	//
	// SINGLE
	TableType *string `json:"TableType,omitempty" xml:"TableType,omitempty"`
}

func (s DescribeTableListByTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTableListByTypeRequest) GoString() string {
	return s.String()
}

func (s *DescribeTableListByTypeRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeTableListByTypeRequest) GetDbName() *string {
	return s.DbName
}

func (s *DescribeTableListByTypeRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *DescribeTableListByTypeRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeTableListByTypeRequest) GetQuery() *string {
	return s.Query
}

func (s *DescribeTableListByTypeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTableListByTypeRequest) GetTableType() *string {
	return s.TableType
}

func (s *DescribeTableListByTypeRequest) SetCurrentPage(v int32) *DescribeTableListByTypeRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeTableListByTypeRequest) SetDbName(v string) *DescribeTableListByTypeRequest {
	s.DbName = &v
	return s
}

func (s *DescribeTableListByTypeRequest) SetDrdsInstanceId(v string) *DescribeTableListByTypeRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeTableListByTypeRequest) SetPageSize(v int32) *DescribeTableListByTypeRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTableListByTypeRequest) SetQuery(v string) *DescribeTableListByTypeRequest {
	s.Query = &v
	return s
}

func (s *DescribeTableListByTypeRequest) SetRegionId(v string) *DescribeTableListByTypeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTableListByTypeRequest) SetTableType(v string) *DescribeTableListByTypeRequest {
	s.TableType = &v
	return s
}

func (s *DescribeTableListByTypeRequest) Validate() error {
	return dara.Validate(s)
}
