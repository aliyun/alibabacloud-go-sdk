// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDynamicTagRuleListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDynamicTagRuleId(v string) *DescribeDynamicTagRuleListRequest
	GetDynamicTagRuleId() *string
	SetPageNumber(v string) *DescribeDynamicTagRuleListRequest
	GetPageNumber() *string
	SetPageSize(v string) *DescribeDynamicTagRuleListRequest
	GetPageSize() *string
	SetTagKey(v string) *DescribeDynamicTagRuleListRequest
	GetTagKey() *string
	SetTagRegionId(v string) *DescribeDynamicTagRuleListRequest
	GetTagRegionId() *string
	SetTagValue(v string) *DescribeDynamicTagRuleListRequest
	GetTagValue() *string
}

type DescribeDynamicTagRuleListRequest struct {
	// The ID of the tag rule.
	//
	// example:
	//
	// 004155fa-15ba-466d-b61a-***********
	DynamicTagRuleId *string `json:"DynamicTagRuleId,omitempty" xml:"DynamicTagRuleId,omitempty"`
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Minimum value: 1. Default value: 30.
	//
	// example:
	//
	// 30
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The tag key.
	//
	// For more information about how to obtain a tag key, see [DescribeTagKeyList](https://help.aliyun.com/document_detail/145558.html).
	//
	// example:
	//
	// tagkey1
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The ID of the region to which the tags belong.
	//
	// example:
	//
	// cn-hangzhou
	TagRegionId *string `json:"TagRegionId,omitempty" xml:"TagRegionId,omitempty"`
	// The tag value.
	//
	// For more information about how to obtain a tag value, see [DescribeTagKeyList](https://help.aliyun.com/document_detail/145557.html).
	//
	// example:
	//
	// *
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeDynamicTagRuleListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDynamicTagRuleListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDynamicTagRuleListRequest) GetDynamicTagRuleId() *string {
	return s.DynamicTagRuleId
}

func (s *DescribeDynamicTagRuleListRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeDynamicTagRuleListRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeDynamicTagRuleListRequest) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeDynamicTagRuleListRequest) GetTagRegionId() *string {
	return s.TagRegionId
}

func (s *DescribeDynamicTagRuleListRequest) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeDynamicTagRuleListRequest) SetDynamicTagRuleId(v string) *DescribeDynamicTagRuleListRequest {
	s.DynamicTagRuleId = &v
	return s
}

func (s *DescribeDynamicTagRuleListRequest) SetPageNumber(v string) *DescribeDynamicTagRuleListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDynamicTagRuleListRequest) SetPageSize(v string) *DescribeDynamicTagRuleListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDynamicTagRuleListRequest) SetTagKey(v string) *DescribeDynamicTagRuleListRequest {
	s.TagKey = &v
	return s
}

func (s *DescribeDynamicTagRuleListRequest) SetTagRegionId(v string) *DescribeDynamicTagRuleListRequest {
	s.TagRegionId = &v
	return s
}

func (s *DescribeDynamicTagRuleListRequest) SetTagValue(v string) *DescribeDynamicTagRuleListRequest {
	s.TagValue = &v
	return s
}

func (s *DescribeDynamicTagRuleListRequest) Validate() error {
	return dara.Validate(s)
}
