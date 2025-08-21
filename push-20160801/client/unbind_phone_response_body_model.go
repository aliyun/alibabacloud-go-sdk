// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindPhoneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnbindPhoneResponseBody
	GetRequestId() *string
}

type UnbindPhoneResponseBody struct {
	// example:
	//
	// 0D1126F0-F8FF-513D-BAFA-F140447BDED4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindPhoneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnbindPhoneResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindPhoneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnbindPhoneResponseBody) SetRequestId(v string) *UnbindPhoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindPhoneResponseBody) Validate() error {
	return dara.Validate(s)
}
