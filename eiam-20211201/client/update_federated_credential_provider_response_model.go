// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFederatedCredentialProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateFederatedCredentialProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateFederatedCredentialProviderResponse
	GetStatusCode() *int32
	SetBody(v *UpdateFederatedCredentialProviderResponseBody) *UpdateFederatedCredentialProviderResponse
	GetBody() *UpdateFederatedCredentialProviderResponseBody
}

type UpdateFederatedCredentialProviderResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFederatedCredentialProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFederatedCredentialProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateFederatedCredentialProviderResponse) GoString() string {
	return s.String()
}

func (s *UpdateFederatedCredentialProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateFederatedCredentialProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateFederatedCredentialProviderResponse) GetBody() *UpdateFederatedCredentialProviderResponseBody {
	return s.Body
}

func (s *UpdateFederatedCredentialProviderResponse) SetHeaders(v map[string]*string) *UpdateFederatedCredentialProviderResponse {
	s.Headers = v
	return s
}

func (s *UpdateFederatedCredentialProviderResponse) SetStatusCode(v int32) *UpdateFederatedCredentialProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFederatedCredentialProviderResponse) SetBody(v *UpdateFederatedCredentialProviderResponseBody) *UpdateFederatedCredentialProviderResponse {
	s.Body = v
	return s
}

func (s *UpdateFederatedCredentialProviderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
