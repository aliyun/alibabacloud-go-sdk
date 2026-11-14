// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTenantMemoryConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTenantMemoryConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTenantMemoryConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetTenantMemoryConfigResponseBody) *GetTenantMemoryConfigResponse
	GetBody() *GetTenantMemoryConfigResponseBody
}

type GetTenantMemoryConfigResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTenantMemoryConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTenantMemoryConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTenantMemoryConfigResponse) GoString() string {
	return s.String()
}

func (s *GetTenantMemoryConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTenantMemoryConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTenantMemoryConfigResponse) GetBody() *GetTenantMemoryConfigResponseBody {
	return s.Body
}

func (s *GetTenantMemoryConfigResponse) SetHeaders(v map[string]*string) *GetTenantMemoryConfigResponse {
	s.Headers = v
	return s
}

func (s *GetTenantMemoryConfigResponse) SetStatusCode(v int32) *GetTenantMemoryConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTenantMemoryConfigResponse) SetBody(v *GetTenantMemoryConfigResponseBody) *GetTenantMemoryConfigResponse {
	s.Body = v
	return s
}

func (s *GetTenantMemoryConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
