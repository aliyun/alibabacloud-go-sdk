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
	SetValues(v *ListTagValuesResponseBodyValues) *ListTagValuesResponseBody
	GetValues() *ListTagValuesResponseBodyValues
}

type ListTagValuesResponseBody struct {
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
	// 8989CA7E-D2E0-4B6D-8282-311106E80150
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information of the tag values.
	Values *ListTagValuesResponseBodyValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Struct"`
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

func (s *ListTagValuesResponseBody) GetValues() *ListTagValuesResponseBodyValues {
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

func (s *ListTagValuesResponseBody) SetValues(v *ListTagValuesResponseBodyValues) *ListTagValuesResponseBody {
	s.Values = v
	return s
}

func (s *ListTagValuesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListTagValuesResponseBodyValues struct {
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListTagValuesResponseBodyValues) String() string {
	return dara.Prettify(s)
}

func (s ListTagValuesResponseBodyValues) GoString() string {
	return s.String()
}

func (s *ListTagValuesResponseBodyValues) GetValue() []*string {
	return s.Value
}

func (s *ListTagValuesResponseBodyValues) SetValue(v []*string) *ListTagValuesResponseBodyValues {
	s.Value = v
	return s
}

func (s *ListTagValuesResponseBodyValues) Validate() error {
	return dara.Validate(s)
}
