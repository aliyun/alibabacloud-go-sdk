// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKopilotListConversationChatMessagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *KopilotListConversationChatMessagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *KopilotListConversationChatMessagesResponse
	GetStatusCode() *int32
	SetBody(v *KopilotListConversationChatMessagesResponseBody) *KopilotListConversationChatMessagesResponse
	GetBody() *KopilotListConversationChatMessagesResponseBody
}

type KopilotListConversationChatMessagesResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *KopilotListConversationChatMessagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s KopilotListConversationChatMessagesResponse) String() string {
	return dara.Prettify(s)
}

func (s KopilotListConversationChatMessagesResponse) GoString() string {
	return s.String()
}

func (s *KopilotListConversationChatMessagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *KopilotListConversationChatMessagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *KopilotListConversationChatMessagesResponse) GetBody() *KopilotListConversationChatMessagesResponseBody {
	return s.Body
}

func (s *KopilotListConversationChatMessagesResponse) SetHeaders(v map[string]*string) *KopilotListConversationChatMessagesResponse {
	s.Headers = v
	return s
}

func (s *KopilotListConversationChatMessagesResponse) SetStatusCode(v int32) *KopilotListConversationChatMessagesResponse {
	s.StatusCode = &v
	return s
}

func (s *KopilotListConversationChatMessagesResponse) SetBody(v *KopilotListConversationChatMessagesResponseBody) *KopilotListConversationChatMessagesResponse {
	s.Body = v
	return s
}

func (s *KopilotListConversationChatMessagesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
