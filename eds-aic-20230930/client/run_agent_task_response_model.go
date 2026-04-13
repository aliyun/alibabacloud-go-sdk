// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunAgentTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunAgentTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunAgentTaskResponse
	GetStatusCode() *int32
	SetBody(v *RunAgentTaskResponseBody) *RunAgentTaskResponse
	GetBody() *RunAgentTaskResponseBody
}

type RunAgentTaskResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunAgentTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunAgentTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s RunAgentTaskResponse) GoString() string {
	return s.String()
}

func (s *RunAgentTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunAgentTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunAgentTaskResponse) GetBody() *RunAgentTaskResponseBody {
	return s.Body
}

func (s *RunAgentTaskResponse) SetHeaders(v map[string]*string) *RunAgentTaskResponse {
	s.Headers = v
	return s
}

func (s *RunAgentTaskResponse) SetStatusCode(v int32) *RunAgentTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *RunAgentTaskResponse) SetBody(v *RunAgentTaskResponseBody) *RunAgentTaskResponse {
	s.Body = v
	return s
}

func (s *RunAgentTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
