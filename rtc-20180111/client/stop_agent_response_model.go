// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopAgentResponse
	GetStatusCode() *int32
	SetBody(v *StopAgentResponseBody) *StopAgentResponse
	GetBody() *StopAgentResponseBody
}

type StopAgentResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s StopAgentResponse) GoString() string {
	return s.String()
}

func (s *StopAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopAgentResponse) GetBody() *StopAgentResponseBody {
	return s.Body
}

func (s *StopAgentResponse) SetHeaders(v map[string]*string) *StopAgentResponse {
	s.Headers = v
	return s
}

func (s *StopAgentResponse) SetStatusCode(v int32) *StopAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *StopAgentResponse) SetBody(v *StopAgentResponseBody) *StopAgentResponse {
	s.Body = v
	return s
}

func (s *StopAgentResponse) Validate() error {
	return dara.Validate(s)
}
