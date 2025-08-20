// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAPSADBInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeAPSADBInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAPSADBInstancesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeAPSADBInstancesRequest
	GetRegionId() *string
}

type DescribeAPSADBInstancesRequest struct {
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeAPSADBInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAPSADBInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAPSADBInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAPSADBInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAPSADBInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAPSADBInstancesRequest) SetPageNumber(v int32) *DescribeAPSADBInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAPSADBInstancesRequest) SetPageSize(v int32) *DescribeAPSADBInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAPSADBInstancesRequest) SetRegionId(v string) *DescribeAPSADBInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAPSADBInstancesRequest) Validate() error {
	return dara.Validate(s)
}
