// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBindingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBinding(v *Binding) *GetBindingResponseBody
	GetBinding() *Binding
	SetRequestId(v string) *GetBindingResponseBody
	GetRequestId() *string
}

type GetBindingResponseBody struct {
	// The details of the binding.
	Binding *Binding `json:"Binding,omitempty" xml:"Binding,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AEFCD467-C928-4A36-951A-6EB5A592****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetBindingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBindingResponseBody) GoString() string {
	return s.String()
}

func (s *GetBindingResponseBody) GetBinding() *Binding {
	return s.Binding
}

func (s *GetBindingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBindingResponseBody) SetBinding(v *Binding) *GetBindingResponseBody {
	s.Binding = v
	return s
}

func (s *GetBindingResponseBody) SetRequestId(v string) *GetBindingResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBindingResponseBody) Validate() error {
	return dara.Validate(s)
}
