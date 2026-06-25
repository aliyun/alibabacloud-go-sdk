// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkspaceRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWorkspaceRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWorkspaceRoleResponse
	GetStatusCode() *int32
	SetBody(v *GetWorkspaceRoleResponseBody) *GetWorkspaceRoleResponse
	GetBody() *GetWorkspaceRoleResponseBody
}

type GetWorkspaceRoleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorkspaceRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWorkspaceRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceRoleResponse) GoString() string {
	return s.String()
}

func (s *GetWorkspaceRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWorkspaceRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWorkspaceRoleResponse) GetBody() *GetWorkspaceRoleResponseBody {
	return s.Body
}

func (s *GetWorkspaceRoleResponse) SetHeaders(v map[string]*string) *GetWorkspaceRoleResponse {
	s.Headers = v
	return s
}

func (s *GetWorkspaceRoleResponse) SetStatusCode(v int32) *GetWorkspaceRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkspaceRoleResponse) SetBody(v *GetWorkspaceRoleResponseBody) *GetWorkspaceRoleResponse {
	s.Body = v
	return s
}

func (s *GetWorkspaceRoleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
