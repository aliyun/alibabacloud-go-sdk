// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomAgentMonitorSessionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCustomAgentMonitorSessionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCustomAgentMonitorSessionsResponse
	GetStatusCode() *int32
	SetBody(v *ListCustomAgentMonitorSessionsResponseBody) *ListCustomAgentMonitorSessionsResponse
	GetBody() *ListCustomAgentMonitorSessionsResponseBody
}

type ListCustomAgentMonitorSessionsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCustomAgentMonitorSessionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCustomAgentMonitorSessionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCustomAgentMonitorSessionsResponse) GoString() string {
	return s.String()
}

func (s *ListCustomAgentMonitorSessionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCustomAgentMonitorSessionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCustomAgentMonitorSessionsResponse) GetBody() *ListCustomAgentMonitorSessionsResponseBody {
	return s.Body
}

func (s *ListCustomAgentMonitorSessionsResponse) SetHeaders(v map[string]*string) *ListCustomAgentMonitorSessionsResponse {
	s.Headers = v
	return s
}

func (s *ListCustomAgentMonitorSessionsResponse) SetStatusCode(v int32) *ListCustomAgentMonitorSessionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCustomAgentMonitorSessionsResponse) SetBody(v *ListCustomAgentMonitorSessionsResponseBody) *ListCustomAgentMonitorSessionsResponse {
	s.Body = v
	return s
}

func (s *ListCustomAgentMonitorSessionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
