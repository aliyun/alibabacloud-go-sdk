// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOAuth2CredentialProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteOAuth2CredentialProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteOAuth2CredentialProviderResponse
	GetStatusCode() *int32
	SetBody(v *DeleteOAuth2CredentialProviderResponseBody) *DeleteOAuth2CredentialProviderResponse
	GetBody() *DeleteOAuth2CredentialProviderResponseBody
}

type DeleteOAuth2CredentialProviderResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteOAuth2CredentialProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteOAuth2CredentialProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteOAuth2CredentialProviderResponse) GoString() string {
	return s.String()
}

func (s *DeleteOAuth2CredentialProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteOAuth2CredentialProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteOAuth2CredentialProviderResponse) GetBody() *DeleteOAuth2CredentialProviderResponseBody {
	return s.Body
}

func (s *DeleteOAuth2CredentialProviderResponse) SetHeaders(v map[string]*string) *DeleteOAuth2CredentialProviderResponse {
	s.Headers = v
	return s
}

func (s *DeleteOAuth2CredentialProviderResponse) SetStatusCode(v int32) *DeleteOAuth2CredentialProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteOAuth2CredentialProviderResponse) SetBody(v *DeleteOAuth2CredentialProviderResponseBody) *DeleteOAuth2CredentialProviderResponse {
	s.Body = v
	return s
}

func (s *DeleteOAuth2CredentialProviderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
