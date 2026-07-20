// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignRbacUserRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssignRbacUserRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssignRbacUserRoleResponse
	GetStatusCode() *int32
	SetBody(v *AssignRbacUserRoleResponseBody) *AssignRbacUserRoleResponse
	GetBody() *AssignRbacUserRoleResponseBody
}

type AssignRbacUserRoleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssignRbacUserRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssignRbacUserRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s AssignRbacUserRoleResponse) GoString() string {
	return s.String()
}

func (s *AssignRbacUserRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssignRbacUserRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssignRbacUserRoleResponse) GetBody() *AssignRbacUserRoleResponseBody {
	return s.Body
}

func (s *AssignRbacUserRoleResponse) SetHeaders(v map[string]*string) *AssignRbacUserRoleResponse {
	s.Headers = v
	return s
}

func (s *AssignRbacUserRoleResponse) SetStatusCode(v int32) *AssignRbacUserRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *AssignRbacUserRoleResponse) SetBody(v *AssignRbacUserRoleResponseBody) *AssignRbacUserRoleResponse {
	s.Body = v
	return s
}

func (s *AssignRbacUserRoleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
