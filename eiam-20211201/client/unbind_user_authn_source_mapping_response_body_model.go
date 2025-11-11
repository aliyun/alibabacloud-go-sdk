// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindUserAuthnSourceMappingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnbindUserAuthnSourceMappingResponseBody
	GetRequestId() *string
}

type UnbindUserAuthnSourceMappingResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindUserAuthnSourceMappingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnbindUserAuthnSourceMappingResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindUserAuthnSourceMappingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnbindUserAuthnSourceMappingResponseBody) SetRequestId(v string) *UnbindUserAuthnSourceMappingResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindUserAuthnSourceMappingResponseBody) Validate() error {
	return dara.Validate(s)
}
