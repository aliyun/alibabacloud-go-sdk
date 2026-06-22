// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKopilotChatStreamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *KopilotChatStreamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *KopilotChatStreamResponse
	GetStatusCode() *int32
	SetId(v string) *KopilotChatStreamResponse
	GetId() *string
	SetEvent(v string) *KopilotChatStreamResponse
	GetEvent() *string
	SetBody(v *KopilotChatStreamResponseBody) *KopilotChatStreamResponse
	GetBody() *KopilotChatStreamResponseBody
}

type KopilotChatStreamResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                        `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                        `json:"event,omitempty" xml:"event,omitempty"`
	Body       *KopilotChatStreamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s KopilotChatStreamResponse) String() string {
	return dara.Prettify(s)
}

func (s KopilotChatStreamResponse) GoString() string {
	return s.String()
}

func (s *KopilotChatStreamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *KopilotChatStreamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *KopilotChatStreamResponse) GetId() *string {
	return s.Id
}

func (s *KopilotChatStreamResponse) GetEvent() *string {
	return s.Event
}

func (s *KopilotChatStreamResponse) GetBody() *KopilotChatStreamResponseBody {
	return s.Body
}

func (s *KopilotChatStreamResponse) SetHeaders(v map[string]*string) *KopilotChatStreamResponse {
	s.Headers = v
	return s
}

func (s *KopilotChatStreamResponse) SetStatusCode(v int32) *KopilotChatStreamResponse {
	s.StatusCode = &v
	return s
}

func (s *KopilotChatStreamResponse) SetId(v string) *KopilotChatStreamResponse {
	s.Id = &v
	return s
}

func (s *KopilotChatStreamResponse) SetEvent(v string) *KopilotChatStreamResponse {
	s.Event = &v
	return s
}

func (s *KopilotChatStreamResponse) SetBody(v *KopilotChatStreamResponseBody) *KopilotChatStreamResponse {
	s.Body = v
	return s
}

func (s *KopilotChatStreamResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
