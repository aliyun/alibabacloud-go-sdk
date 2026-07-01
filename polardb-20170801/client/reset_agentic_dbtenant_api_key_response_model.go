// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetAgenticDBTenantApiKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetAgenticDBTenantApiKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetAgenticDBTenantApiKeyResponse
	GetStatusCode() *int32
	SetBody(v *ResetAgenticDBTenantApiKeyResponseBody) *ResetAgenticDBTenantApiKeyResponse
	GetBody() *ResetAgenticDBTenantApiKeyResponseBody
}

type ResetAgenticDBTenantApiKeyResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetAgenticDBTenantApiKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetAgenticDBTenantApiKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetAgenticDBTenantApiKeyResponse) GoString() string {
	return s.String()
}

func (s *ResetAgenticDBTenantApiKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetAgenticDBTenantApiKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetAgenticDBTenantApiKeyResponse) GetBody() *ResetAgenticDBTenantApiKeyResponseBody {
	return s.Body
}

func (s *ResetAgenticDBTenantApiKeyResponse) SetHeaders(v map[string]*string) *ResetAgenticDBTenantApiKeyResponse {
	s.Headers = v
	return s
}

func (s *ResetAgenticDBTenantApiKeyResponse) SetStatusCode(v int32) *ResetAgenticDBTenantApiKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetAgenticDBTenantApiKeyResponse) SetBody(v *ResetAgenticDBTenantApiKeyResponseBody) *ResetAgenticDBTenantApiKeyResponse {
	s.Body = v
	return s
}

func (s *ResetAgenticDBTenantApiKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
