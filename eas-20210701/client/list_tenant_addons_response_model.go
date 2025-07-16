// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTenantAddonsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTenantAddonsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTenantAddonsResponse
	GetStatusCode() *int32
	SetBody(v *ListTenantAddonsResponseBody) *ListTenantAddonsResponse
	GetBody() *ListTenantAddonsResponseBody
}

type ListTenantAddonsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTenantAddonsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTenantAddonsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTenantAddonsResponse) GoString() string {
	return s.String()
}

func (s *ListTenantAddonsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTenantAddonsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTenantAddonsResponse) GetBody() *ListTenantAddonsResponseBody {
	return s.Body
}

func (s *ListTenantAddonsResponse) SetHeaders(v map[string]*string) *ListTenantAddonsResponse {
	s.Headers = v
	return s
}

func (s *ListTenantAddonsResponse) SetStatusCode(v int32) *ListTenantAddonsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTenantAddonsResponse) SetBody(v *ListTenantAddonsResponseBody) *ListTenantAddonsResponse {
	s.Body = v
	return s
}

func (s *ListTenantAddonsResponse) Validate() error {
	return dara.Validate(s)
}
