// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBindingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBindings(v []*Binding) *ListBindingsResponseBody
	GetBindings() []*Binding
	SetNextToken(v string) *ListBindingsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListBindingsResponseBody
	GetRequestId() *string
}

type ListBindingsResponseBody struct {
	// The bindings between the dataset and OSS buckets.
	Bindings []*Binding `json:"Bindings,omitempty" xml:"Bindings,omitempty" type:"Repeated"`
	// 	- The pagination token that is used in the next request to retrieve a new page of results if the total number of results exceeds the value of the MaxResults parameter.
	//
	// 	- The next request returns remaining results starting from the position marked by the NextToken parameter value.
	//
	// 	- This parameter has a non-empty value only when not all bindings are returned.
	//
	// example:
	//
	// immtest:dataset001:examplebucket01
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EFDFD356-C928-4A36-951A-6EB5A592****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListBindingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBindingsResponseBody) GoString() string {
	return s.String()
}

func (s *ListBindingsResponseBody) GetBindings() []*Binding {
	return s.Bindings
}

func (s *ListBindingsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListBindingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBindingsResponseBody) SetBindings(v []*Binding) *ListBindingsResponseBody {
	s.Bindings = v
	return s
}

func (s *ListBindingsResponseBody) SetNextToken(v string) *ListBindingsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListBindingsResponseBody) SetRequestId(v string) *ListBindingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBindingsResponseBody) Validate() error {
	return dara.Validate(s)
}
