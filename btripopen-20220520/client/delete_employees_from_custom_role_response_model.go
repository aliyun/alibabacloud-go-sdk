// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEmployeesFromCustomRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEmployeesFromCustomRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEmployeesFromCustomRoleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEmployeesFromCustomRoleResponseBody) *DeleteEmployeesFromCustomRoleResponse
	GetBody() *DeleteEmployeesFromCustomRoleResponseBody
}

type DeleteEmployeesFromCustomRoleResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEmployeesFromCustomRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEmployeesFromCustomRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEmployeesFromCustomRoleResponse) GoString() string {
	return s.String()
}

func (s *DeleteEmployeesFromCustomRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEmployeesFromCustomRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEmployeesFromCustomRoleResponse) GetBody() *DeleteEmployeesFromCustomRoleResponseBody {
	return s.Body
}

func (s *DeleteEmployeesFromCustomRoleResponse) SetHeaders(v map[string]*string) *DeleteEmployeesFromCustomRoleResponse {
	s.Headers = v
	return s
}

func (s *DeleteEmployeesFromCustomRoleResponse) SetStatusCode(v int32) *DeleteEmployeesFromCustomRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEmployeesFromCustomRoleResponse) SetBody(v *DeleteEmployeesFromCustomRoleResponseBody) *DeleteEmployeesFromCustomRoleResponse {
	s.Body = v
	return s
}

func (s *DeleteEmployeesFromCustomRoleResponse) Validate() error {
	return dara.Validate(s)
}
