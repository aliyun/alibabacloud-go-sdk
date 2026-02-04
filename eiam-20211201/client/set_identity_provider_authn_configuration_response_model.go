// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetIdentityProviderAuthnConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetIdentityProviderAuthnConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetIdentityProviderAuthnConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *SetIdentityProviderAuthnConfigurationResponseBody) *SetIdentityProviderAuthnConfigurationResponse
	GetBody() *SetIdentityProviderAuthnConfigurationResponseBody
}

type SetIdentityProviderAuthnConfigurationResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetIdentityProviderAuthnConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetIdentityProviderAuthnConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s SetIdentityProviderAuthnConfigurationResponse) GoString() string {
	return s.String()
}

func (s *SetIdentityProviderAuthnConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetIdentityProviderAuthnConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetIdentityProviderAuthnConfigurationResponse) GetBody() *SetIdentityProviderAuthnConfigurationResponseBody {
	return s.Body
}

func (s *SetIdentityProviderAuthnConfigurationResponse) SetHeaders(v map[string]*string) *SetIdentityProviderAuthnConfigurationResponse {
	s.Headers = v
	return s
}

func (s *SetIdentityProviderAuthnConfigurationResponse) SetStatusCode(v int32) *SetIdentityProviderAuthnConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *SetIdentityProviderAuthnConfigurationResponse) SetBody(v *SetIdentityProviderAuthnConfigurationResponseBody) *SetIdentityProviderAuthnConfigurationResponse {
	s.Body = v
	return s
}

func (s *SetIdentityProviderAuthnConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
