// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentRuntimeEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAgentRuntimeEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAgentRuntimeEndpointResponse
	GetStatusCode() *int32
	SetBody(v *AgentRuntimeEndpointResult) *UpdateAgentRuntimeEndpointResponse
	GetBody() *AgentRuntimeEndpointResult
}

type UpdateAgentRuntimeEndpointResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AgentRuntimeEndpointResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAgentRuntimeEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentRuntimeEndpointResponse) GoString() string {
	return s.String()
}

func (s *UpdateAgentRuntimeEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAgentRuntimeEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAgentRuntimeEndpointResponse) GetBody() *AgentRuntimeEndpointResult {
	return s.Body
}

func (s *UpdateAgentRuntimeEndpointResponse) SetHeaders(v map[string]*string) *UpdateAgentRuntimeEndpointResponse {
	s.Headers = v
	return s
}

func (s *UpdateAgentRuntimeEndpointResponse) SetStatusCode(v int32) *UpdateAgentRuntimeEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAgentRuntimeEndpointResponse) SetBody(v *AgentRuntimeEndpointResult) *UpdateAgentRuntimeEndpointResponse {
	s.Body = v
	return s
}

func (s *UpdateAgentRuntimeEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
