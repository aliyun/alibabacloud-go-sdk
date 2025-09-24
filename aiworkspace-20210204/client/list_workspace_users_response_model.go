// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkspaceUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWorkspaceUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWorkspaceUsersResponse
	GetStatusCode() *int32
	SetBody(v *ListWorkspaceUsersResponseBody) *ListWorkspaceUsersResponse
	GetBody() *ListWorkspaceUsersResponseBody
}

type ListWorkspaceUsersResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWorkspaceUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWorkspaceUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspaceUsersResponse) GoString() string {
	return s.String()
}

func (s *ListWorkspaceUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWorkspaceUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWorkspaceUsersResponse) GetBody() *ListWorkspaceUsersResponseBody {
	return s.Body
}

func (s *ListWorkspaceUsersResponse) SetHeaders(v map[string]*string) *ListWorkspaceUsersResponse {
	s.Headers = v
	return s
}

func (s *ListWorkspaceUsersResponse) SetStatusCode(v int32) *ListWorkspaceUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkspaceUsersResponse) SetBody(v *ListWorkspaceUsersResponseBody) *ListWorkspaceUsersResponse {
	s.Body = v
	return s
}

func (s *ListWorkspaceUsersResponse) Validate() error {
	return dara.Validate(s)
}
