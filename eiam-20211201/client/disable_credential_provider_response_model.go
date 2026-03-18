// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableCredentialProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableCredentialProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableCredentialProviderResponse
	GetStatusCode() *int32
	SetBody(v *DisableCredentialProviderResponseBody) *DisableCredentialProviderResponse
	GetBody() *DisableCredentialProviderResponseBody
}

type DisableCredentialProviderResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableCredentialProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableCredentialProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableCredentialProviderResponse) GoString() string {
	return s.String()
}

func (s *DisableCredentialProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableCredentialProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableCredentialProviderResponse) GetBody() *DisableCredentialProviderResponseBody {
	return s.Body
}

func (s *DisableCredentialProviderResponse) SetHeaders(v map[string]*string) *DisableCredentialProviderResponse {
	s.Headers = v
	return s
}

func (s *DisableCredentialProviderResponse) SetStatusCode(v int32) *DisableCredentialProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableCredentialProviderResponse) SetBody(v *DisableCredentialProviderResponseBody) *DisableCredentialProviderResponse {
	s.Body = v
	return s
}

func (s *DisableCredentialProviderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
