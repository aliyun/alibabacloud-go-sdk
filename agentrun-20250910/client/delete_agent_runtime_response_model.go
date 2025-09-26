// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentRuntimeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAgentRuntimeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAgentRuntimeResponse
	GetStatusCode() *int32
	SetBody(v *AgentRuntimeResult) *DeleteAgentRuntimeResponse
	GetBody() *AgentRuntimeResult
}

type DeleteAgentRuntimeResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AgentRuntimeResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAgentRuntimeResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentRuntimeResponse) GoString() string {
	return s.String()
}

func (s *DeleteAgentRuntimeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAgentRuntimeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAgentRuntimeResponse) GetBody() *AgentRuntimeResult {
	return s.Body
}

func (s *DeleteAgentRuntimeResponse) SetHeaders(v map[string]*string) *DeleteAgentRuntimeResponse {
	s.Headers = v
	return s
}

func (s *DeleteAgentRuntimeResponse) SetStatusCode(v int32) *DeleteAgentRuntimeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAgentRuntimeResponse) SetBody(v *AgentRuntimeResult) *DeleteAgentRuntimeResponse {
	s.Body = v
	return s
}

func (s *DeleteAgentRuntimeResponse) Validate() error {
	return dara.Validate(s)
}
