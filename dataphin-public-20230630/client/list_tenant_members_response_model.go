// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTenantMembersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTenantMembersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTenantMembersResponse
	GetStatusCode() *int32
	SetBody(v *ListTenantMembersResponseBody) *ListTenantMembersResponse
	GetBody() *ListTenantMembersResponseBody
}

type ListTenantMembersResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTenantMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTenantMembersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTenantMembersResponse) GoString() string {
	return s.String()
}

func (s *ListTenantMembersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTenantMembersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTenantMembersResponse) GetBody() *ListTenantMembersResponseBody {
	return s.Body
}

func (s *ListTenantMembersResponse) SetHeaders(v map[string]*string) *ListTenantMembersResponse {
	s.Headers = v
	return s
}

func (s *ListTenantMembersResponse) SetStatusCode(v int32) *ListTenantMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTenantMembersResponse) SetBody(v *ListTenantMembersResponseBody) *ListTenantMembersResponse {
	s.Body = v
	return s
}

func (s *ListTenantMembersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
