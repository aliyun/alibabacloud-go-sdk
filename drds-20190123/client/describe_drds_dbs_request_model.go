// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsDBsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDrdsInstanceId(v string) *DescribeDrdsDBsRequest
	GetDrdsInstanceId() *string
	SetPageNumber(v int32) *DescribeDrdsDBsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDrdsDBsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeDrdsDBsRequest
	GetRegionId() *string
}

type DescribeDrdsDBsRequest struct {
	// The ID of the PolarDB-X 1.0 instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drdshbga1138****
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The number of the page to return. The value of this parameter must be an integer that is greater than 0. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of databases to return on each page. Valid values: **30**, **50**, and **100**.
	//
	// Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region in which the PolarDB-X 1.0 instance is created.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDrdsDBsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsDBsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBsRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *DescribeDrdsDBsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDrdsDBsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDrdsDBsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDrdsDBsRequest) SetDrdsInstanceId(v string) *DescribeDrdsDBsRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsDBsRequest) SetPageNumber(v int32) *DescribeDrdsDBsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDrdsDBsRequest) SetPageSize(v int32) *DescribeDrdsDBsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDrdsDBsRequest) SetRegionId(v string) *DescribeDrdsDBsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDrdsDBsRequest) Validate() error {
	return dara.Validate(s)
}
