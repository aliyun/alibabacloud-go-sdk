// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTenantConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyTenantConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyTenantConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyTenantConfigResponseBody) *ModifyTenantConfigResponse
	GetBody() *ModifyTenantConfigResponseBody
}

type ModifyTenantConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyTenantConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyTenantConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyTenantConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyTenantConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyTenantConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyTenantConfigResponse) GetBody() *ModifyTenantConfigResponseBody {
	return s.Body
}

func (s *ModifyTenantConfigResponse) SetHeaders(v map[string]*string) *ModifyTenantConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyTenantConfigResponse) SetStatusCode(v int32) *ModifyTenantConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyTenantConfigResponse) SetBody(v *ModifyTenantConfigResponseBody) *ModifyTenantConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyTenantConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
