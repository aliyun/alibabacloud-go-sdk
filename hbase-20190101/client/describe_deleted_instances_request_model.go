// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeletedInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeDeletedInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDeletedInstancesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeDeletedInstancesRequest
	GetRegionId() *string
}

type DescribeDeletedInstancesRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDeletedInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeletedInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDeletedInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDeletedInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDeletedInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDeletedInstancesRequest) SetPageNumber(v int32) *DescribeDeletedInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDeletedInstancesRequest) SetPageSize(v int32) *DescribeDeletedInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDeletedInstancesRequest) SetRegionId(v string) *DescribeDeletedInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDeletedInstancesRequest) Validate() error {
	return dara.Validate(s)
}
