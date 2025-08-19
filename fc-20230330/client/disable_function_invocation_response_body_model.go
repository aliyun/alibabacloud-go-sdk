// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableFunctionInvocationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetSuccess(v bool) *DisableFunctionInvocationResponseBody
	GetSuccess() *bool
}

type DisableFunctionInvocationResponseBody struct {
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DisableFunctionInvocationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableFunctionInvocationResponseBody) GoString() string {
	return s.String()
}

func (s *DisableFunctionInvocationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DisableFunctionInvocationResponseBody) SetSuccess(v bool) *DisableFunctionInvocationResponseBody {
	s.Success = &v
	return s
}

func (s *DisableFunctionInvocationResponseBody) Validate() error {
	return dara.Validate(s)
}
