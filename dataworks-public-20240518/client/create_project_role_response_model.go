// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProjectRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateProjectRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateProjectRoleResponse
	GetStatusCode() *int32
	SetBody(v *CreateProjectRoleResponseBody) *CreateProjectRoleResponse
	GetBody() *CreateProjectRoleResponseBody
}

type CreateProjectRoleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateProjectRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateProjectRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectRoleResponse) GoString() string {
	return s.String()
}

func (s *CreateProjectRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateProjectRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateProjectRoleResponse) GetBody() *CreateProjectRoleResponseBody {
	return s.Body
}

func (s *CreateProjectRoleResponse) SetHeaders(v map[string]*string) *CreateProjectRoleResponse {
	s.Headers = v
	return s
}

func (s *CreateProjectRoleResponse) SetStatusCode(v int32) *CreateProjectRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProjectRoleResponse) SetBody(v *CreateProjectRoleResponseBody) *CreateProjectRoleResponse {
	s.Body = v
	return s
}

func (s *CreateProjectRoleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
