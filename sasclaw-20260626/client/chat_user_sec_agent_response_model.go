// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatUserSecAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChatUserSecAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChatUserSecAgentResponse
	GetStatusCode() *int32
	SetId(v string) *ChatUserSecAgentResponse
	GetId() *string
	SetEvent(v string) *ChatUserSecAgentResponse
	GetEvent() *string
	SetBody(v string) *ChatUserSecAgentResponse
	GetBody() *string
}

type ChatUserSecAgentResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string            `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string            `json:"event,omitempty" xml:"event,omitempty"`
	Body       *string            `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChatUserSecAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s ChatUserSecAgentResponse) GoString() string {
	return s.String()
}

func (s *ChatUserSecAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChatUserSecAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChatUserSecAgentResponse) GetId() *string {
	return s.Id
}

func (s *ChatUserSecAgentResponse) GetEvent() *string {
	return s.Event
}

func (s *ChatUserSecAgentResponse) GetBody() *string {
	return s.Body
}

func (s *ChatUserSecAgentResponse) SetHeaders(v map[string]*string) *ChatUserSecAgentResponse {
	s.Headers = v
	return s
}

func (s *ChatUserSecAgentResponse) SetStatusCode(v int32) *ChatUserSecAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatUserSecAgentResponse) SetId(v string) *ChatUserSecAgentResponse {
	s.Id = &v
	return s
}

func (s *ChatUserSecAgentResponse) SetEvent(v string) *ChatUserSecAgentResponse {
	s.Event = &v
	return s
}

func (s *ChatUserSecAgentResponse) SetBody(v string) *ChatUserSecAgentResponse {
	s.Body = &v
	return s
}

func (s *ChatUserSecAgentResponse) Validate() error {
	return dara.Validate(s)
}
