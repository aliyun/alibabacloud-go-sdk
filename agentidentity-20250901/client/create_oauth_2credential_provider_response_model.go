// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOAuth2CredentialProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOAuth2CredentialProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOAuth2CredentialProviderResponse
	GetStatusCode() *int32
	SetBody(v *CreateOAuth2CredentialProviderResponseBody) *CreateOAuth2CredentialProviderResponse
	GetBody() *CreateOAuth2CredentialProviderResponseBody
}

type CreateOAuth2CredentialProviderResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOAuth2CredentialProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOAuth2CredentialProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOAuth2CredentialProviderResponse) GoString() string {
	return s.String()
}

func (s *CreateOAuth2CredentialProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOAuth2CredentialProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOAuth2CredentialProviderResponse) GetBody() *CreateOAuth2CredentialProviderResponseBody {
	return s.Body
}

func (s *CreateOAuth2CredentialProviderResponse) SetHeaders(v map[string]*string) *CreateOAuth2CredentialProviderResponse {
	s.Headers = v
	return s
}

func (s *CreateOAuth2CredentialProviderResponse) SetStatusCode(v int32) *CreateOAuth2CredentialProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOAuth2CredentialProviderResponse) SetBody(v *CreateOAuth2CredentialProviderResponseBody) *CreateOAuth2CredentialProviderResponse {
	s.Body = v
	return s
}

func (s *CreateOAuth2CredentialProviderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
