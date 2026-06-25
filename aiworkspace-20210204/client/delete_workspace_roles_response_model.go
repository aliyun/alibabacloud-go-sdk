// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkspaceRolesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteWorkspaceRolesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteWorkspaceRolesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteWorkspaceRolesResponseBody) *DeleteWorkspaceRolesResponse
	GetBody() *DeleteWorkspaceRolesResponseBody
}

type DeleteWorkspaceRolesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWorkspaceRolesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWorkspaceRolesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkspaceRolesResponse) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceRolesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteWorkspaceRolesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteWorkspaceRolesResponse) GetBody() *DeleteWorkspaceRolesResponseBody {
	return s.Body
}

func (s *DeleteWorkspaceRolesResponse) SetHeaders(v map[string]*string) *DeleteWorkspaceRolesResponse {
	s.Headers = v
	return s
}

func (s *DeleteWorkspaceRolesResponse) SetStatusCode(v int32) *DeleteWorkspaceRolesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWorkspaceRolesResponse) SetBody(v *DeleteWorkspaceRolesResponseBody) *DeleteWorkspaceRolesResponse {
	s.Body = v
	return s
}

func (s *DeleteWorkspaceRolesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
