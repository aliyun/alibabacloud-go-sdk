// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendAgentConversationMessageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendAgentConversationMessageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendAgentConversationMessageResponse
	GetStatusCode() *int32
	SetId(v string) *SendAgentConversationMessageResponse
	GetId() *string
	SetEvent(v string) *SendAgentConversationMessageResponse
	GetEvent() *string
	SetBody(v string) *SendAgentConversationMessageResponse
	GetBody() *string
}

type SendAgentConversationMessageResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string            `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string            `json:"event,omitempty" xml:"event,omitempty"`
	Body       *string            `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendAgentConversationMessageResponse) String() string {
	return dara.Prettify(s)
}

func (s SendAgentConversationMessageResponse) GoString() string {
	return s.String()
}

func (s *SendAgentConversationMessageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendAgentConversationMessageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendAgentConversationMessageResponse) GetId() *string {
	return s.Id
}

func (s *SendAgentConversationMessageResponse) GetEvent() *string {
	return s.Event
}

func (s *SendAgentConversationMessageResponse) GetBody() *string {
	return s.Body
}

func (s *SendAgentConversationMessageResponse) SetHeaders(v map[string]*string) *SendAgentConversationMessageResponse {
	s.Headers = v
	return s
}

func (s *SendAgentConversationMessageResponse) SetStatusCode(v int32) *SendAgentConversationMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *SendAgentConversationMessageResponse) SetId(v string) *SendAgentConversationMessageResponse {
	s.Id = &v
	return s
}

func (s *SendAgentConversationMessageResponse) SetEvent(v string) *SendAgentConversationMessageResponse {
	s.Event = &v
	return s
}

func (s *SendAgentConversationMessageResponse) SetBody(v string) *SendAgentConversationMessageResponse {
	s.Body = &v
	return s
}

func (s *SendAgentConversationMessageResponse) Validate() error {
	return dara.Validate(s)
}
