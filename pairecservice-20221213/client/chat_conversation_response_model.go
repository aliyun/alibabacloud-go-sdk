// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatConversationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChatConversationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChatConversationResponse
	GetStatusCode() *int32
	SetId(v string) *ChatConversationResponse
	GetId() *string
	SetEvent(v string) *ChatConversationResponse
	GetEvent() *string
	SetBody(v *ChatConversationResponseBody) *ChatConversationResponse
	GetBody() *ChatConversationResponseBody
}

type ChatConversationResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                       `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                       `json:"event,omitempty" xml:"event,omitempty"`
	Body       *ChatConversationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChatConversationResponse) String() string {
	return dara.Prettify(s)
}

func (s ChatConversationResponse) GoString() string {
	return s.String()
}

func (s *ChatConversationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChatConversationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChatConversationResponse) GetId() *string {
	return s.Id
}

func (s *ChatConversationResponse) GetEvent() *string {
	return s.Event
}

func (s *ChatConversationResponse) GetBody() *ChatConversationResponseBody {
	return s.Body
}

func (s *ChatConversationResponse) SetHeaders(v map[string]*string) *ChatConversationResponse {
	s.Headers = v
	return s
}

func (s *ChatConversationResponse) SetStatusCode(v int32) *ChatConversationResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatConversationResponse) SetId(v string) *ChatConversationResponse {
	s.Id = &v
	return s
}

func (s *ChatConversationResponse) SetEvent(v string) *ChatConversationResponse {
	s.Event = &v
	return s
}

func (s *ChatConversationResponse) SetBody(v *ChatConversationResponseBody) *ChatConversationResponse {
	s.Body = v
	return s
}

func (s *ChatConversationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
