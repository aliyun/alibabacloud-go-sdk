// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserTagValuesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeUserTagValuesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeUserTagValuesRequest
	GetNextToken() *string
	SetRegionId(v string) *DescribeUserTagValuesRequest
	GetRegionId() *string
	SetTagFilterValue(v string) *DescribeUserTagValuesRequest
	GetTagFilterValue() *string
	SetTagKey(v string) *DescribeUserTagValuesRequest
	GetTagKey() *string
}

type DescribeUserTagValuesRequest struct {
	// Number of items per page in a paginated query. The maximum value is 100.
	//
	// Default value:
	//
	// - If no value is set or the set value is less than 10, the default value is 10.
	//
	// - If the set value is greater than 100, the default value is 100.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Query token (Token). The value should be the NextToken parameter value from the previous call to this interface. This parameter is not required for the initial call. If NextToken is set, the PageSize and PageNumber request parameters become invalid, and the TotalCount in the response data is also invalid.
	//
	// example:
	//
	// AAAAAdDWBF2
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID of the consistency replication group.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Tag content filter
	//
	// example:
	//
	// keyValue
	TagFilterValue *string `json:"TagFilterValue,omitempty" xml:"TagFilterValue,omitempty"`
	// Tag key.
	//
	// example:
	//
	// TestKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s DescribeUserTagValuesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserTagValuesRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserTagValuesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeUserTagValuesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeUserTagValuesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeUserTagValuesRequest) GetTagFilterValue() *string {
	return s.TagFilterValue
}

func (s *DescribeUserTagValuesRequest) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeUserTagValuesRequest) SetMaxResults(v int32) *DescribeUserTagValuesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeUserTagValuesRequest) SetNextToken(v string) *DescribeUserTagValuesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeUserTagValuesRequest) SetRegionId(v string) *DescribeUserTagValuesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeUserTagValuesRequest) SetTagFilterValue(v string) *DescribeUserTagValuesRequest {
	s.TagFilterValue = &v
	return s
}

func (s *DescribeUserTagValuesRequest) SetTagKey(v string) *DescribeUserTagValuesRequest {
	s.TagKey = &v
	return s
}

func (s *DescribeUserTagValuesRequest) Validate() error {
	return dara.Validate(s)
}
