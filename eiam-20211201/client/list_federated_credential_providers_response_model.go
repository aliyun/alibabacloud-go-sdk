// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFederatedCredentialProvidersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFederatedCredentialProvidersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFederatedCredentialProvidersResponse
	GetStatusCode() *int32
	SetBody(v *ListFederatedCredentialProvidersResponseBody) *ListFederatedCredentialProvidersResponse
	GetBody() *ListFederatedCredentialProvidersResponseBody
}

type ListFederatedCredentialProvidersResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFederatedCredentialProvidersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFederatedCredentialProvidersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFederatedCredentialProvidersResponse) GoString() string {
	return s.String()
}

func (s *ListFederatedCredentialProvidersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFederatedCredentialProvidersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFederatedCredentialProvidersResponse) GetBody() *ListFederatedCredentialProvidersResponseBody {
	return s.Body
}

func (s *ListFederatedCredentialProvidersResponse) SetHeaders(v map[string]*string) *ListFederatedCredentialProvidersResponse {
	s.Headers = v
	return s
}

func (s *ListFederatedCredentialProvidersResponse) SetStatusCode(v int32) *ListFederatedCredentialProvidersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFederatedCredentialProvidersResponse) SetBody(v *ListFederatedCredentialProvidersResponseBody) *ListFederatedCredentialProvidersResponse {
	s.Body = v
	return s
}

func (s *ListFederatedCredentialProvidersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
