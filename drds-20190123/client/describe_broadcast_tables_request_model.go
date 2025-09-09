// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBroadcastTablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeBroadcastTablesRequest
	GetCurrentPage() *int32
	SetDbName(v string) *DescribeBroadcastTablesRequest
	GetDbName() *string
	SetDrdsInstanceId(v string) *DescribeBroadcastTablesRequest
	GetDrdsInstanceId() *string
	SetPageSize(v int32) *DescribeBroadcastTablesRequest
	GetPageSize() *int32
	SetQuery(v string) *DescribeBroadcastTablesRequest
	GetQuery() *string
	SetRegionId(v string) *DescribeBroadcastTablesRequest
	GetRegionId() *string
}

type DescribeBroadcastTablesRequest struct {
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
	// test_db
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds************
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The content of the query.
	//
	// example:
	//
	// tb1
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeBroadcastTablesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBroadcastTablesRequest) GoString() string {
	return s.String()
}

func (s *DescribeBroadcastTablesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeBroadcastTablesRequest) GetDbName() *string {
	return s.DbName
}

func (s *DescribeBroadcastTablesRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *DescribeBroadcastTablesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBroadcastTablesRequest) GetQuery() *string {
	return s.Query
}

func (s *DescribeBroadcastTablesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeBroadcastTablesRequest) SetCurrentPage(v int32) *DescribeBroadcastTablesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeBroadcastTablesRequest) SetDbName(v string) *DescribeBroadcastTablesRequest {
	s.DbName = &v
	return s
}

func (s *DescribeBroadcastTablesRequest) SetDrdsInstanceId(v string) *DescribeBroadcastTablesRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeBroadcastTablesRequest) SetPageSize(v int32) *DescribeBroadcastTablesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBroadcastTablesRequest) SetQuery(v string) *DescribeBroadcastTablesRequest {
	s.Query = &v
	return s
}

func (s *DescribeBroadcastTablesRequest) SetRegionId(v string) *DescribeBroadcastTablesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeBroadcastTablesRequest) Validate() error {
	return dara.Validate(s)
}
