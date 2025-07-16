// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRoleUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRoleUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRoleUsersResponse
	GetStatusCode() *int32
}

type UpdateRoleUsersResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s UpdateRoleUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRoleUsersResponse) GoString() string {
	return s.String()
}

func (s *UpdateRoleUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRoleUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRoleUsersResponse) SetHeaders(v map[string]*string) *UpdateRoleUsersResponse {
	s.Headers = v
	return s
}

func (s *UpdateRoleUsersResponse) SetStatusCode(v int32) *UpdateRoleUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRoleUsersResponse) Validate() error {
	return dara.Validate(s)
}
