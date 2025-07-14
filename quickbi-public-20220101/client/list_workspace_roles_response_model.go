// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkspaceRolesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWorkspaceRolesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWorkspaceRolesResponse
	GetStatusCode() *int32
	SetBody(v *ListWorkspaceRolesResponseBody) *ListWorkspaceRolesResponse
	GetBody() *ListWorkspaceRolesResponseBody
}

type ListWorkspaceRolesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWorkspaceRolesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWorkspaceRolesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspaceRolesResponse) GoString() string {
	return s.String()
}

func (s *ListWorkspaceRolesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWorkspaceRolesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWorkspaceRolesResponse) GetBody() *ListWorkspaceRolesResponseBody {
	return s.Body
}

func (s *ListWorkspaceRolesResponse) SetHeaders(v map[string]*string) *ListWorkspaceRolesResponse {
	s.Headers = v
	return s
}

func (s *ListWorkspaceRolesResponse) SetStatusCode(v int32) *ListWorkspaceRolesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkspaceRolesResponse) SetBody(v *ListWorkspaceRolesResponseBody) *ListWorkspaceRolesResponse {
	s.Body = v
	return s
}

func (s *ListWorkspaceRolesResponse) Validate() error {
	return dara.Validate(s)
}
