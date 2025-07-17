// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFunctionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateFunctionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateFunctionResponseBody
	GetSuccess() *bool
}

type UpdateFunctionResponseBody struct {
	// The request ID. You can troubleshoot issues based on the ID.
	//
	// example:
	//
	// 12123960-CB2C-5086-868E-C6C1D024XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// true
	//
	// false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateFunctionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFunctionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateFunctionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateFunctionResponseBody) SetRequestId(v string) *UpdateFunctionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFunctionResponseBody) SetSuccess(v bool) *UpdateFunctionResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateFunctionResponseBody) Validate() error {
	return dara.Validate(s)
}
