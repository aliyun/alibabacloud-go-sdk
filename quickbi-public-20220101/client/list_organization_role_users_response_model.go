// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOrganizationRoleUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOrganizationRoleUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOrganizationRoleUsersResponse
	GetStatusCode() *int32
	SetBody(v *ListOrganizationRoleUsersResponseBody) *ListOrganizationRoleUsersResponse
	GetBody() *ListOrganizationRoleUsersResponseBody
}

type ListOrganizationRoleUsersResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOrganizationRoleUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOrganizationRoleUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOrganizationRoleUsersResponse) GoString() string {
	return s.String()
}

func (s *ListOrganizationRoleUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOrganizationRoleUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOrganizationRoleUsersResponse) GetBody() *ListOrganizationRoleUsersResponseBody {
	return s.Body
}

func (s *ListOrganizationRoleUsersResponse) SetHeaders(v map[string]*string) *ListOrganizationRoleUsersResponse {
	s.Headers = v
	return s
}

func (s *ListOrganizationRoleUsersResponse) SetStatusCode(v int32) *ListOrganizationRoleUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOrganizationRoleUsersResponse) SetBody(v *ListOrganizationRoleUsersResponseBody) *ListOrganizationRoleUsersResponse {
	s.Body = v
	return s
}

func (s *ListOrganizationRoleUsersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
