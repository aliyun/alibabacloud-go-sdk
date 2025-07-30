// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserPasswordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateUserPasswordResponseBody
	GetRequestId() *string
}

type UpdateUserPasswordResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateUserPasswordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserPasswordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateUserPasswordResponseBody) SetRequestId(v string) *UpdateUserPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserPasswordResponseBody) Validate() error {
	return dara.Validate(s)
}
