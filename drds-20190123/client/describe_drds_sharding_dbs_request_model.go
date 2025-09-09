// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsShardingDbsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *DescribeDrdsShardingDbsRequest
	GetDbName() *string
	SetDbNamePattern(v string) *DescribeDrdsShardingDbsRequest
	GetDbNamePattern() *string
	SetDrdsInstanceId(v string) *DescribeDrdsShardingDbsRequest
	GetDrdsInstanceId() *string
	SetPageNumber(v int64) *DescribeDrdsShardingDbsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeDrdsShardingDbsRequest
	GetPageSize() *int64
}

type DescribeDrdsShardingDbsRequest struct {
	// The name of the database whose shards you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds_test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The matching pattern of the database name.
	//
	// example:
	//
	// test
	DbNamePattern *string `json:"DbNamePattern,omitempty" xml:"DbNamePattern,omitempty"`
	// The ID of the PolarDB-X 1.0 instance whose database shards you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// drdshbgaf3c63qbo
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of database shards returned on each page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDrdsShardingDbsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsShardingDbsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsShardingDbsRequest) GetDbName() *string {
	return s.DbName
}

func (s *DescribeDrdsShardingDbsRequest) GetDbNamePattern() *string {
	return s.DbNamePattern
}

func (s *DescribeDrdsShardingDbsRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *DescribeDrdsShardingDbsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeDrdsShardingDbsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeDrdsShardingDbsRequest) SetDbName(v string) *DescribeDrdsShardingDbsRequest {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsShardingDbsRequest) SetDbNamePattern(v string) *DescribeDrdsShardingDbsRequest {
	s.DbNamePattern = &v
	return s
}

func (s *DescribeDrdsShardingDbsRequest) SetDrdsInstanceId(v string) *DescribeDrdsShardingDbsRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsShardingDbsRequest) SetPageNumber(v int64) *DescribeDrdsShardingDbsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDrdsShardingDbsRequest) SetPageSize(v int64) *DescribeDrdsShardingDbsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDrdsShardingDbsRequest) Validate() error {
	return dara.Validate(s)
}
