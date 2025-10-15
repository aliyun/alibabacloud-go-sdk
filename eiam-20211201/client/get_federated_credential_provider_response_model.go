// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFederatedCredentialProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFederatedCredentialProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFederatedCredentialProviderResponse
	GetStatusCode() *int32
	SetBody(v *GetFederatedCredentialProviderResponseBody) *GetFederatedCredentialProviderResponse
	GetBody() *GetFederatedCredentialProviderResponseBody
}

type GetFederatedCredentialProviderResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFederatedCredentialProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFederatedCredentialProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFederatedCredentialProviderResponse) GoString() string {
	return s.String()
}

func (s *GetFederatedCredentialProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFederatedCredentialProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFederatedCredentialProviderResponse) GetBody() *GetFederatedCredentialProviderResponseBody {
	return s.Body
}

func (s *GetFederatedCredentialProviderResponse) SetHeaders(v map[string]*string) *GetFederatedCredentialProviderResponse {
	s.Headers = v
	return s
}

func (s *GetFederatedCredentialProviderResponse) SetStatusCode(v int32) *GetFederatedCredentialProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFederatedCredentialProviderResponse) SetBody(v *GetFederatedCredentialProviderResponseBody) *GetFederatedCredentialProviderResponse {
	s.Body = v
	return s
}

func (s *GetFederatedCredentialProviderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
