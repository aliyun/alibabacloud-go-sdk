// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataAgentMcpResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataAgentMcpResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataAgentMcpResponse
	GetStatusCode() *int32
	SetBody(v *ListDataAgentMcpResponseBody) *ListDataAgentMcpResponse
	GetBody() *ListDataAgentMcpResponseBody
}

type ListDataAgentMcpResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataAgentMcpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataAgentMcpResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataAgentMcpResponse) GoString() string {
	return s.String()
}

func (s *ListDataAgentMcpResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataAgentMcpResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataAgentMcpResponse) GetBody() *ListDataAgentMcpResponseBody {
	return s.Body
}

func (s *ListDataAgentMcpResponse) SetHeaders(v map[string]*string) *ListDataAgentMcpResponse {
	s.Headers = v
	return s
}

func (s *ListDataAgentMcpResponse) SetStatusCode(v int32) *ListDataAgentMcpResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataAgentMcpResponse) SetBody(v *ListDataAgentMcpResponseBody) *ListDataAgentMcpResponse {
	s.Body = v
	return s
}

func (s *ListDataAgentMcpResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
