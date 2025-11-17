// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkspaceUserRolesByUserIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWorkspaceUserRolesByUserIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWorkspaceUserRolesByUserIdResponse
	GetStatusCode() *int32
	SetBody(v *ListWorkspaceUserRolesByUserIdResponseBody) *ListWorkspaceUserRolesByUserIdResponse
	GetBody() *ListWorkspaceUserRolesByUserIdResponseBody
}

type ListWorkspaceUserRolesByUserIdResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWorkspaceUserRolesByUserIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWorkspaceUserRolesByUserIdResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspaceUserRolesByUserIdResponse) GoString() string {
	return s.String()
}

func (s *ListWorkspaceUserRolesByUserIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWorkspaceUserRolesByUserIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWorkspaceUserRolesByUserIdResponse) GetBody() *ListWorkspaceUserRolesByUserIdResponseBody {
	return s.Body
}

func (s *ListWorkspaceUserRolesByUserIdResponse) SetHeaders(v map[string]*string) *ListWorkspaceUserRolesByUserIdResponse {
	s.Headers = v
	return s
}

func (s *ListWorkspaceUserRolesByUserIdResponse) SetStatusCode(v int32) *ListWorkspaceUserRolesByUserIdResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkspaceUserRolesByUserIdResponse) SetBody(v *ListWorkspaceUserRolesByUserIdResponseBody) *ListWorkspaceUserRolesByUserIdResponse {
	s.Body = v
	return s
}

func (s *ListWorkspaceUserRolesByUserIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
