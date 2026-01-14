// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgentImportNumberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AgentImportNumberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AgentImportNumberResponse
	GetStatusCode() *int32
	SetBody(v *AgentImportNumberResponseBody) *AgentImportNumberResponse
	GetBody() *AgentImportNumberResponseBody
}

type AgentImportNumberResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AgentImportNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AgentImportNumberResponse) String() string {
	return dara.Prettify(s)
}

func (s AgentImportNumberResponse) GoString() string {
	return s.String()
}

func (s *AgentImportNumberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AgentImportNumberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AgentImportNumberResponse) GetBody() *AgentImportNumberResponseBody {
	return s.Body
}

func (s *AgentImportNumberResponse) SetHeaders(v map[string]*string) *AgentImportNumberResponse {
	s.Headers = v
	return s
}

func (s *AgentImportNumberResponse) SetStatusCode(v int32) *AgentImportNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *AgentImportNumberResponse) SetBody(v *AgentImportNumberResponseBody) *AgentImportNumberResponse {
	s.Body = v
	return s
}

func (s *AgentImportNumberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
