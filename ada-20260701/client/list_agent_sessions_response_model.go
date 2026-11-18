// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentSessionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAgentSessionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAgentSessionsResponse
	GetStatusCode() *int32
	SetBody(v *ListAgentSessionsResponseBody) *ListAgentSessionsResponse
	GetBody() *ListAgentSessionsResponseBody
}

type ListAgentSessionsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAgentSessionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAgentSessionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAgentSessionsResponse) GoString() string {
	return s.String()
}

func (s *ListAgentSessionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAgentSessionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAgentSessionsResponse) GetBody() *ListAgentSessionsResponseBody {
	return s.Body
}

func (s *ListAgentSessionsResponse) SetHeaders(v map[string]*string) *ListAgentSessionsResponse {
	s.Headers = v
	return s
}

func (s *ListAgentSessionsResponse) SetStatusCode(v int32) *ListAgentSessionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAgentSessionsResponse) SetBody(v *ListAgentSessionsResponseBody) *ListAgentSessionsResponse {
	s.Body = v
	return s
}

func (s *ListAgentSessionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
