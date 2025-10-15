// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationFederatedCredentialsForProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApplicationFederatedCredentialsForProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApplicationFederatedCredentialsForProviderResponse
	GetStatusCode() *int32
	SetBody(v *ListApplicationFederatedCredentialsForProviderResponseBody) *ListApplicationFederatedCredentialsForProviderResponse
	GetBody() *ListApplicationFederatedCredentialsForProviderResponseBody
}

type ListApplicationFederatedCredentialsForProviderResponse struct {
	Headers    map[string]*string                                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApplicationFederatedCredentialsForProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApplicationFederatedCredentialsForProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationFederatedCredentialsForProviderResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationFederatedCredentialsForProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApplicationFederatedCredentialsForProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApplicationFederatedCredentialsForProviderResponse) GetBody() *ListApplicationFederatedCredentialsForProviderResponseBody {
	return s.Body
}

func (s *ListApplicationFederatedCredentialsForProviderResponse) SetHeaders(v map[string]*string) *ListApplicationFederatedCredentialsForProviderResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponse) SetStatusCode(v int32) *ListApplicationFederatedCredentialsForProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponse) SetBody(v *ListApplicationFederatedCredentialsForProviderResponseBody) *ListApplicationFederatedCredentialsForProviderResponse {
	s.Body = v
	return s
}

func (s *ListApplicationFederatedCredentialsForProviderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
