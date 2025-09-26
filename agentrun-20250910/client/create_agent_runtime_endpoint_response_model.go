// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentRuntimeEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAgentRuntimeEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAgentRuntimeEndpointResponse
	GetStatusCode() *int32
	SetBody(v *AgentRuntimeEndpointResult) *CreateAgentRuntimeEndpointResponse
	GetBody() *AgentRuntimeEndpointResult
}

type CreateAgentRuntimeEndpointResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AgentRuntimeEndpointResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAgentRuntimeEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentRuntimeEndpointResponse) GoString() string {
	return s.String()
}

func (s *CreateAgentRuntimeEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAgentRuntimeEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAgentRuntimeEndpointResponse) GetBody() *AgentRuntimeEndpointResult {
	return s.Body
}

func (s *CreateAgentRuntimeEndpointResponse) SetHeaders(v map[string]*string) *CreateAgentRuntimeEndpointResponse {
	s.Headers = v
	return s
}

func (s *CreateAgentRuntimeEndpointResponse) SetStatusCode(v int32) *CreateAgentRuntimeEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAgentRuntimeEndpointResponse) SetBody(v *AgentRuntimeEndpointResult) *CreateAgentRuntimeEndpointResponse {
	s.Body = v
	return s
}

func (s *CreateAgentRuntimeEndpointResponse) Validate() error {
	return dara.Validate(s)
}
