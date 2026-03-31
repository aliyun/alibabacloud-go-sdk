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
	SetRequestId(v string) *ListTagValuesResponseBody
	GetRequestId() *string
	SetValues(v []*string) *ListTagValuesResponseBody
	GetValues() []*string
}

type ListTagValuesResponseBody struct {
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0*****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 705114BB-EAEF-5CC4-8837-F1D4****BB5B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tag values.
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
