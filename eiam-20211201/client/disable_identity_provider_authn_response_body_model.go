// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableIdentityProviderAuthnResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableIdentityProviderAuthnResponseBody
	GetRequestId() *string
}

type DisableIdentityProviderAuthnResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableIdentityProviderAuthnResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableIdentityProviderAuthnResponseBody) GoString() string {
	return s.String()
}

func (s *DisableIdentityProviderAuthnResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableIdentityProviderAuthnResponseBody) SetRequestId(v string) *DisableIdentityProviderAuthnResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableIdentityProviderAuthnResponseBody) Validate() error {
	return dara.Validate(s)
}
