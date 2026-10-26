// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentConversationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAgentConversationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAgentConversationsResponse
	GetStatusCode() *int32
	SetBody(v *ListAgentConversationsResponseBody) *ListAgentConversationsResponse
	GetBody() *ListAgentConversationsResponseBody
}

type ListAgentConversationsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAgentConversationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAgentConversationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAgentConversationsResponse) GoString() string {
	return s.String()
}

func (s *ListAgentConversationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAgentConversationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAgentConversationsResponse) GetBody() *ListAgentConversationsResponseBody {
	return s.Body
}

func (s *ListAgentConversationsResponse) SetHeaders(v map[string]*string) *ListAgentConversationsResponse {
	s.Headers = v
	return s
}

func (s *ListAgentConversationsResponse) SetStatusCode(v int32) *ListAgentConversationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAgentConversationsResponse) SetBody(v *ListAgentConversationsResponseBody) *ListAgentConversationsResponse {
	s.Body = v
	return s
}

func (s *ListAgentConversationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
