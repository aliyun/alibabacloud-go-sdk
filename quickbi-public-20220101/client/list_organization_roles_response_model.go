// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOrganizationRolesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOrganizationRolesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOrganizationRolesResponse
	GetStatusCode() *int32
	SetBody(v *ListOrganizationRolesResponseBody) *ListOrganizationRolesResponse
	GetBody() *ListOrganizationRolesResponseBody
}

type ListOrganizationRolesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOrganizationRolesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOrganizationRolesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOrganizationRolesResponse) GoString() string {
	return s.String()
}

func (s *ListOrganizationRolesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOrganizationRolesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOrganizationRolesResponse) GetBody() *ListOrganizationRolesResponseBody {
	return s.Body
}

func (s *ListOrganizationRolesResponse) SetHeaders(v map[string]*string) *ListOrganizationRolesResponse {
	s.Headers = v
	return s
}

func (s *ListOrganizationRolesResponse) SetStatusCode(v int32) *ListOrganizationRolesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOrganizationRolesResponse) SetBody(v *ListOrganizationRolesResponseBody) *ListOrganizationRolesResponse {
	s.Body = v
	return s
}

func (s *ListOrganizationRolesResponse) Validate() error {
	return dara.Validate(s)
}
