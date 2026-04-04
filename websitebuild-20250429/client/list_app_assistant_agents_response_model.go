// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppAssistantAgentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAppAssistantAgentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAppAssistantAgentsResponse
	GetStatusCode() *int32
	SetBody(v *ListAppAssistantAgentsResponseBody) *ListAppAssistantAgentsResponse
	GetBody() *ListAppAssistantAgentsResponseBody
}

type ListAppAssistantAgentsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppAssistantAgentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAppAssistantAgentsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAppAssistantAgentsResponse) GoString() string {
	return s.String()
}

func (s *ListAppAssistantAgentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAppAssistantAgentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAppAssistantAgentsResponse) GetBody() *ListAppAssistantAgentsResponseBody {
	return s.Body
}

func (s *ListAppAssistantAgentsResponse) SetHeaders(v map[string]*string) *ListAppAssistantAgentsResponse {
	s.Headers = v
	return s
}

func (s *ListAppAssistantAgentsResponse) SetStatusCode(v int32) *ListAppAssistantAgentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppAssistantAgentsResponse) SetBody(v *ListAppAssistantAgentsResponseBody) *ListAppAssistantAgentsResponse {
	s.Body = v
	return s
}

func (s *ListAppAssistantAgentsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
