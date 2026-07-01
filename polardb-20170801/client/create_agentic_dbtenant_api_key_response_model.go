// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgenticDBTenantApiKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAgenticDBTenantApiKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAgenticDBTenantApiKeyResponse
	GetStatusCode() *int32
	SetBody(v *CreateAgenticDBTenantApiKeyResponseBody) *CreateAgenticDBTenantApiKeyResponse
	GetBody() *CreateAgenticDBTenantApiKeyResponseBody
}

type CreateAgenticDBTenantApiKeyResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAgenticDBTenantApiKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAgenticDBTenantApiKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAgenticDBTenantApiKeyResponse) GoString() string {
	return s.String()
}

func (s *CreateAgenticDBTenantApiKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAgenticDBTenantApiKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAgenticDBTenantApiKeyResponse) GetBody() *CreateAgenticDBTenantApiKeyResponseBody {
	return s.Body
}

func (s *CreateAgenticDBTenantApiKeyResponse) SetHeaders(v map[string]*string) *CreateAgenticDBTenantApiKeyResponse {
	s.Headers = v
	return s
}

func (s *CreateAgenticDBTenantApiKeyResponse) SetStatusCode(v int32) *CreateAgenticDBTenantApiKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAgenticDBTenantApiKeyResponse) SetBody(v *CreateAgenticDBTenantApiKeyResponseBody) *CreateAgenticDBTenantApiKeyResponse {
	s.Body = v
	return s
}

func (s *CreateAgenticDBTenantApiKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
