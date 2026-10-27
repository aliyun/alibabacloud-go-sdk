// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIdentityProviderSCIMCredentialResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateIdentityProviderSCIMCredentialResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateIdentityProviderSCIMCredentialResponse
	GetStatusCode() *int32
	SetBody(v *CreateIdentityProviderSCIMCredentialResponseBody) *CreateIdentityProviderSCIMCredentialResponse
	GetBody() *CreateIdentityProviderSCIMCredentialResponseBody
}

type CreateIdentityProviderSCIMCredentialResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateIdentityProviderSCIMCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIdentityProviderSCIMCredentialResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentityProviderSCIMCredentialResponse) GoString() string {
	return s.String()
}

func (s *CreateIdentityProviderSCIMCredentialResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateIdentityProviderSCIMCredentialResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateIdentityProviderSCIMCredentialResponse) GetBody() *CreateIdentityProviderSCIMCredentialResponseBody {
	return s.Body
}

func (s *CreateIdentityProviderSCIMCredentialResponse) SetHeaders(v map[string]*string) *CreateIdentityProviderSCIMCredentialResponse {
	s.Headers = v
	return s
}

func (s *CreateIdentityProviderSCIMCredentialResponse) SetStatusCode(v int32) *CreateIdentityProviderSCIMCredentialResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIdentityProviderSCIMCredentialResponse) SetBody(v *CreateIdentityProviderSCIMCredentialResponseBody) *CreateIdentityProviderSCIMCredentialResponse {
	s.Body = v
	return s
}

func (s *CreateIdentityProviderSCIMCredentialResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
