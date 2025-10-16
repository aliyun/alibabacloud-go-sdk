// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentRuntimeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAgentRuntimeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAgentRuntimeResponse
	GetStatusCode() *int32
	SetBody(v *AgentRuntimeResult) *CreateAgentRuntimeResponse
	GetBody() *AgentRuntimeResult
}

type CreateAgentRuntimeResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AgentRuntimeResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAgentRuntimeResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentRuntimeResponse) GoString() string {
	return s.String()
}

func (s *CreateAgentRuntimeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAgentRuntimeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAgentRuntimeResponse) GetBody() *AgentRuntimeResult {
	return s.Body
}

func (s *CreateAgentRuntimeResponse) SetHeaders(v map[string]*string) *CreateAgentRuntimeResponse {
	s.Headers = v
	return s
}

func (s *CreateAgentRuntimeResponse) SetStatusCode(v int32) *CreateAgentRuntimeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAgentRuntimeResponse) SetBody(v *AgentRuntimeResult) *CreateAgentRuntimeResponse {
	s.Body = v
	return s
}

func (s *CreateAgentRuntimeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
