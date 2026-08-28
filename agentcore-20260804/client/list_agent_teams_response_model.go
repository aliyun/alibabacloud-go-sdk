// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentTeamsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAgentTeamsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAgentTeamsResponse
	GetStatusCode() *int32
	SetBody(v *ListAgentTeamsResponseBody) *ListAgentTeamsResponse
	GetBody() *ListAgentTeamsResponseBody
}

type ListAgentTeamsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAgentTeamsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAgentTeamsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAgentTeamsResponse) GoString() string {
	return s.String()
}

func (s *ListAgentTeamsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAgentTeamsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAgentTeamsResponse) GetBody() *ListAgentTeamsResponseBody {
	return s.Body
}

func (s *ListAgentTeamsResponse) SetHeaders(v map[string]*string) *ListAgentTeamsResponse {
	s.Headers = v
	return s
}

func (s *ListAgentTeamsResponse) SetStatusCode(v int32) *ListAgentTeamsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAgentTeamsResponse) SetBody(v *ListAgentTeamsResponseBody) *ListAgentTeamsResponse {
	s.Body = v
	return s
}

func (s *ListAgentTeamsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
