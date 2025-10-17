// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAgentResponse
	GetStatusCode() *int32
	SetBody(v *ListAgentResponseBody) *ListAgentResponse
	GetBody() *ListAgentResponseBody
}

type ListAgentResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAgentResponse) GoString() string {
	return s.String()
}

func (s *ListAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAgentResponse) GetBody() *ListAgentResponseBody {
	return s.Body
}

func (s *ListAgentResponse) SetHeaders(v map[string]*string) *ListAgentResponse {
	s.Headers = v
	return s
}

func (s *ListAgentResponse) SetStatusCode(v int32) *ListAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAgentResponse) SetBody(v *ListAgentResponseBody) *ListAgentResponse {
	s.Body = v
	return s
}

func (s *ListAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
