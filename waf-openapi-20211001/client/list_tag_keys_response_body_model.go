// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagKeysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKeys(v []*ListTagKeysResponseBodyKeys) *ListTagKeysResponseBody
	GetKeys() []*ListTagKeysResponseBodyKeys
	SetNextToken(v string) *ListTagKeysResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTagKeysResponseBody
	GetRequestId() *string
}

type ListTagKeysResponseBody struct {
	// The keys and types of the tags.
	Keys []*ListTagKeysResponseBodyKeys `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Repeated"`
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
	// 8CB8BB88-24C7-5608-BF5E-4DCA****CF1C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListTagKeysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTagKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBody) GetKeys() []*ListTagKeysResponseBodyKeys {
	return s.Keys
}

func (s *ListTagKeysResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTagKeysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTagKeysResponseBody) SetKeys(v []*ListTagKeysResponseBodyKeys) *ListTagKeysResponseBody {
	s.Keys = v
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

type ListTagKeysResponseBodyKeys struct {
	// The type of the tag. Valid values:
	//
	// 	- custom
	//
	// 	- system
	//
	// example:
	//
	// custom
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The key of the tag.
	//
	// example:
	//
	// demoTagKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s ListTagKeysResponseBodyKeys) String() string {
	return dara.Prettify(s)
}

func (s ListTagKeysResponseBodyKeys) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBodyKeys) GetCategory() *string {
	return s.Category
}

func (s *ListTagKeysResponseBodyKeys) GetKey() *string {
	return s.Key
}

func (s *ListTagKeysResponseBodyKeys) SetCategory(v string) *ListTagKeysResponseBodyKeys {
	s.Category = &v
	return s
}

func (s *ListTagKeysResponseBodyKeys) SetKey(v string) *ListTagKeysResponseBodyKeys {
	s.Key = &v
	return s
}

func (s *ListTagKeysResponseBodyKeys) Validate() error {
	return dara.Validate(s)
}
