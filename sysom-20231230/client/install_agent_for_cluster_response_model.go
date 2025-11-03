// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallAgentForClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InstallAgentForClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InstallAgentForClusterResponse
	GetStatusCode() *int32
	SetBody(v *InstallAgentForClusterResponseBody) *InstallAgentForClusterResponse
	GetBody() *InstallAgentForClusterResponseBody
}

type InstallAgentForClusterResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstallAgentForClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallAgentForClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s InstallAgentForClusterResponse) GoString() string {
	return s.String()
}

func (s *InstallAgentForClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InstallAgentForClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InstallAgentForClusterResponse) GetBody() *InstallAgentForClusterResponseBody {
	return s.Body
}

func (s *InstallAgentForClusterResponse) SetHeaders(v map[string]*string) *InstallAgentForClusterResponse {
	s.Headers = v
	return s
}

func (s *InstallAgentForClusterResponse) SetStatusCode(v int32) *InstallAgentForClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallAgentForClusterResponse) SetBody(v *InstallAgentForClusterResponseBody) *InstallAgentForClusterResponse {
	s.Body = v
	return s
}

func (s *InstallAgentForClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
