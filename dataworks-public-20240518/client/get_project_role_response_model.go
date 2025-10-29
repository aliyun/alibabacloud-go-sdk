// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProjectRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetProjectRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetProjectRoleResponse
	GetStatusCode() *int32
	SetBody(v *GetProjectRoleResponseBody) *GetProjectRoleResponse
	GetBody() *GetProjectRoleResponseBody
}

type GetProjectRoleResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProjectRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProjectRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetProjectRoleResponse) GoString() string {
	return s.String()
}

func (s *GetProjectRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetProjectRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetProjectRoleResponse) GetBody() *GetProjectRoleResponseBody {
	return s.Body
}

func (s *GetProjectRoleResponse) SetHeaders(v map[string]*string) *GetProjectRoleResponse {
	s.Headers = v
	return s
}

func (s *GetProjectRoleResponse) SetStatusCode(v int32) *GetProjectRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProjectRoleResponse) SetBody(v *GetProjectRoleResponseBody) *GetProjectRoleResponse {
	s.Body = v
	return s
}

func (s *GetProjectRoleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
