// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsDbInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *DescribeDrdsDbInstancesRequest
	GetDbName() *string
	SetDrdsInstanceId(v string) *DescribeDrdsDbInstancesRequest
	GetDrdsInstanceId() *string
	SetPageNumber(v int32) *DescribeDrdsDbInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDrdsDbInstancesRequest
	GetPageSize() *int32
}

type DescribeDrdsDbInstancesRequest struct {
	// The name of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// dbname
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drdshbga1138****
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDrdsDbInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsDbInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbInstancesRequest) GetDbName() *string {
	return s.DbName
}

func (s *DescribeDrdsDbInstancesRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *DescribeDrdsDbInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDrdsDbInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDrdsDbInstancesRequest) SetDbName(v string) *DescribeDrdsDbInstancesRequest {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsDbInstancesRequest) SetDrdsInstanceId(v string) *DescribeDrdsDbInstancesRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsDbInstancesRequest) SetPageNumber(v int32) *DescribeDrdsDbInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDrdsDbInstancesRequest) SetPageSize(v int32) *DescribeDrdsDbInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDrdsDbInstancesRequest) Validate() error {
	return dara.Validate(s)
}
