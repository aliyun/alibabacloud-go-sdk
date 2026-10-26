// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSteerAgentConversationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SteerAgentConversationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SteerAgentConversationResponse
	GetStatusCode() *int32
	SetBody(v *SteerAgentConversationResponseBody) *SteerAgentConversationResponse
	GetBody() *SteerAgentConversationResponseBody
}

type SteerAgentConversationResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SteerAgentConversationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SteerAgentConversationResponse) String() string {
	return dara.Prettify(s)
}

func (s SteerAgentConversationResponse) GoString() string {
	return s.String()
}

func (s *SteerAgentConversationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SteerAgentConversationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SteerAgentConversationResponse) GetBody() *SteerAgentConversationResponseBody {
	return s.Body
}

func (s *SteerAgentConversationResponse) SetHeaders(v map[string]*string) *SteerAgentConversationResponse {
	s.Headers = v
	return s
}

func (s *SteerAgentConversationResponse) SetStatusCode(v int32) *SteerAgentConversationResponse {
	s.StatusCode = &v
	return s
}

func (s *SteerAgentConversationResponse) SetBody(v *SteerAgentConversationResponseBody) *SteerAgentConversationResponse {
	s.Body = v
	return s
}

func (s *SteerAgentConversationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
