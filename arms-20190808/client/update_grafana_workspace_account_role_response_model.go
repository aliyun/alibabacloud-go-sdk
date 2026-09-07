// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGrafanaWorkspaceAccountRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGrafanaWorkspaceAccountRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGrafanaWorkspaceAccountRoleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGrafanaWorkspaceAccountRoleResponseBody) *UpdateGrafanaWorkspaceAccountRoleResponse
	GetBody() *UpdateGrafanaWorkspaceAccountRoleResponseBody
}

type UpdateGrafanaWorkspaceAccountRoleResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGrafanaWorkspaceAccountRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGrafanaWorkspaceAccountRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGrafanaWorkspaceAccountRoleResponse) GoString() string {
	return s.String()
}

func (s *UpdateGrafanaWorkspaceAccountRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGrafanaWorkspaceAccountRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGrafanaWorkspaceAccountRoleResponse) GetBody() *UpdateGrafanaWorkspaceAccountRoleResponseBody {
	return s.Body
}

func (s *UpdateGrafanaWorkspaceAccountRoleResponse) SetHeaders(v map[string]*string) *UpdateGrafanaWorkspaceAccountRoleResponse {
	s.Headers = v
	return s
}

func (s *UpdateGrafanaWorkspaceAccountRoleResponse) SetStatusCode(v int32) *UpdateGrafanaWorkspaceAccountRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGrafanaWorkspaceAccountRoleResponse) SetBody(v *UpdateGrafanaWorkspaceAccountRoleResponseBody) *UpdateGrafanaWorkspaceAccountRoleResponse {
	s.Body = v
	return s
}

func (s *UpdateGrafanaWorkspaceAccountRoleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
