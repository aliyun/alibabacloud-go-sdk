// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFederatedCredentialProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFederatedCredentialProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFederatedCredentialProviderResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFederatedCredentialProviderResponseBody) *DeleteFederatedCredentialProviderResponse
	GetBody() *DeleteFederatedCredentialProviderResponseBody
}

type DeleteFederatedCredentialProviderResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFederatedCredentialProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFederatedCredentialProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFederatedCredentialProviderResponse) GoString() string {
	return s.String()
}

func (s *DeleteFederatedCredentialProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFederatedCredentialProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFederatedCredentialProviderResponse) GetBody() *DeleteFederatedCredentialProviderResponseBody {
	return s.Body
}

func (s *DeleteFederatedCredentialProviderResponse) SetHeaders(v map[string]*string) *DeleteFederatedCredentialProviderResponse {
	s.Headers = v
	return s
}

func (s *DeleteFederatedCredentialProviderResponse) SetStatusCode(v int32) *DeleteFederatedCredentialProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFederatedCredentialProviderResponse) SetBody(v *DeleteFederatedCredentialProviderResponseBody) *DeleteFederatedCredentialProviderResponse {
	s.Body = v
	return s
}

func (s *DeleteFederatedCredentialProviderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
