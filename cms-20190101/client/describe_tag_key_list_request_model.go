// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagKeyListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeTagKeyListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeTagKeyListRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeTagKeyListRequest
	GetRegionId() *string
}

type DescribeTagKeyListRequest struct {
	// The page number.
	//
	// Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Maximum value: 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeTagKeyListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagKeyListRequest) GoString() string {
	return s.String()
}

func (s *DescribeTagKeyListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeTagKeyListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeTagKeyListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTagKeyListRequest) SetPageNumber(v int32) *DescribeTagKeyListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeTagKeyListRequest) SetPageSize(v int32) *DescribeTagKeyListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTagKeyListRequest) SetRegionId(v string) *DescribeTagKeyListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTagKeyListRequest) Validate() error {
	return dara.Validate(s)
}
