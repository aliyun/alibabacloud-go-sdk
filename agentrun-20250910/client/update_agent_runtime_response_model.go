// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentRuntimeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAgentRuntimeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAgentRuntimeResponse
	GetStatusCode() *int32
	SetBody(v *AgentRuntimeResult) *UpdateAgentRuntimeResponse
	GetBody() *AgentRuntimeResult
}

type UpdateAgentRuntimeResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AgentRuntimeResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAgentRuntimeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentRuntimeResponse) GoString() string {
	return s.String()
}

func (s *UpdateAgentRuntimeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAgentRuntimeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAgentRuntimeResponse) GetBody() *AgentRuntimeResult {
	return s.Body
}

func (s *UpdateAgentRuntimeResponse) SetHeaders(v map[string]*string) *UpdateAgentRuntimeResponse {
	s.Headers = v
	return s
}

func (s *UpdateAgentRuntimeResponse) SetStatusCode(v int32) *UpdateAgentRuntimeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAgentRuntimeResponse) SetBody(v *AgentRuntimeResult) *UpdateAgentRuntimeResponse {
	s.Body = v
	return s
}

func (s *UpdateAgentRuntimeResponse) Validate() error {
	return dara.Validate(s)
}
