// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserTagKeysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeUserTagKeysRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeUserTagKeysRequest
	GetNextToken() *string
	SetRegionId(v string) *DescribeUserTagKeysRequest
	GetRegionId() *string
	SetTagFilterKey(v string) *DescribeUserTagKeysRequest
	GetTagFilterKey() *string
}

type DescribeUserTagKeysRequest struct {
	// Number of items per page in paginated queries. The maximum value is 100.
	//
	// Default value:
	//
	// - If no value is set or the set value is less than 10, the default is 10.
	//
	// - If the set value is greater than 100, the default is 100.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The query token returned by this call (Token).
	//
	// example:
	//
	// f07b150eadfa1d7a
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region to which the resource belongs. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to view the latest list of Alibaba Cloud regions.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The tagKey for filtering the query.
	//
	// example:
	//
	// tagKey
	TagFilterKey *string `json:"TagFilterKey,omitempty" xml:"TagFilterKey,omitempty"`
}

func (s DescribeUserTagKeysRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserTagKeysRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserTagKeysRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeUserTagKeysRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeUserTagKeysRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeUserTagKeysRequest) GetTagFilterKey() *string {
	return s.TagFilterKey
}

func (s *DescribeUserTagKeysRequest) SetMaxResults(v int32) *DescribeUserTagKeysRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeUserTagKeysRequest) SetNextToken(v string) *DescribeUserTagKeysRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeUserTagKeysRequest) SetRegionId(v string) *DescribeUserTagKeysRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeUserTagKeysRequest) SetTagFilterKey(v string) *DescribeUserTagKeysRequest {
	s.TagFilterKey = &v
	return s
}

func (s *DescribeUserTagKeysRequest) Validate() error {
	return dara.Validate(s)
}
