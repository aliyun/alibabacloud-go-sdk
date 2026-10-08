// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTenantRolesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTenantRolesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTenantRolesResponse
	GetStatusCode() *int32
	SetBody(v *ListTenantRolesResponseBody) *ListTenantRolesResponse
	GetBody() *ListTenantRolesResponseBody
}

type ListTenantRolesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTenantRolesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTenantRolesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTenantRolesResponse) GoString() string {
	return s.String()
}

func (s *ListTenantRolesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTenantRolesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTenantRolesResponse) GetBody() *ListTenantRolesResponseBody {
	return s.Body
}

func (s *ListTenantRolesResponse) SetHeaders(v map[string]*string) *ListTenantRolesResponse {
	s.Headers = v
	return s
}

func (s *ListTenantRolesResponse) SetStatusCode(v int32) *ListTenantRolesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTenantRolesResponse) SetBody(v *ListTenantRolesResponseBody) *ListTenantRolesResponse {
	s.Body = v
	return s
}

func (s *ListTenantRolesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
