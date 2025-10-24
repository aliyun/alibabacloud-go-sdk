// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallMonitorAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UninstallMonitorAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UninstallMonitorAgentResponse
	GetStatusCode() *int32
	SetBody(v *UninstallMonitorAgentResponseBody) *UninstallMonitorAgentResponse
	GetBody() *UninstallMonitorAgentResponseBody
}

type UninstallMonitorAgentResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UninstallMonitorAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UninstallMonitorAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s UninstallMonitorAgentResponse) GoString() string {
	return s.String()
}

func (s *UninstallMonitorAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UninstallMonitorAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UninstallMonitorAgentResponse) GetBody() *UninstallMonitorAgentResponseBody {
	return s.Body
}

func (s *UninstallMonitorAgentResponse) SetHeaders(v map[string]*string) *UninstallMonitorAgentResponse {
	s.Headers = v
	return s
}

func (s *UninstallMonitorAgentResponse) SetStatusCode(v int32) *UninstallMonitorAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *UninstallMonitorAgentResponse) SetBody(v *UninstallMonitorAgentResponseBody) *UninstallMonitorAgentResponse {
	s.Body = v
	return s
}

func (s *UninstallMonitorAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
