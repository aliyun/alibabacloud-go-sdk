// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetIdentityProviderAuthnConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetIdentityProviderAuthnConfigurationResponseBody
	GetRequestId() *string
}

type SetIdentityProviderAuthnConfigurationResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetIdentityProviderAuthnConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetIdentityProviderAuthnConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *SetIdentityProviderAuthnConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetIdentityProviderAuthnConfigurationResponseBody) SetRequestId(v string) *SetIdentityProviderAuthnConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetIdentityProviderAuthnConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}
