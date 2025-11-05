// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserTagValuesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeUserTagValuesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeUserTagValuesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeUserTagValuesResponseBody
	GetRequestId() *string
	SetTagValues(v []*string) *DescribeUserTagValuesResponseBody
	GetTagValues() []*string
}

type DescribeUserTagValuesResponseBody struct {
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
	// NextToken
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Request ID. We return the request ID regardless of whether the API call was successful or not.
	//
	// example:
	//
	// C123F94F-4E38-19AE-942A-A8D6F44F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Tag values corresponding to the tag key.
	TagValues []*string `json:"TagValues,omitempty" xml:"TagValues,omitempty" type:"Repeated"`
}

func (s DescribeUserTagValuesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserTagValuesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserTagValuesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeUserTagValuesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeUserTagValuesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserTagValuesResponseBody) GetTagValues() []*string {
	return s.TagValues
}

func (s *DescribeUserTagValuesResponseBody) SetMaxResults(v int32) *DescribeUserTagValuesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeUserTagValuesResponseBody) SetNextToken(v string) *DescribeUserTagValuesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeUserTagValuesResponseBody) SetRequestId(v string) *DescribeUserTagValuesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserTagValuesResponseBody) SetTagValues(v []*string) *DescribeUserTagValuesResponseBody {
	s.TagValues = v
	return s
}

func (s *DescribeUserTagValuesResponseBody) Validate() error {
	return dara.Validate(s)
}
