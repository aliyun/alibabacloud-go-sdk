// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagValuesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListTagValuesResponseBody
	GetNextToken() *string
	SetPageSize(v int32) *ListTagValuesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListTagValuesResponseBody
	GetRequestId() *string
	SetValues(v []*string) *ListTagValuesResponseBody
	GetValues() []*string
}

type ListTagValuesResponseBody struct {
	// The token that determines the start point of the next query. If this parameter is empty, all results are returned.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// AB444F46-1CFF-4B06-B8F0-B45D3158****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Details of the tag values.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListTagValuesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTagValuesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagValuesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTagValuesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTagValuesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTagValuesResponseBody) GetValues() []*string {
	return s.Values
}

func (s *ListTagValuesResponseBody) SetNextToken(v string) *ListTagValuesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagValuesResponseBody) SetPageSize(v int32) *ListTagValuesResponseBody {
	s.PageSize = &v
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
