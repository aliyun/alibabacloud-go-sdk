// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentRuntimeEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAgentRuntimeEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAgentRuntimeEndpointResponse
	GetStatusCode() *int32
	SetBody(v *AgentRuntimeEndpointResult) *DeleteAgentRuntimeEndpointResponse
	GetBody() *AgentRuntimeEndpointResult
}

type DeleteAgentRuntimeEndpointResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AgentRuntimeEndpointResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAgentRuntimeEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentRuntimeEndpointResponse) GoString() string {
	return s.String()
}

func (s *DeleteAgentRuntimeEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAgentRuntimeEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAgentRuntimeEndpointResponse) GetBody() *AgentRuntimeEndpointResult {
	return s.Body
}

func (s *DeleteAgentRuntimeEndpointResponse) SetHeaders(v map[string]*string) *DeleteAgentRuntimeEndpointResponse {
	s.Headers = v
	return s
}

func (s *DeleteAgentRuntimeEndpointResponse) SetStatusCode(v int32) *DeleteAgentRuntimeEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAgentRuntimeEndpointResponse) SetBody(v *AgentRuntimeEndpointResult) *DeleteAgentRuntimeEndpointResponse {
	s.Body = v
	return s
}

func (s *DeleteAgentRuntimeEndpointResponse) Validate() error {
	return dara.Validate(s)
}
