// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallMonitoringAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UninstallMonitoringAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UninstallMonitoringAgentResponse
	GetStatusCode() *int32
	SetBody(v *UninstallMonitoringAgentResponseBody) *UninstallMonitoringAgentResponse
	GetBody() *UninstallMonitoringAgentResponseBody
}

type UninstallMonitoringAgentResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UninstallMonitoringAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UninstallMonitoringAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s UninstallMonitoringAgentResponse) GoString() string {
	return s.String()
}

func (s *UninstallMonitoringAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UninstallMonitoringAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UninstallMonitoringAgentResponse) GetBody() *UninstallMonitoringAgentResponseBody {
	return s.Body
}

func (s *UninstallMonitoringAgentResponse) SetHeaders(v map[string]*string) *UninstallMonitoringAgentResponse {
	s.Headers = v
	return s
}

func (s *UninstallMonitoringAgentResponse) SetStatusCode(v int32) *UninstallMonitoringAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *UninstallMonitoringAgentResponse) SetBody(v *UninstallMonitoringAgentResponseBody) *UninstallMonitoringAgentResponse {
	s.Body = v
	return s
}

func (s *UninstallMonitoringAgentResponse) Validate() error {
	return dara.Validate(s)
}
