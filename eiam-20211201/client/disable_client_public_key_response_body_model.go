// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableClientPublicKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableClientPublicKeyResponseBody
	GetRequestId() *string
}

type DisableClientPublicKeyResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableClientPublicKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableClientPublicKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DisableClientPublicKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableClientPublicKeyResponseBody) SetRequestId(v string) *DisableClientPublicKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableClientPublicKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
