// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateWebAuthnAuthenticatorRegistrationUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateWebAuthnAuthenticatorRegistrationUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateWebAuthnAuthenticatorRegistrationUrlResponse
	GetStatusCode() *int32
	SetBody(v *GenerateWebAuthnAuthenticatorRegistrationUrlResponseBody) *GenerateWebAuthnAuthenticatorRegistrationUrlResponse
	GetBody() *GenerateWebAuthnAuthenticatorRegistrationUrlResponseBody
}

type GenerateWebAuthnAuthenticatorRegistrationUrlResponse struct {
	Headers    map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateWebAuthnAuthenticatorRegistrationUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateWebAuthnAuthenticatorRegistrationUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateWebAuthnAuthenticatorRegistrationUrlResponse) GoString() string {
	return s.String()
}

func (s *GenerateWebAuthnAuthenticatorRegistrationUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateWebAuthnAuthenticatorRegistrationUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateWebAuthnAuthenticatorRegistrationUrlResponse) GetBody() *GenerateWebAuthnAuthenticatorRegistrationUrlResponseBody {
	return s.Body
}

func (s *GenerateWebAuthnAuthenticatorRegistrationUrlResponse) SetHeaders(v map[string]*string) *GenerateWebAuthnAuthenticatorRegistrationUrlResponse {
	s.Headers = v
	return s
}

func (s *GenerateWebAuthnAuthenticatorRegistrationUrlResponse) SetStatusCode(v int32) *GenerateWebAuthnAuthenticatorRegistrationUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateWebAuthnAuthenticatorRegistrationUrlResponse) SetBody(v *GenerateWebAuthnAuthenticatorRegistrationUrlResponseBody) *GenerateWebAuthnAuthenticatorRegistrationUrlResponse {
	s.Body = v
	return s
}

func (s *GenerateWebAuthnAuthenticatorRegistrationUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
