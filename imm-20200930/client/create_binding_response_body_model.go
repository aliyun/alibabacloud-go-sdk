// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBindingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBinding(v *Binding) *CreateBindingResponseBody
	GetBinding() *Binding
	SetRequestId(v string) *CreateBindingResponseBody
	GetRequestId() *string
}

type CreateBindingResponseBody struct {
	// The binding relationship.
	Binding *Binding `json:"Binding,omitempty" xml:"Binding,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5F74C5C9-5AC0-49F9-914D-E01589D3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateBindingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBindingResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBindingResponseBody) GetBinding() *Binding {
	return s.Binding
}

func (s *CreateBindingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBindingResponseBody) SetBinding(v *Binding) *CreateBindingResponseBody {
	s.Binding = v
	return s
}

func (s *CreateBindingResponseBody) SetRequestId(v string) *CreateBindingResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBindingResponseBody) Validate() error {
	return dara.Validate(s)
}
