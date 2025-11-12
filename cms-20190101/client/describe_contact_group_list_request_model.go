// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContactGroupListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeContactGroupListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeContactGroupListRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeContactGroupListRequest
	GetRegionId() *string
}

type DescribeContactGroupListRequest struct {
	// Page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Number of records per page.
	//
	// example:
	//
	// 10
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeContactGroupListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeContactGroupListRequest) GoString() string {
	return s.String()
}

func (s *DescribeContactGroupListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeContactGroupListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeContactGroupListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeContactGroupListRequest) SetPageNumber(v int32) *DescribeContactGroupListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeContactGroupListRequest) SetPageSize(v int32) *DescribeContactGroupListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeContactGroupListRequest) SetRegionId(v string) *DescribeContactGroupListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeContactGroupListRequest) Validate() error {
	return dara.Validate(s)
}
