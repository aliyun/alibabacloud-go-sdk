// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIAgentEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAIAgentEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAIAgentEventResponse
	GetStatusCode() *int32
	SetBody(v *ListAIAgentEventResponseBody) *ListAIAgentEventResponse
	GetBody() *ListAIAgentEventResponseBody
}

type ListAIAgentEventResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAIAgentEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAIAgentEventResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAIAgentEventResponse) GoString() string {
	return s.String()
}

func (s *ListAIAgentEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAIAgentEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAIAgentEventResponse) GetBody() *ListAIAgentEventResponseBody {
	return s.Body
}

func (s *ListAIAgentEventResponse) SetHeaders(v map[string]*string) *ListAIAgentEventResponse {
	s.Headers = v
	return s
}

func (s *ListAIAgentEventResponse) SetStatusCode(v int32) *ListAIAgentEventResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAIAgentEventResponse) SetBody(v *ListAIAgentEventResponseBody) *ListAIAgentEventResponse {
	s.Body = v
	return s
}

func (s *ListAIAgentEventResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
