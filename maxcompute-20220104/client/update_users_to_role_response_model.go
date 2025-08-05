// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUsersToRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateUsersToRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateUsersToRoleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateUsersToRoleResponseBody) *UpdateUsersToRoleResponse
	GetBody() *UpdateUsersToRoleResponseBody
}

type UpdateUsersToRoleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUsersToRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUsersToRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateUsersToRoleResponse) GoString() string {
	return s.String()
}

func (s *UpdateUsersToRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateUsersToRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateUsersToRoleResponse) GetBody() *UpdateUsersToRoleResponseBody {
	return s.Body
}

func (s *UpdateUsersToRoleResponse) SetHeaders(v map[string]*string) *UpdateUsersToRoleResponse {
	s.Headers = v
	return s
}

func (s *UpdateUsersToRoleResponse) SetStatusCode(v int32) *UpdateUsersToRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUsersToRoleResponse) SetBody(v *UpdateUsersToRoleResponseBody) *UpdateUsersToRoleResponse {
	s.Body = v
	return s
}

func (s *UpdateUsersToRoleResponse) Validate() error {
	return dara.Validate(s)
}
