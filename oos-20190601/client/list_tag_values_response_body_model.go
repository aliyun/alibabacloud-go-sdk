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
	SetValues(v []*string) *ListTagValuesResponseBody
	GetValues() []*string
}

type ListTagValuesResponseBody struct {
	// The maximum number of results on each page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The returned value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// 83u29j2dj3dskds
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 65591133-1188-4935-B78F-20F72
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tag values returned.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
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

func (s *ListTagValuesResponseBody) GetValues() []*string {
	return s.Values
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

func (s *ListTagValuesResponseBody) SetValues(v []*string) *ListTagValuesResponseBody {
	s.Values = v
	return s
}

func (s *ListTagValuesResponseBody) Validate() error {
	return dara.Validate(s)
}
