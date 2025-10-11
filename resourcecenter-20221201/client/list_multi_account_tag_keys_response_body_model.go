// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultiAccountTagKeysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListMultiAccountTagKeysResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListMultiAccountTagKeysResponseBody
	GetRequestId() *string
	SetTagKeys(v []*string) *ListMultiAccountTagKeysResponseBody
	GetTagKeys() []*string
}

type ListMultiAccountTagKeysResponseBody struct {
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAARfZmVDe9NvRXloR5+8CK9nNJufMdRA7W1miLC1P****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// FA6086F9-6363-51A5-A507-88E3201EBCCB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tag keys.
	TagKeys []*string `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Repeated"`
}

func (s ListMultiAccountTagKeysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMultiAccountTagKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListMultiAccountTagKeysResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMultiAccountTagKeysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMultiAccountTagKeysResponseBody) GetTagKeys() []*string {
	return s.TagKeys
}

func (s *ListMultiAccountTagKeysResponseBody) SetNextToken(v string) *ListMultiAccountTagKeysResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMultiAccountTagKeysResponseBody) SetRequestId(v string) *ListMultiAccountTagKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMultiAccountTagKeysResponseBody) SetTagKeys(v []*string) *ListMultiAccountTagKeysResponseBody {
	s.TagKeys = v
	return s
}

func (s *ListMultiAccountTagKeysResponseBody) Validate() error {
	return dara.Validate(s)
}
