// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMonitoringAgentProcessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMonitoringAgentProcessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMonitoringAgentProcessResponse
	GetStatusCode() *int32
	SetBody(v *CreateMonitoringAgentProcessResponseBody) *CreateMonitoringAgentProcessResponse
	GetBody() *CreateMonitoringAgentProcessResponseBody
}

type CreateMonitoringAgentProcessResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMonitoringAgentProcessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMonitoringAgentProcessResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMonitoringAgentProcessResponse) GoString() string {
	return s.String()
}

func (s *CreateMonitoringAgentProcessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMonitoringAgentProcessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMonitoringAgentProcessResponse) GetBody() *CreateMonitoringAgentProcessResponseBody {
	return s.Body
}

func (s *CreateMonitoringAgentProcessResponse) SetHeaders(v map[string]*string) *CreateMonitoringAgentProcessResponse {
	s.Headers = v
	return s
}

func (s *CreateMonitoringAgentProcessResponse) SetStatusCode(v int32) *CreateMonitoringAgentProcessResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMonitoringAgentProcessResponse) SetBody(v *CreateMonitoringAgentProcessResponseBody) *CreateMonitoringAgentProcessResponse {
	s.Body = v
	return s
}

func (s *CreateMonitoringAgentProcessResponse) Validate() error {
	return dara.Validate(s)
}
