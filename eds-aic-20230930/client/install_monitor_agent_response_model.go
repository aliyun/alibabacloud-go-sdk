// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallMonitorAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InstallMonitorAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InstallMonitorAgentResponse
	GetStatusCode() *int32
	SetBody(v *InstallMonitorAgentResponseBody) *InstallMonitorAgentResponse
	GetBody() *InstallMonitorAgentResponseBody
}

type InstallMonitorAgentResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstallMonitorAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallMonitorAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s InstallMonitorAgentResponse) GoString() string {
	return s.String()
}

func (s *InstallMonitorAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InstallMonitorAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InstallMonitorAgentResponse) GetBody() *InstallMonitorAgentResponseBody {
	return s.Body
}

func (s *InstallMonitorAgentResponse) SetHeaders(v map[string]*string) *InstallMonitorAgentResponse {
	s.Headers = v
	return s
}

func (s *InstallMonitorAgentResponse) SetStatusCode(v int32) *InstallMonitorAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallMonitorAgentResponse) SetBody(v *InstallMonitorAgentResponseBody) *InstallMonitorAgentResponse {
	s.Body = v
	return s
}

func (s *InstallMonitorAgentResponse) Validate() error {
	return dara.Validate(s)
}
