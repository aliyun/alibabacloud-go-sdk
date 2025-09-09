// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDbInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbInstType(v string) *DescribeDbInstancesRequest
	GetDbInstType() *string
	SetDrdsInstanceId(v string) *DescribeDbInstancesRequest
	GetDrdsInstanceId() *string
	SetPageNumber(v int32) *DescribeDbInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDbInstancesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeDbInstancesRequest
	GetRegionId() *string
	SetSearch(v string) *DescribeDbInstancesRequest
	GetSearch() *string
}

type DescribeDbInstancesRequest struct {
	// Storage layer type. Valid values: **POLARDB*	- or **RDS**.
	//
	// example:
	//
	// POLARDB
	DbInstType *string `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	// The ID of a DRDS instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds************
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the storage or cluster.
	//
	// example:
	//
	// pc-***************
	Search *string `json:"Search,omitempty" xml:"Search,omitempty"`
}

func (s DescribeDbInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDbInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDbInstancesRequest) GetDbInstType() *string {
	return s.DbInstType
}

func (s *DescribeDbInstancesRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *DescribeDbInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDbInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDbInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDbInstancesRequest) GetSearch() *string {
	return s.Search
}

func (s *DescribeDbInstancesRequest) SetDbInstType(v string) *DescribeDbInstancesRequest {
	s.DbInstType = &v
	return s
}

func (s *DescribeDbInstancesRequest) SetDrdsInstanceId(v string) *DescribeDbInstancesRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDbInstancesRequest) SetPageNumber(v int32) *DescribeDbInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDbInstancesRequest) SetPageSize(v int32) *DescribeDbInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDbInstancesRequest) SetRegionId(v string) *DescribeDbInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDbInstancesRequest) SetSearch(v string) *DescribeDbInstancesRequest {
	s.Search = &v
	return s
}

func (s *DescribeDbInstancesRequest) Validate() error {
	return dara.Validate(s)
}
