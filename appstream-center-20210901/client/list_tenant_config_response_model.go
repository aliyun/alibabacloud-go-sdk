// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTenantConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTenantConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTenantConfigResponse
	GetStatusCode() *int32
	SetBody(v *ListTenantConfigResponseBody) *ListTenantConfigResponse
	GetBody() *ListTenantConfigResponseBody
}

type ListTenantConfigResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTenantConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTenantConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTenantConfigResponse) GoString() string {
	return s.String()
}

func (s *ListTenantConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTenantConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTenantConfigResponse) GetBody() *ListTenantConfigResponseBody {
	return s.Body
}

func (s *ListTenantConfigResponse) SetHeaders(v map[string]*string) *ListTenantConfigResponse {
	s.Headers = v
	return s
}

func (s *ListTenantConfigResponse) SetStatusCode(v int32) *ListTenantConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTenantConfigResponse) SetBody(v *ListTenantConfigResponseBody) *ListTenantConfigResponse {
	s.Body = v
	return s
}

func (s *ListTenantConfigResponse) Validate() error {
	return dara.Validate(s)
}
