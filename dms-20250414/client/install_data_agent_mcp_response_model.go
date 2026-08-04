// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallDataAgentMcpResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InstallDataAgentMcpResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InstallDataAgentMcpResponse
	GetStatusCode() *int32
	SetBody(v *InstallDataAgentMcpResponseBody) *InstallDataAgentMcpResponse
	GetBody() *InstallDataAgentMcpResponseBody
}

type InstallDataAgentMcpResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstallDataAgentMcpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallDataAgentMcpResponse) String() string {
	return dara.Prettify(s)
}

func (s InstallDataAgentMcpResponse) GoString() string {
	return s.String()
}

func (s *InstallDataAgentMcpResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InstallDataAgentMcpResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InstallDataAgentMcpResponse) GetBody() *InstallDataAgentMcpResponseBody {
	return s.Body
}

func (s *InstallDataAgentMcpResponse) SetHeaders(v map[string]*string) *InstallDataAgentMcpResponse {
	s.Headers = v
	return s
}

func (s *InstallDataAgentMcpResponse) SetStatusCode(v int32) *InstallDataAgentMcpResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallDataAgentMcpResponse) SetBody(v *InstallDataAgentMcpResponseBody) *InstallDataAgentMcpResponse {
	s.Body = v
	return s
}

func (s *InstallDataAgentMcpResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
