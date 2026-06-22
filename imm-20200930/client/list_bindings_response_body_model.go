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
	// The list of binding information between datasets and OSS buckets.
	Bindings []*Binding `json:"Bindings,omitempty" xml:"Bindings,omitempty" type:"Repeated"`
	// - The pagination token that is used when the total number of bindings exceeds the MaxResults value.
	//
	// - Use this value as the NextToken in the next request to return the remaining results.
	//
	// - This parameter has a value only when not all bindings are returned.
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
	if s.Bindings != nil {
		for _, item := range s.Bindings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
