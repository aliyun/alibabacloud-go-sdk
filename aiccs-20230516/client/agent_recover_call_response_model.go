// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgentRecoverCallResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AgentRecoverCallResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AgentRecoverCallResponse
	GetStatusCode() *int32
	SetBody(v *AgentRecoverCallResponseBody) *AgentRecoverCallResponse
	GetBody() *AgentRecoverCallResponseBody
}

type AgentRecoverCallResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AgentRecoverCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AgentRecoverCallResponse) String() string {
	return dara.Prettify(s)
}

func (s AgentRecoverCallResponse) GoString() string {
	return s.String()
}

func (s *AgentRecoverCallResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AgentRecoverCallResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AgentRecoverCallResponse) GetBody() *AgentRecoverCallResponseBody {
	return s.Body
}

func (s *AgentRecoverCallResponse) SetHeaders(v map[string]*string) *AgentRecoverCallResponse {
	s.Headers = v
	return s
}

func (s *AgentRecoverCallResponse) SetStatusCode(v int32) *AgentRecoverCallResponse {
	s.StatusCode = &v
	return s
}

func (s *AgentRecoverCallResponse) SetBody(v *AgentRecoverCallResponseBody) *AgentRecoverCallResponse {
	s.Body = v
	return s
}

func (s *AgentRecoverCallResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
