// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProjectRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteProjectRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteProjectRoleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteProjectRoleResponseBody) *DeleteProjectRoleResponse
	GetBody() *DeleteProjectRoleResponseBody
}

type DeleteProjectRoleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteProjectRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteProjectRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteProjectRoleResponse) GoString() string {
	return s.String()
}

func (s *DeleteProjectRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteProjectRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteProjectRoleResponse) GetBody() *DeleteProjectRoleResponseBody {
	return s.Body
}

func (s *DeleteProjectRoleResponse) SetHeaders(v map[string]*string) *DeleteProjectRoleResponse {
	s.Headers = v
	return s
}

func (s *DeleteProjectRoleResponse) SetStatusCode(v int32) *DeleteProjectRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProjectRoleResponse) SetBody(v *DeleteProjectRoleResponseBody) *DeleteProjectRoleResponse {
	s.Body = v
	return s
}

func (s *DeleteProjectRoleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
