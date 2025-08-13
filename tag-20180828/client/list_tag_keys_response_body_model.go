// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagKeysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKeys(v *ListTagKeysResponseBodyKeys) *ListTagKeysResponseBody
	GetKeys() *ListTagKeysResponseBodyKeys
	SetNextToken(v string) *ListTagKeysResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTagKeysResponseBody
	GetRequestId() *string
}

type ListTagKeysResponseBody struct {
	// The information of the tag keys.
	Keys *ListTagKeysResponseBodyKeys `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Struct"`
	// Indicates whether the next query is required. The value of this parameter may be empty.
	//
	// 	- If the value of this parameter is empty (`"NextToken": ""`), all results are returned, and the next query is not required.
	//
	// 	- If the value of this parameter is not empty, the next query is required, and the value is the token used to start the next query.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// DC09A6AA-2713-4E10-A2E9-E6C5C43A8842
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListTagKeysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTagKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBody) GetKeys() *ListTagKeysResponseBodyKeys {
	return s.Keys
}

func (s *ListTagKeysResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTagKeysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTagKeysResponseBody) SetKeys(v *ListTagKeysResponseBodyKeys) *ListTagKeysResponseBody {
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
	Key []*ListTagKeysResponseBodyKeysKey `json:"Key,omitempty" xml:"Key,omitempty" type:"Repeated"`
}

func (s ListTagKeysResponseBodyKeys) String() string {
	return dara.Prettify(s)
}

func (s ListTagKeysResponseBodyKeys) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBodyKeys) GetKey() []*ListTagKeysResponseBodyKeysKey {
	return s.Key
}

func (s *ListTagKeysResponseBodyKeys) SetKey(v []*ListTagKeysResponseBodyKeysKey) *ListTagKeysResponseBodyKeys {
	s.Key = v
	return s
}

func (s *ListTagKeysResponseBodyKeys) Validate() error {
	return dara.Validate(s)
}

type ListTagKeysResponseBodyKeysKey struct {
	// The type of the resource tag. Valid values:
	//
	// 	- custom
	//
	// 	- system
	//
	// example:
	//
	// custom
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The description of the tag key.
	//
	// example:
	//
	// Business team
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The tag key.
	//
	// example:
	//
	// team
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s ListTagKeysResponseBodyKeysKey) String() string {
	return dara.Prettify(s)
}

func (s ListTagKeysResponseBodyKeysKey) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBodyKeysKey) GetCategory() *string {
	return s.Category
}

func (s *ListTagKeysResponseBodyKeysKey) GetDescription() *string {
	return s.Description
}

func (s *ListTagKeysResponseBodyKeysKey) GetKey() *string {
	return s.Key
}

func (s *ListTagKeysResponseBodyKeysKey) SetCategory(v string) *ListTagKeysResponseBodyKeysKey {
	s.Category = &v
	return s
}

func (s *ListTagKeysResponseBodyKeysKey) SetDescription(v string) *ListTagKeysResponseBodyKeysKey {
	s.Description = &v
	return s
}

func (s *ListTagKeysResponseBodyKeysKey) SetKey(v string) *ListTagKeysResponseBodyKeysKey {
	s.Key = &v
	return s
}

func (s *ListTagKeysResponseBodyKeysKey) Validate() error {
	return dara.Validate(s)
}
