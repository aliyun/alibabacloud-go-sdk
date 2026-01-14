// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgentCancelCallResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AgentCancelCallResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AgentCancelCallResponse
	GetStatusCode() *int32
	SetBody(v *AgentCancelCallResponseBody) *AgentCancelCallResponse
	GetBody() *AgentCancelCallResponseBody
}

type AgentCancelCallResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AgentCancelCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AgentCancelCallResponse) String() string {
	return dara.Prettify(s)
}

func (s AgentCancelCallResponse) GoString() string {
	return s.String()
}

func (s *AgentCancelCallResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AgentCancelCallResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AgentCancelCallResponse) GetBody() *AgentCancelCallResponseBody {
	return s.Body
}

func (s *AgentCancelCallResponse) SetHeaders(v map[string]*string) *AgentCancelCallResponse {
	s.Headers = v
	return s
}

func (s *AgentCancelCallResponse) SetStatusCode(v int32) *AgentCancelCallResponse {
	s.StatusCode = &v
	return s
}

func (s *AgentCancelCallResponse) SetBody(v *AgentCancelCallResponseBody) *AgentCancelCallResponse {
	s.Body = v
	return s
}

func (s *AgentCancelCallResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
