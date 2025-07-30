// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableUserResponseBody
	GetRequestId() *string
}

type DisableUserResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableUserResponseBody) GoString() string {
	return s.String()
}

func (s *DisableUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableUserResponseBody) SetRequestId(v string) *DisableUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableUserResponseBody) Validate() error {
	return dara.Validate(s)
}
