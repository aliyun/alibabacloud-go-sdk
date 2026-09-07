// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGrafanaWorkspaceAccountRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteGrafanaWorkspaceAccountRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteGrafanaWorkspaceAccountRoleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteGrafanaWorkspaceAccountRoleResponseBody) *DeleteGrafanaWorkspaceAccountRoleResponse
	GetBody() *DeleteGrafanaWorkspaceAccountRoleResponseBody
}

type DeleteGrafanaWorkspaceAccountRoleResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGrafanaWorkspaceAccountRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGrafanaWorkspaceAccountRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteGrafanaWorkspaceAccountRoleResponse) GoString() string {
	return s.String()
}

func (s *DeleteGrafanaWorkspaceAccountRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteGrafanaWorkspaceAccountRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteGrafanaWorkspaceAccountRoleResponse) GetBody() *DeleteGrafanaWorkspaceAccountRoleResponseBody {
	return s.Body
}

func (s *DeleteGrafanaWorkspaceAccountRoleResponse) SetHeaders(v map[string]*string) *DeleteGrafanaWorkspaceAccountRoleResponse {
	s.Headers = v
	return s
}

func (s *DeleteGrafanaWorkspaceAccountRoleResponse) SetStatusCode(v int32) *DeleteGrafanaWorkspaceAccountRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGrafanaWorkspaceAccountRoleResponse) SetBody(v *DeleteGrafanaWorkspaceAccountRoleResponseBody) *DeleteGrafanaWorkspaceAccountRoleResponse {
	s.Body = v
	return s
}

func (s *DeleteGrafanaWorkspaceAccountRoleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
