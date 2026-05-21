// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeTagsRequest
	GetRegionId() *string
	SetTagsOnly(v bool) *DescribeTagsRequest
	GetTagsOnly() *bool
}

type DescribeTagsRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// true
	TagsOnly *bool `json:"tagsOnly,omitempty" xml:"tagsOnly,omitempty"`
}

func (s DescribeTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsRequest) GoString() string {
	return s.String()
}

func (s *DescribeTagsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTagsRequest) GetTagsOnly() *bool {
	return s.TagsOnly
}

func (s *DescribeTagsRequest) SetRegionId(v string) *DescribeTagsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTagsRequest) SetTagsOnly(v bool) *DescribeTagsRequest {
	s.TagsOnly = &v
	return s
}

func (s *DescribeTagsRequest) Validate() error {
	return dara.Validate(s)
}
