// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOAuth2CredentialProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateOAuth2CredentialProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateOAuth2CredentialProviderResponse
	GetStatusCode() *int32
	SetBody(v *UpdateOAuth2CredentialProviderResponseBody) *UpdateOAuth2CredentialProviderResponse
	GetBody() *UpdateOAuth2CredentialProviderResponseBody
}

type UpdateOAuth2CredentialProviderResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateOAuth2CredentialProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateOAuth2CredentialProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateOAuth2CredentialProviderResponse) GoString() string {
	return s.String()
}

func (s *UpdateOAuth2CredentialProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateOAuth2CredentialProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateOAuth2CredentialProviderResponse) GetBody() *UpdateOAuth2CredentialProviderResponseBody {
	return s.Body
}

func (s *UpdateOAuth2CredentialProviderResponse) SetHeaders(v map[string]*string) *UpdateOAuth2CredentialProviderResponse {
	s.Headers = v
	return s
}

func (s *UpdateOAuth2CredentialProviderResponse) SetStatusCode(v int32) *UpdateOAuth2CredentialProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateOAuth2CredentialProviderResponse) SetBody(v *UpdateOAuth2CredentialProviderResponseBody) *UpdateOAuth2CredentialProviderResponse {
	s.Body = v
	return s
}

func (s *UpdateOAuth2CredentialProviderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
