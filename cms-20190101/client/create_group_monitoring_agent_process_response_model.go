// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupMonitoringAgentProcessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateGroupMonitoringAgentProcessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateGroupMonitoringAgentProcessResponse
	GetStatusCode() *int32
	SetBody(v *CreateGroupMonitoringAgentProcessResponseBody) *CreateGroupMonitoringAgentProcessResponse
	GetBody() *CreateGroupMonitoringAgentProcessResponseBody
}

type CreateGroupMonitoringAgentProcessResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGroupMonitoringAgentProcessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGroupMonitoringAgentProcessResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupMonitoringAgentProcessResponse) GoString() string {
	return s.String()
}

func (s *CreateGroupMonitoringAgentProcessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateGroupMonitoringAgentProcessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateGroupMonitoringAgentProcessResponse) GetBody() *CreateGroupMonitoringAgentProcessResponseBody {
	return s.Body
}

func (s *CreateGroupMonitoringAgentProcessResponse) SetHeaders(v map[string]*string) *CreateGroupMonitoringAgentProcessResponse {
	s.Headers = v
	return s
}

func (s *CreateGroupMonitoringAgentProcessResponse) SetStatusCode(v int32) *CreateGroupMonitoringAgentProcessResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGroupMonitoringAgentProcessResponse) SetBody(v *CreateGroupMonitoringAgentProcessResponseBody) *CreateGroupMonitoringAgentProcessResponse {
	s.Body = v
	return s
}

func (s *CreateGroupMonitoringAgentProcessResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
