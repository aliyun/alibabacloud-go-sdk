// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRbacRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRbacRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRbacRoleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRbacRoleResponseBody) *UpdateRbacRoleResponse
	GetBody() *UpdateRbacRoleResponseBody
}

type UpdateRbacRoleResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRbacRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRbacRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRbacRoleResponse) GoString() string {
	return s.String()
}

func (s *UpdateRbacRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRbacRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRbacRoleResponse) GetBody() *UpdateRbacRoleResponseBody {
	return s.Body
}

func (s *UpdateRbacRoleResponse) SetHeaders(v map[string]*string) *UpdateRbacRoleResponse {
	s.Headers = v
	return s
}

func (s *UpdateRbacRoleResponse) SetStatusCode(v int32) *UpdateRbacRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRbacRoleResponse) SetBody(v *UpdateRbacRoleResponseBody) *UpdateRbacRoleResponse {
	s.Body = v
	return s
}

func (s *UpdateRbacRoleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
