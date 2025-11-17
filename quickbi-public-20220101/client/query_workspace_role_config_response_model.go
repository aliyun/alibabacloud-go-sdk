// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryWorkspaceRoleConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryWorkspaceRoleConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryWorkspaceRoleConfigResponse
	GetStatusCode() *int32
	SetBody(v *QueryWorkspaceRoleConfigResponseBody) *QueryWorkspaceRoleConfigResponse
	GetBody() *QueryWorkspaceRoleConfigResponseBody
}

type QueryWorkspaceRoleConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryWorkspaceRoleConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryWorkspaceRoleConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryWorkspaceRoleConfigResponse) GoString() string {
	return s.String()
}

func (s *QueryWorkspaceRoleConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryWorkspaceRoleConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryWorkspaceRoleConfigResponse) GetBody() *QueryWorkspaceRoleConfigResponseBody {
	return s.Body
}

func (s *QueryWorkspaceRoleConfigResponse) SetHeaders(v map[string]*string) *QueryWorkspaceRoleConfigResponse {
	s.Headers = v
	return s
}

func (s *QueryWorkspaceRoleConfigResponse) SetStatusCode(v int32) *QueryWorkspaceRoleConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryWorkspaceRoleConfigResponse) SetBody(v *QueryWorkspaceRoleConfigResponseBody) *QueryWorkspaceRoleConfigResponse {
	s.Body = v
	return s
}

func (s *QueryWorkspaceRoleConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
