// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgenticDBTenantApiKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAgenticDBTenantApiKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAgenticDBTenantApiKeyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAgenticDBTenantApiKeyResponseBody) *DeleteAgenticDBTenantApiKeyResponse
	GetBody() *DeleteAgenticDBTenantApiKeyResponseBody
}

type DeleteAgenticDBTenantApiKeyResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAgenticDBTenantApiKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAgenticDBTenantApiKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgenticDBTenantApiKeyResponse) GoString() string {
	return s.String()
}

func (s *DeleteAgenticDBTenantApiKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAgenticDBTenantApiKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAgenticDBTenantApiKeyResponse) GetBody() *DeleteAgenticDBTenantApiKeyResponseBody {
	return s.Body
}

func (s *DeleteAgenticDBTenantApiKeyResponse) SetHeaders(v map[string]*string) *DeleteAgenticDBTenantApiKeyResponse {
	s.Headers = v
	return s
}

func (s *DeleteAgenticDBTenantApiKeyResponse) SetStatusCode(v int32) *DeleteAgenticDBTenantApiKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAgenticDBTenantApiKeyResponse) SetBody(v *DeleteAgenticDBTenantApiKeyResponseBody) *DeleteAgenticDBTenantApiKeyResponse {
	s.Body = v
	return s
}

func (s *DeleteAgenticDBTenantApiKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
