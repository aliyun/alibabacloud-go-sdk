// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateWebAuthnAuthenticatorRegistrationUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GenerateWebAuthnAuthenticatorRegistrationUrlResponseBody
	GetRequestId() *string
	SetWebAuthnAuthenticatorRegistrationUrl(v *GenerateWebAuthnAuthenticatorRegistrationUrlResponseBodyWebAuthnAuthenticatorRegistrationUrl) *GenerateWebAuthnAuthenticatorRegistrationUrlResponseBody
	GetWebAuthnAuthenticatorRegistrationUrl() *GenerateWebAuthnAuthenticatorRegistrationUrlResponseBodyWebAuthnAuthenticatorRegistrationUrl
}

type GenerateWebAuthnAuthenticatorRegistrationUrlResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId                            *string                                                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	WebAuthnAuthenticatorRegistrationUrl *GenerateWebAuthnAuthenticatorRegistrationUrlResponseBodyWebAuthnAuthenticatorRegistrationUrl `json:"WebAuthnAuthenticatorRegistrationUrl,omitempty" xml:"WebAuthnAuthenticatorRegistrationUrl,omitempty" type:"Struct"`
}

func (s GenerateWebAuthnAuthenticatorRegistrationUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateWebAuthnAuthenticatorRegistrationUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateWebAuthnAuthenticatorRegistrationUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateWebAuthnAuthenticatorRegistrationUrlResponseBody) GetWebAuthnAuthenticatorRegistrationUrl() *GenerateWebAuthnAuthenticatorRegistrationUrlResponseBodyWebAuthnAuthenticatorRegistrationUrl {
	return s.WebAuthnAuthenticatorRegistrationUrl
}

func (s *GenerateWebAuthnAuthenticatorRegistrationUrlResponseBody) SetRequestId(v string) *GenerateWebAuthnAuthenticatorRegistrationUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateWebAuthnAuthenticatorRegistrationUrlResponseBody) SetWebAuthnAuthenticatorRegistrationUrl(v *GenerateWebAuthnAuthenticatorRegistrationUrlResponseBodyWebAuthnAuthenticatorRegistrationUrl) *GenerateWebAuthnAuthenticatorRegistrationUrlResponseBody {
	s.WebAuthnAuthenticatorRegistrationUrl = v
	return s
}

func (s *GenerateWebAuthnAuthenticatorRegistrationUrlResponseBody) Validate() error {
	if s.WebAuthnAuthenticatorRegistrationUrl != nil {
		if err := s.WebAuthnAuthenticatorRegistrationUrl.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GenerateWebAuthnAuthenticatorRegistrationUrlResponseBodyWebAuthnAuthenticatorRegistrationUrl struct {
	// 注册WebAuthn认证器URL
	//
	// example:
	//
	// https://012cnaliyunidaas..com./login/webauthn/registration
	RegistrationUrl *string `json:"RegistrationUrl,omitempty" xml:"RegistrationUrl,omitempty"`
	// 注册WebAuthn认证器URL参数
	//
	// example:
	//
	// eyJhbGciOiJkaXIiLCJlbmMiOiJBMjU2R0NNIiwia2lkIjoia2V5X29ueWNzbXNib3Y1bmV2anlncHgyZnlsbjdhIn0
	RegistrationUrlParameters *string `json:"RegistrationUrlParameters,omitempty" xml:"RegistrationUrlParameters,omitempty"`
}

func (s GenerateWebAuthnAuthenticatorRegistrationUrlResponseBodyWebAuthnAuthenticatorRegistrationUrl) String() string {
	return dara.Prettify(s)
}

func (s GenerateWebAuthnAuthenticatorRegistrationUrlResponseBodyWebAuthnAuthenticatorRegistrationUrl) GoString() string {
	return s.String()
}

func (s *GenerateWebAuthnAuthenticatorRegistrationUrlResponseBodyWebAuthnAuthenticatorRegistrationUrl) GetRegistrationUrl() *string {
	return s.RegistrationUrl
}

func (s *GenerateWebAuthnAuthenticatorRegistrationUrlResponseBodyWebAuthnAuthenticatorRegistrationUrl) GetRegistrationUrlParameters() *string {
	return s.RegistrationUrlParameters
}

func (s *GenerateWebAuthnAuthenticatorRegistrationUrlResponseBodyWebAuthnAuthenticatorRegistrationUrl) SetRegistrationUrl(v string) *GenerateWebAuthnAuthenticatorRegistrationUrlResponseBodyWebAuthnAuthenticatorRegistrationUrl {
	s.RegistrationUrl = &v
	return s
}

func (s *GenerateWebAuthnAuthenticatorRegistrationUrlResponseBodyWebAuthnAuthenticatorRegistrationUrl) SetRegistrationUrlParameters(v string) *GenerateWebAuthnAuthenticatorRegistrationUrlResponseBodyWebAuthnAuthenticatorRegistrationUrl {
	s.RegistrationUrlParameters = &v
	return s
}

func (s *GenerateWebAuthnAuthenticatorRegistrationUrlResponseBodyWebAuthnAuthenticatorRegistrationUrl) Validate() error {
	return dara.Validate(s)
}
