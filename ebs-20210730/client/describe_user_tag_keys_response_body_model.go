// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserTagKeysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeUserTagKeysResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeUserTagKeysResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeUserTagKeysResponseBody
	GetRequestId() *string
	SetTagKeys(v []*string) *DescribeUserTagKeysResponseBody
	GetTagKeys() []*string
}

type DescribeUserTagKeysResponseBody struct {
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
	// The token for the next query. An empty NextToken indicates there are no more results.
	//
	// example:
	//
	// f07b150eadfa1d7a
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Request ID.
	//
	// example:
	//
	// C123F94F-4E38-19AE-942A-A8D6F44F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// List of matching tag keys.
	TagKeys []*string `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Repeated"`
}

func (s DescribeUserTagKeysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserTagKeysResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserTagKeysResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeUserTagKeysResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeUserTagKeysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserTagKeysResponseBody) GetTagKeys() []*string {
	return s.TagKeys
}

func (s *DescribeUserTagKeysResponseBody) SetMaxResults(v int32) *DescribeUserTagKeysResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeUserTagKeysResponseBody) SetNextToken(v string) *DescribeUserTagKeysResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeUserTagKeysResponseBody) SetRequestId(v string) *DescribeUserTagKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserTagKeysResponseBody) SetTagKeys(v []*string) *DescribeUserTagKeysResponseBody {
	s.TagKeys = v
	return s
}

func (s *DescribeUserTagKeysResponseBody) Validate() error {
	return dara.Validate(s)
}
