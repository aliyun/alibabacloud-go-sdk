// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataAgentWorkspaceMemberRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDataAgentWorkspaceMemberRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDataAgentWorkspaceMemberRoleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDataAgentWorkspaceMemberRoleResponseBody) *UpdateDataAgentWorkspaceMemberRoleResponse
	GetBody() *UpdateDataAgentWorkspaceMemberRoleResponseBody
}

type UpdateDataAgentWorkspaceMemberRoleResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDataAgentWorkspaceMemberRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDataAgentWorkspaceMemberRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataAgentWorkspaceMemberRoleResponse) GoString() string {
	return s.String()
}

func (s *UpdateDataAgentWorkspaceMemberRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDataAgentWorkspaceMemberRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDataAgentWorkspaceMemberRoleResponse) GetBody() *UpdateDataAgentWorkspaceMemberRoleResponseBody {
	return s.Body
}

func (s *UpdateDataAgentWorkspaceMemberRoleResponse) SetHeaders(v map[string]*string) *UpdateDataAgentWorkspaceMemberRoleResponse {
	s.Headers = v
	return s
}

func (s *UpdateDataAgentWorkspaceMemberRoleResponse) SetStatusCode(v int32) *UpdateDataAgentWorkspaceMemberRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDataAgentWorkspaceMemberRoleResponse) SetBody(v *UpdateDataAgentWorkspaceMemberRoleResponseBody) *UpdateDataAgentWorkspaceMemberRoleResponse {
	s.Body = v
	return s
}

func (s *UpdateDataAgentWorkspaceMemberRoleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
