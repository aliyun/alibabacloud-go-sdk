// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgentGroupPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AgentGroupPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AgentGroupPageResponse
	GetStatusCode() *int32
	SetBody(v *AgentGroupPageResponseBody) *AgentGroupPageResponse
	GetBody() *AgentGroupPageResponseBody
}

type AgentGroupPageResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AgentGroupPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AgentGroupPageResponse) String() string {
	return dara.Prettify(s)
}

func (s AgentGroupPageResponse) GoString() string {
	return s.String()
}

func (s *AgentGroupPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AgentGroupPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AgentGroupPageResponse) GetBody() *AgentGroupPageResponseBody {
	return s.Body
}

func (s *AgentGroupPageResponse) SetHeaders(v map[string]*string) *AgentGroupPageResponse {
	s.Headers = v
	return s
}

func (s *AgentGroupPageResponse) SetStatusCode(v int32) *AgentGroupPageResponse {
	s.StatusCode = &v
	return s
}

func (s *AgentGroupPageResponse) SetBody(v *AgentGroupPageResponseBody) *AgentGroupPageResponse {
	s.Body = v
	return s
}

func (s *AgentGroupPageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
