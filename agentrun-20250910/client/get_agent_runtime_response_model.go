// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentRuntimeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgentRuntimeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgentRuntimeResponse
	GetStatusCode() *int32
	SetBody(v *AgentRuntimeResult) *GetAgentRuntimeResponse
	GetBody() *AgentRuntimeResult
}

type GetAgentRuntimeResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AgentRuntimeResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgentRuntimeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgentRuntimeResponse) GoString() string {
	return s.String()
}

func (s *GetAgentRuntimeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgentRuntimeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgentRuntimeResponse) GetBody() *AgentRuntimeResult {
	return s.Body
}

func (s *GetAgentRuntimeResponse) SetHeaders(v map[string]*string) *GetAgentRuntimeResponse {
	s.Headers = v
	return s
}

func (s *GetAgentRuntimeResponse) SetStatusCode(v int32) *GetAgentRuntimeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentRuntimeResponse) SetBody(v *AgentRuntimeResult) *GetAgentRuntimeResponse {
	s.Body = v
	return s
}

func (s *GetAgentRuntimeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
