// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentRuntimeEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgentRuntimeEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgentRuntimeEndpointResponse
	GetStatusCode() *int32
	SetBody(v *AgentRuntimeEndpointResult) *GetAgentRuntimeEndpointResponse
	GetBody() *AgentRuntimeEndpointResult
}

type GetAgentRuntimeEndpointResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AgentRuntimeEndpointResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgentRuntimeEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgentRuntimeEndpointResponse) GoString() string {
	return s.String()
}

func (s *GetAgentRuntimeEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgentRuntimeEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgentRuntimeEndpointResponse) GetBody() *AgentRuntimeEndpointResult {
	return s.Body
}

func (s *GetAgentRuntimeEndpointResponse) SetHeaders(v map[string]*string) *GetAgentRuntimeEndpointResponse {
	s.Headers = v
	return s
}

func (s *GetAgentRuntimeEndpointResponse) SetStatusCode(v int32) *GetAgentRuntimeEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentRuntimeEndpointResponse) SetBody(v *AgentRuntimeEndpointResult) *GetAgentRuntimeEndpointResponse {
	s.Body = v
	return s
}

func (s *GetAgentRuntimeEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
