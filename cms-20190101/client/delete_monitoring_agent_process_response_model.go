// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMonitoringAgentProcessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMonitoringAgentProcessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMonitoringAgentProcessResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMonitoringAgentProcessResponseBody) *DeleteMonitoringAgentProcessResponse
	GetBody() *DeleteMonitoringAgentProcessResponseBody
}

type DeleteMonitoringAgentProcessResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMonitoringAgentProcessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMonitoringAgentProcessResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMonitoringAgentProcessResponse) GoString() string {
	return s.String()
}

func (s *DeleteMonitoringAgentProcessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMonitoringAgentProcessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMonitoringAgentProcessResponse) GetBody() *DeleteMonitoringAgentProcessResponseBody {
	return s.Body
}

func (s *DeleteMonitoringAgentProcessResponse) SetHeaders(v map[string]*string) *DeleteMonitoringAgentProcessResponse {
	s.Headers = v
	return s
}

func (s *DeleteMonitoringAgentProcessResponse) SetStatusCode(v int32) *DeleteMonitoringAgentProcessResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMonitoringAgentProcessResponse) SetBody(v *DeleteMonitoringAgentProcessResponseBody) *DeleteMonitoringAgentProcessResponse {
	s.Body = v
	return s
}

func (s *DeleteMonitoringAgentProcessResponse) Validate() error {
	return dara.Validate(s)
}
