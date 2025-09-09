// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeTablesRequest
	GetCurrentPage() *int32
	SetDbName(v string) *DescribeTablesRequest
	GetDbName() *string
	SetDrdsInstanceId(v string) *DescribeTablesRequest
	GetDrdsInstanceId() *string
	SetPageSize(v int32) *DescribeTablesRequest
	GetPageSize() *int32
	SetQuery(v string) *DescribeTablesRequest
	GetQuery() *string
	SetRegionId(v string) *DescribeTablesRequest
	GetRegionId() *string
}

type DescribeTablesRequest struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The name of the database whose tables you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds************
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The number of tables returned on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The query condition. The value of this parameter is the ID of the PolarDB-X 1.0 instance.
	//
	// example:
	//
	// drds************
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeTablesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTablesRequest) GoString() string {
	return s.String()
}

func (s *DescribeTablesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeTablesRequest) GetDbName() *string {
	return s.DbName
}

func (s *DescribeTablesRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *DescribeTablesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeTablesRequest) GetQuery() *string {
	return s.Query
}

func (s *DescribeTablesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTablesRequest) SetCurrentPage(v int32) *DescribeTablesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeTablesRequest) SetDbName(v string) *DescribeTablesRequest {
	s.DbName = &v
	return s
}

func (s *DescribeTablesRequest) SetDrdsInstanceId(v string) *DescribeTablesRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeTablesRequest) SetPageSize(v int32) *DescribeTablesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTablesRequest) SetQuery(v string) *DescribeTablesRequest {
	s.Query = &v
	return s
}

func (s *DescribeTablesRequest) SetRegionId(v string) *DescribeTablesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTablesRequest) Validate() error {
	return dara.Validate(s)
}
