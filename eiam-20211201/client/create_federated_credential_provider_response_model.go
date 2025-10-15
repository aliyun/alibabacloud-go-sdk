// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFederatedCredentialProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateFederatedCredentialProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateFederatedCredentialProviderResponse
	GetStatusCode() *int32
	SetBody(v *CreateFederatedCredentialProviderResponseBody) *CreateFederatedCredentialProviderResponse
	GetBody() *CreateFederatedCredentialProviderResponseBody
}

type CreateFederatedCredentialProviderResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFederatedCredentialProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFederatedCredentialProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateFederatedCredentialProviderResponse) GoString() string {
	return s.String()
}

func (s *CreateFederatedCredentialProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateFederatedCredentialProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateFederatedCredentialProviderResponse) GetBody() *CreateFederatedCredentialProviderResponseBody {
	return s.Body
}

func (s *CreateFederatedCredentialProviderResponse) SetHeaders(v map[string]*string) *CreateFederatedCredentialProviderResponse {
	s.Headers = v
	return s
}

func (s *CreateFederatedCredentialProviderResponse) SetStatusCode(v int32) *CreateFederatedCredentialProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFederatedCredentialProviderResponse) SetBody(v *CreateFederatedCredentialProviderResponseBody) *CreateFederatedCredentialProviderResponse {
	s.Body = v
	return s
}

func (s *CreateFederatedCredentialProviderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
