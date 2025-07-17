// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateNodeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateNodeResponseBody
	GetSuccess() *bool
}

type UpdateNodeResponseBody struct {
	// The request ID. You can troubleshoot issues based on the ID.
	//
	// example:
	//
	// 99EBE7CF-69C0-5089-BE3E-79563C31XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateNodeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateNodeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateNodeResponseBody) SetRequestId(v string) *UpdateNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateNodeResponseBody) SetSuccess(v bool) *UpdateNodeResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateNodeResponseBody) Validate() error {
	return dara.Validate(s)
}
