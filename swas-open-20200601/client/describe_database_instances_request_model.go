// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDatabaseInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseInstanceIds(v string) *DescribeDatabaseInstancesRequest
	GetDatabaseInstanceIds() *string
	SetPageNumber(v int32) *DescribeDatabaseInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDatabaseInstancesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeDatabaseInstancesRequest
	GetRegionId() *string
}

type DescribeDatabaseInstancesRequest struct {
	// The IDs of the lightweight database instances. The value can be a JSON array that consists of multiple instance IDs. You can specify up to 100 instance IDs. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// ["swasdb-xxx******","swasdb-yyy******"]
	DatabaseInstanceIds *string `json:"DatabaseInstanceIds,omitempty" xml:"DatabaseInstanceIds,omitempty"`
	// The page number of the lightweight database instance list.
	//
	// Starting value: 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page for a paged query.
	//
	// Maximum value: 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the lightweight database instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDatabaseInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatabaseInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseInstancesRequest) GetDatabaseInstanceIds() *string {
	return s.DatabaseInstanceIds
}

func (s *DescribeDatabaseInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDatabaseInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDatabaseInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDatabaseInstancesRequest) SetDatabaseInstanceIds(v string) *DescribeDatabaseInstancesRequest {
	s.DatabaseInstanceIds = &v
	return s
}

func (s *DescribeDatabaseInstancesRequest) SetPageNumber(v int32) *DescribeDatabaseInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDatabaseInstancesRequest) SetPageSize(v int32) *DescribeDatabaseInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDatabaseInstancesRequest) SetRegionId(v string) *DescribeDatabaseInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDatabaseInstancesRequest) Validate() error {
	return dara.Validate(s)
}
