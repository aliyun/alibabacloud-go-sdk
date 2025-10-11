// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagKeysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListTagKeysResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTagKeysResponseBody
	GetRequestId() *string
	SetTagKeys(v []*string) *ListTagKeysResponseBody
	GetTagKeys() []*string
}

type ListTagKeysResponseBody struct {
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAUDnubHKJbVTCdlIGYUPtsu3EoN3bfdgjDA****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 44C8A952-D6B0-5BC8-82D5-93BA02E26F2E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tag keys.
	TagKeys []*string `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Repeated"`
}

func (s ListTagKeysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTagKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTagKeysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTagKeysResponseBody) GetTagKeys() []*string {
	return s.TagKeys
}

func (s *ListTagKeysResponseBody) SetNextToken(v string) *ListTagKeysResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagKeysResponseBody) SetRequestId(v string) *ListTagKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagKeysResponseBody) SetTagKeys(v []*string) *ListTagKeysResponseBody {
	s.TagKeys = v
	return s
}

func (s *ListTagKeysResponseBody) Validate() error {
	return dara.Validate(s)
}
