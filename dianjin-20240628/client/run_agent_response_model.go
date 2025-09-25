// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunAgentResponse
	GetStatusCode() *int32
	SetBody(v *RunAgentResponseBody) *RunAgentResponse
	GetBody() *RunAgentResponseBody
}

type RunAgentResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s RunAgentResponse) GoString() string {
	return s.String()
}

func (s *RunAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunAgentResponse) GetBody() *RunAgentResponseBody {
	return s.Body
}

func (s *RunAgentResponse) SetHeaders(v map[string]*string) *RunAgentResponse {
	s.Headers = v
	return s
}

func (s *RunAgentResponse) SetStatusCode(v int32) *RunAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *RunAgentResponse) SetBody(v *RunAgentResponseBody) *RunAgentResponse {
	s.Body = v
	return s
}

func (s *RunAgentResponse) Validate() error {
	return dara.Validate(s)
}
