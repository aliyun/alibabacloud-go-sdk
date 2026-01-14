// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgentCallListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AgentCallListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AgentCallListResponse
	GetStatusCode() *int32
	SetBody(v *AgentCallListResponseBody) *AgentCallListResponse
	GetBody() *AgentCallListResponseBody
}

type AgentCallListResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AgentCallListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AgentCallListResponse) String() string {
	return dara.Prettify(s)
}

func (s AgentCallListResponse) GoString() string {
	return s.String()
}

func (s *AgentCallListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AgentCallListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AgentCallListResponse) GetBody() *AgentCallListResponseBody {
	return s.Body
}

func (s *AgentCallListResponse) SetHeaders(v map[string]*string) *AgentCallListResponse {
	s.Headers = v
	return s
}

func (s *AgentCallListResponse) SetStatusCode(v int32) *AgentCallListResponse {
	s.StatusCode = &v
	return s
}

func (s *AgentCallListResponse) SetBody(v *AgentCallListResponseBody) *AgentCallListResponse {
	s.Body = v
	return s
}

func (s *AgentCallListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
