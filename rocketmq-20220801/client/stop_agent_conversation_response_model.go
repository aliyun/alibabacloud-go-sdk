// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopAgentConversationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopAgentConversationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopAgentConversationResponse
	GetStatusCode() *int32
	SetBody(v *StopAgentConversationResponseBody) *StopAgentConversationResponse
	GetBody() *StopAgentConversationResponseBody
}

type StopAgentConversationResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopAgentConversationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopAgentConversationResponse) String() string {
	return dara.Prettify(s)
}

func (s StopAgentConversationResponse) GoString() string {
	return s.String()
}

func (s *StopAgentConversationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopAgentConversationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopAgentConversationResponse) GetBody() *StopAgentConversationResponseBody {
	return s.Body
}

func (s *StopAgentConversationResponse) SetHeaders(v map[string]*string) *StopAgentConversationResponse {
	s.Headers = v
	return s
}

func (s *StopAgentConversationResponse) SetStatusCode(v int32) *StopAgentConversationResponse {
	s.StatusCode = &v
	return s
}

func (s *StopAgentConversationResponse) SetBody(v *StopAgentConversationResponseBody) *StopAgentConversationResponse {
	s.Body = v
	return s
}

func (s *StopAgentConversationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
