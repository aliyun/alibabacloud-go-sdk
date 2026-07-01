// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyAgenticDBTenantApiKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VerifyAgenticDBTenantApiKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VerifyAgenticDBTenantApiKeyResponse
	GetStatusCode() *int32
	SetBody(v *VerifyAgenticDBTenantApiKeyResponseBody) *VerifyAgenticDBTenantApiKeyResponse
	GetBody() *VerifyAgenticDBTenantApiKeyResponseBody
}

type VerifyAgenticDBTenantApiKeyResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyAgenticDBTenantApiKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyAgenticDBTenantApiKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s VerifyAgenticDBTenantApiKeyResponse) GoString() string {
	return s.String()
}

func (s *VerifyAgenticDBTenantApiKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VerifyAgenticDBTenantApiKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VerifyAgenticDBTenantApiKeyResponse) GetBody() *VerifyAgenticDBTenantApiKeyResponseBody {
	return s.Body
}

func (s *VerifyAgenticDBTenantApiKeyResponse) SetHeaders(v map[string]*string) *VerifyAgenticDBTenantApiKeyResponse {
	s.Headers = v
	return s
}

func (s *VerifyAgenticDBTenantApiKeyResponse) SetStatusCode(v int32) *VerifyAgenticDBTenantApiKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyAgenticDBTenantApiKeyResponse) SetBody(v *VerifyAgenticDBTenantApiKeyResponseBody) *VerifyAgenticDBTenantApiKeyResponse {
	s.Body = v
	return s
}

func (s *VerifyAgenticDBTenantApiKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
