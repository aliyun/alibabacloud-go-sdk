// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagValuesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListTagValuesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListTagValuesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTagValuesResponseBody
	GetRequestId() *string
	SetTagValues(v []*string) *ListTagValuesResponseBody
	GetTagValues() []*string
	SetTotalCount(v int32) *ListTagValuesResponseBody
	GetTotalCount() *int32
}

type ListTagValuesResponseBody struct {
	// The number of entries per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value is returned for **NextToken**, the value is the token that determines the start point of the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 593B0448-D13E-4C56-AC0D-FDF0FDE0E9A3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tag values.
	TagValues []*string `json:"TagValues,omitempty" xml:"TagValues,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTagValuesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTagValuesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagValuesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTagValuesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTagValuesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTagValuesResponseBody) GetTagValues() []*string {
	return s.TagValues
}

func (s *ListTagValuesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTagValuesResponseBody) SetMaxResults(v int32) *ListTagValuesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTagValuesResponseBody) SetNextToken(v string) *ListTagValuesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagValuesResponseBody) SetRequestId(v string) *ListTagValuesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagValuesResponseBody) SetTagValues(v []*string) *ListTagValuesResponseBody {
	s.TagValues = v
	return s
}

func (s *ListTagValuesResponseBody) SetTotalCount(v int32) *ListTagValuesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTagValuesResponseBody) Validate() error {
	return dara.Validate(s)
}
