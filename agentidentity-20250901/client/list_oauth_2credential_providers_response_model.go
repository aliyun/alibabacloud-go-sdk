// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOAuth2CredentialProvidersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOAuth2CredentialProvidersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOAuth2CredentialProvidersResponse
	GetStatusCode() *int32
	SetBody(v *ListOAuth2CredentialProvidersResponseBody) *ListOAuth2CredentialProvidersResponse
	GetBody() *ListOAuth2CredentialProvidersResponseBody
}

type ListOAuth2CredentialProvidersResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOAuth2CredentialProvidersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOAuth2CredentialProvidersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOAuth2CredentialProvidersResponse) GoString() string {
	return s.String()
}

func (s *ListOAuth2CredentialProvidersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOAuth2CredentialProvidersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOAuth2CredentialProvidersResponse) GetBody() *ListOAuth2CredentialProvidersResponseBody {
	return s.Body
}

func (s *ListOAuth2CredentialProvidersResponse) SetHeaders(v map[string]*string) *ListOAuth2CredentialProvidersResponse {
	s.Headers = v
	return s
}

func (s *ListOAuth2CredentialProvidersResponse) SetStatusCode(v int32) *ListOAuth2CredentialProvidersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOAuth2CredentialProvidersResponse) SetBody(v *ListOAuth2CredentialProvidersResponseBody) *ListOAuth2CredentialProvidersResponse {
	s.Body = v
	return s
}

func (s *ListOAuth2CredentialProvidersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
