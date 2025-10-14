// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGroupMonitoringAgentProcessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteGroupMonitoringAgentProcessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteGroupMonitoringAgentProcessResponse
	GetStatusCode() *int32
	SetBody(v *DeleteGroupMonitoringAgentProcessResponseBody) *DeleteGroupMonitoringAgentProcessResponse
	GetBody() *DeleteGroupMonitoringAgentProcessResponseBody
}

type DeleteGroupMonitoringAgentProcessResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGroupMonitoringAgentProcessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGroupMonitoringAgentProcessResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteGroupMonitoringAgentProcessResponse) GoString() string {
	return s.String()
}

func (s *DeleteGroupMonitoringAgentProcessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteGroupMonitoringAgentProcessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteGroupMonitoringAgentProcessResponse) GetBody() *DeleteGroupMonitoringAgentProcessResponseBody {
	return s.Body
}

func (s *DeleteGroupMonitoringAgentProcessResponse) SetHeaders(v map[string]*string) *DeleteGroupMonitoringAgentProcessResponse {
	s.Headers = v
	return s
}

func (s *DeleteGroupMonitoringAgentProcessResponse) SetStatusCode(v int32) *DeleteGroupMonitoringAgentProcessResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGroupMonitoringAgentProcessResponse) SetBody(v *DeleteGroupMonitoringAgentProcessResponseBody) *DeleteGroupMonitoringAgentProcessResponse {
	s.Body = v
	return s
}

func (s *DeleteGroupMonitoringAgentProcessResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
