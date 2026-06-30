// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAgentTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAgentTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateAgentTaskResponseBody) *CreateAgentTaskResponse
	GetBody() *CreateAgentTaskResponseBody
}

type CreateAgentTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAgentTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAgentTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateAgentTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAgentTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAgentTaskResponse) GetBody() *CreateAgentTaskResponseBody {
	return s.Body
}

func (s *CreateAgentTaskResponse) SetHeaders(v map[string]*string) *CreateAgentTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateAgentTaskResponse) SetStatusCode(v int32) *CreateAgentTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAgentTaskResponse) SetBody(v *CreateAgentTaskResponseBody) *CreateAgentTaskResponse {
	s.Body = v
	return s
}

func (s *CreateAgentTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
