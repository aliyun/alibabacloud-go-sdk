// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgentCardResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AgentCardResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AgentCardResponse
	GetStatusCode() *int32
	SetBody(v interface{}) *AgentCardResponse
	GetBody() interface{}
}

type AgentCardResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       interface{}        `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AgentCardResponse) String() string {
	return dara.Prettify(s)
}

func (s AgentCardResponse) GoString() string {
	return s.String()
}

func (s *AgentCardResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AgentCardResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AgentCardResponse) GetBody() interface{} {
	return s.Body
}

func (s *AgentCardResponse) SetHeaders(v map[string]*string) *AgentCardResponse {
	s.Headers = v
	return s
}

func (s *AgentCardResponse) SetStatusCode(v int32) *AgentCardResponse {
	s.StatusCode = &v
	return s
}

func (s *AgentCardResponse) SetBody(v interface{}) *AgentCardResponse {
	s.Body = v
	return s
}

func (s *AgentCardResponse) Validate() error {
	return dara.Validate(s)
}
