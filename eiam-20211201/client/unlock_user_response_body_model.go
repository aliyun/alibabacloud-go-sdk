// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnlockUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnlockUserResponseBody
	GetRequestId() *string
}

type UnlockUserResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnlockUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnlockUserResponseBody) GoString() string {
	return s.String()
}

func (s *UnlockUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnlockUserResponseBody) SetRequestId(v string) *UnlockUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnlockUserResponseBody) Validate() error {
	return dara.Validate(s)
}
