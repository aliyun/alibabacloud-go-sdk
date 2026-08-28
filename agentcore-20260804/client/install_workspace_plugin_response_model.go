// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallWorkspacePluginResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InstallWorkspacePluginResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InstallWorkspacePluginResponse
	GetStatusCode() *int32
	SetBody(v *InstallWorkspacePluginResponseBody) *InstallWorkspacePluginResponse
	GetBody() *InstallWorkspacePluginResponseBody
}

type InstallWorkspacePluginResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstallWorkspacePluginResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallWorkspacePluginResponse) String() string {
	return dara.Prettify(s)
}

func (s InstallWorkspacePluginResponse) GoString() string {
	return s.String()
}

func (s *InstallWorkspacePluginResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InstallWorkspacePluginResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InstallWorkspacePluginResponse) GetBody() *InstallWorkspacePluginResponseBody {
	return s.Body
}

func (s *InstallWorkspacePluginResponse) SetHeaders(v map[string]*string) *InstallWorkspacePluginResponse {
	s.Headers = v
	return s
}

func (s *InstallWorkspacePluginResponse) SetStatusCode(v int32) *InstallWorkspacePluginResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallWorkspacePluginResponse) SetBody(v *InstallWorkspacePluginResponseBody) *InstallWorkspacePluginResponse {
	s.Body = v
	return s
}

func (s *InstallWorkspacePluginResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
