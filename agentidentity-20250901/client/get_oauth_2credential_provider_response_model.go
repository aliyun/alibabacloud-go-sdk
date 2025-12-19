// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOAuth2CredentialProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOAuth2CredentialProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOAuth2CredentialProviderResponse
	GetStatusCode() *int32
	SetBody(v *GetOAuth2CredentialProviderResponseBody) *GetOAuth2CredentialProviderResponse
	GetBody() *GetOAuth2CredentialProviderResponseBody
}

type GetOAuth2CredentialProviderResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOAuth2CredentialProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOAuth2CredentialProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOAuth2CredentialProviderResponse) GoString() string {
	return s.String()
}

func (s *GetOAuth2CredentialProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOAuth2CredentialProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOAuth2CredentialProviderResponse) GetBody() *GetOAuth2CredentialProviderResponseBody {
	return s.Body
}

func (s *GetOAuth2CredentialProviderResponse) SetHeaders(v map[string]*string) *GetOAuth2CredentialProviderResponse {
	s.Headers = v
	return s
}

func (s *GetOAuth2CredentialProviderResponse) SetStatusCode(v int32) *GetOAuth2CredentialProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOAuth2CredentialProviderResponse) SetBody(v *GetOAuth2CredentialProviderResponseBody) *GetOAuth2CredentialProviderResponse {
	s.Body = v
	return s
}

func (s *GetOAuth2CredentialProviderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
