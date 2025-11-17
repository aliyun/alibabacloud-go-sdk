// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserRoleInfoInWorkspaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryUserRoleInfoInWorkspaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryUserRoleInfoInWorkspaceResponse
	GetStatusCode() *int32
	SetBody(v *QueryUserRoleInfoInWorkspaceResponseBody) *QueryUserRoleInfoInWorkspaceResponse
	GetBody() *QueryUserRoleInfoInWorkspaceResponseBody
}

type QueryUserRoleInfoInWorkspaceResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryUserRoleInfoInWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryUserRoleInfoInWorkspaceResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryUserRoleInfoInWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *QueryUserRoleInfoInWorkspaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryUserRoleInfoInWorkspaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryUserRoleInfoInWorkspaceResponse) GetBody() *QueryUserRoleInfoInWorkspaceResponseBody {
	return s.Body
}

func (s *QueryUserRoleInfoInWorkspaceResponse) SetHeaders(v map[string]*string) *QueryUserRoleInfoInWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *QueryUserRoleInfoInWorkspaceResponse) SetStatusCode(v int32) *QueryUserRoleInfoInWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUserRoleInfoInWorkspaceResponse) SetBody(v *QueryUserRoleInfoInWorkspaceResponseBody) *QueryUserRoleInfoInWorkspaceResponse {
	s.Body = v
	return s
}

func (s *QueryUserRoleInfoInWorkspaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
