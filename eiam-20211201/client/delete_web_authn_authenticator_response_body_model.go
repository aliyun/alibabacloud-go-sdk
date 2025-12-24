// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWebAuthnAuthenticatorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteWebAuthnAuthenticatorResponseBody
	GetRequestId() *string
}

type DeleteWebAuthnAuthenticatorResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteWebAuthnAuthenticatorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWebAuthnAuthenticatorResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWebAuthnAuthenticatorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWebAuthnAuthenticatorResponseBody) SetRequestId(v string) *DeleteWebAuthnAuthenticatorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWebAuthnAuthenticatorResponseBody) Validate() error {
	return dara.Validate(s)
}
