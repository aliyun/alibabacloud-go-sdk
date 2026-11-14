// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTenantMemoryConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyTenantMemoryConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyTenantMemoryConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyTenantMemoryConfigResponseBody) *ModifyTenantMemoryConfigResponse
	GetBody() *ModifyTenantMemoryConfigResponseBody
}

type ModifyTenantMemoryConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyTenantMemoryConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyTenantMemoryConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyTenantMemoryConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyTenantMemoryConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyTenantMemoryConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyTenantMemoryConfigResponse) GetBody() *ModifyTenantMemoryConfigResponseBody {
	return s.Body
}

func (s *ModifyTenantMemoryConfigResponse) SetHeaders(v map[string]*string) *ModifyTenantMemoryConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyTenantMemoryConfigResponse) SetStatusCode(v int32) *ModifyTenantMemoryConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyTenantMemoryConfigResponse) SetBody(v *ModifyTenantMemoryConfigResponseBody) *ModifyTenantMemoryConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyTenantMemoryConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
