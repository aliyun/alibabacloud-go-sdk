// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagValueListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeTagValueListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeTagValueListRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeTagValueListRequest
	GetRegionId() *string
	SetTagKey(v string) *DescribeTagValueListRequest
	GetTagKey() *string
}

type DescribeTagValueListRequest struct {
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
	// Valid values: 1 to 100. Default value: 50.
	//
	// example:
	//
	// 10
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The tag key.
	//
	// For more information about how to obtain a tag key, see [DescribeTagKeyList](https://help.aliyun.com/document_detail/145558.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// tagKey1
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s DescribeTagValueListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagValueListRequest) GoString() string {
	return s.String()
}

func (s *DescribeTagValueListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeTagValueListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeTagValueListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTagValueListRequest) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeTagValueListRequest) SetPageNumber(v int32) *DescribeTagValueListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeTagValueListRequest) SetPageSize(v int32) *DescribeTagValueListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTagValueListRequest) SetRegionId(v string) *DescribeTagValueListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTagValueListRequest) SetTagKey(v string) *DescribeTagValueListRequest {
	s.TagKey = &v
	return s
}

func (s *DescribeTagValueListRequest) Validate() error {
	return dara.Validate(s)
}
