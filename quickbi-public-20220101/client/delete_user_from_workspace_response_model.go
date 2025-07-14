// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserFromWorkspaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteUserFromWorkspaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteUserFromWorkspaceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteUserFromWorkspaceResponseBody) *DeleteUserFromWorkspaceResponse
	GetBody() *DeleteUserFromWorkspaceResponseBody
}

type DeleteUserFromWorkspaceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserFromWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUserFromWorkspaceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserFromWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserFromWorkspaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteUserFromWorkspaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteUserFromWorkspaceResponse) GetBody() *DeleteUserFromWorkspaceResponseBody {
	return s.Body
}

func (s *DeleteUserFromWorkspaceResponse) SetHeaders(v map[string]*string) *DeleteUserFromWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserFromWorkspaceResponse) SetStatusCode(v int32) *DeleteUserFromWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserFromWorkspaceResponse) SetBody(v *DeleteUserFromWorkspaceResponseBody) *DeleteUserFromWorkspaceResponse {
	s.Body = v
	return s
}

func (s *DeleteUserFromWorkspaceResponse) Validate() error {
	return dara.Validate(s)
}
