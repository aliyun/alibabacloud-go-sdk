// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableFederatedCredentialProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableFederatedCredentialProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableFederatedCredentialProviderResponse
	GetStatusCode() *int32
	SetBody(v *DisableFederatedCredentialProviderResponseBody) *DisableFederatedCredentialProviderResponse
	GetBody() *DisableFederatedCredentialProviderResponseBody
}

type DisableFederatedCredentialProviderResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableFederatedCredentialProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableFederatedCredentialProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableFederatedCredentialProviderResponse) GoString() string {
	return s.String()
}

func (s *DisableFederatedCredentialProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableFederatedCredentialProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableFederatedCredentialProviderResponse) GetBody() *DisableFederatedCredentialProviderResponseBody {
	return s.Body
}

func (s *DisableFederatedCredentialProviderResponse) SetHeaders(v map[string]*string) *DisableFederatedCredentialProviderResponse {
	s.Headers = v
	return s
}

func (s *DisableFederatedCredentialProviderResponse) SetStatusCode(v int32) *DisableFederatedCredentialProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableFederatedCredentialProviderResponse) SetBody(v *DisableFederatedCredentialProviderResponseBody) *DisableFederatedCredentialProviderResponse {
	s.Body = v
	return s
}

func (s *DisableFederatedCredentialProviderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
