// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseAgentTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PauseAgentTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PauseAgentTaskResponse
	GetStatusCode() *int32
	SetBody(v *PauseAgentTaskResponseBody) *PauseAgentTaskResponse
	GetBody() *PauseAgentTaskResponseBody
}

type PauseAgentTaskResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PauseAgentTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PauseAgentTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s PauseAgentTaskResponse) GoString() string {
	return s.String()
}

func (s *PauseAgentTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PauseAgentTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PauseAgentTaskResponse) GetBody() *PauseAgentTaskResponseBody {
	return s.Body
}

func (s *PauseAgentTaskResponse) SetHeaders(v map[string]*string) *PauseAgentTaskResponse {
	s.Headers = v
	return s
}

func (s *PauseAgentTaskResponse) SetStatusCode(v int32) *PauseAgentTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *PauseAgentTaskResponse) SetBody(v *PauseAgentTaskResponseBody) *PauseAgentTaskResponse {
	s.Body = v
	return s
}

func (s *PauseAgentTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
