// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallMonitoringAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InstallMonitoringAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InstallMonitoringAgentResponse
	GetStatusCode() *int32
	SetBody(v *InstallMonitoringAgentResponseBody) *InstallMonitoringAgentResponse
	GetBody() *InstallMonitoringAgentResponseBody
}

type InstallMonitoringAgentResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstallMonitoringAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallMonitoringAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s InstallMonitoringAgentResponse) GoString() string {
	return s.String()
}

func (s *InstallMonitoringAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InstallMonitoringAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InstallMonitoringAgentResponse) GetBody() *InstallMonitoringAgentResponseBody {
	return s.Body
}

func (s *InstallMonitoringAgentResponse) SetHeaders(v map[string]*string) *InstallMonitoringAgentResponse {
	s.Headers = v
	return s
}

func (s *InstallMonitoringAgentResponse) SetStatusCode(v int32) *InstallMonitoringAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallMonitoringAgentResponse) SetBody(v *InstallMonitoringAgentResponseBody) *InstallMonitoringAgentResponse {
	s.Body = v
	return s
}

func (s *InstallMonitoringAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
