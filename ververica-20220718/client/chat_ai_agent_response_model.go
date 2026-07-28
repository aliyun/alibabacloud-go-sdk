// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatAiAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChatAiAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChatAiAgentResponse
	GetStatusCode() *int32
	SetId(v string) *ChatAiAgentResponse
	GetId() *string
	SetEvent(v string) *ChatAiAgentResponse
	GetEvent() *string
	SetBody(v *ChatAiAgentResponseBody) *ChatAiAgentResponse
	GetBody() *ChatAiAgentResponseBody
}

type ChatAiAgentResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                  `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                  `json:"event,omitempty" xml:"event,omitempty"`
	Body       *ChatAiAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChatAiAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s ChatAiAgentResponse) GoString() string {
	return s.String()
}

func (s *ChatAiAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChatAiAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChatAiAgentResponse) GetId() *string {
	return s.Id
}

func (s *ChatAiAgentResponse) GetEvent() *string {
	return s.Event
}

func (s *ChatAiAgentResponse) GetBody() *ChatAiAgentResponseBody {
	return s.Body
}

func (s *ChatAiAgentResponse) SetHeaders(v map[string]*string) *ChatAiAgentResponse {
	s.Headers = v
	return s
}

func (s *ChatAiAgentResponse) SetStatusCode(v int32) *ChatAiAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatAiAgentResponse) SetId(v string) *ChatAiAgentResponse {
	s.Id = &v
	return s
}

func (s *ChatAiAgentResponse) SetEvent(v string) *ChatAiAgentResponse {
	s.Event = &v
	return s
}

func (s *ChatAiAgentResponse) SetBody(v *ChatAiAgentResponseBody) *ChatAiAgentResponse {
	s.Body = v
	return s
}

func (s *ChatAiAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
