// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagKeysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKeys(v []*string) *ListTagKeysResponseBody
	GetKeys() []*string
	SetMaxResults(v int32) *ListTagKeysResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListTagKeysResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTagKeysResponseBody
	GetRequestId() *string
}

type ListTagKeysResponseBody struct {
	// The tag keys.
	Keys []*string `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Repeated"`
	// The maximum number of entries to return on each page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page.
	//
	// example:
	//
	// 87y29h80h20h3f2
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 36210B73-8262-4D08-9D3A-7F96789733C8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListTagKeysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTagKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBody) GetKeys() []*string {
	return s.Keys
}

func (s *ListTagKeysResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTagKeysResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTagKeysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTagKeysResponseBody) SetKeys(v []*string) *ListTagKeysResponseBody {
	s.Keys = v
	return s
}

func (s *ListTagKeysResponseBody) SetMaxResults(v int32) *ListTagKeysResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTagKeysResponseBody) SetNextToken(v string) *ListTagKeysResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagKeysResponseBody) SetRequestId(v string) *ListTagKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagKeysResponseBody) Validate() error {
	return dara.Validate(s)
}
