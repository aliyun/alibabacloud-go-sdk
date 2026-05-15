// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProjectRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateProjectRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateProjectRoleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateProjectRoleResponseBody) *UpdateProjectRoleResponse
	GetBody() *UpdateProjectRoleResponseBody
}

type UpdateProjectRoleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateProjectRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateProjectRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectRoleResponse) GoString() string {
	return s.String()
}

func (s *UpdateProjectRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateProjectRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateProjectRoleResponse) GetBody() *UpdateProjectRoleResponseBody {
	return s.Body
}

func (s *UpdateProjectRoleResponse) SetHeaders(v map[string]*string) *UpdateProjectRoleResponse {
	s.Headers = v
	return s
}

func (s *UpdateProjectRoleResponse) SetStatusCode(v int32) *UpdateProjectRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateProjectRoleResponse) SetBody(v *UpdateProjectRoleResponseBody) *UpdateProjectRoleResponse {
	s.Body = v
	return s
}

func (s *UpdateProjectRoleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
