// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReconnectAppChatResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReconnectAppChatResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReconnectAppChatResponse
	GetStatusCode() *int32
	SetId(v string) *ReconnectAppChatResponse
	GetId() *string
	SetEvent(v string) *ReconnectAppChatResponse
	GetEvent() *string
	SetBody(v string) *ReconnectAppChatResponse
	GetBody() *string
}

type ReconnectAppChatResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string            `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string            `json:"event,omitempty" xml:"event,omitempty"`
	Body       *string            `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReconnectAppChatResponse) String() string {
	return dara.Prettify(s)
}

func (s ReconnectAppChatResponse) GoString() string {
	return s.String()
}

func (s *ReconnectAppChatResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReconnectAppChatResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReconnectAppChatResponse) GetId() *string {
	return s.Id
}

func (s *ReconnectAppChatResponse) GetEvent() *string {
	return s.Event
}

func (s *ReconnectAppChatResponse) GetBody() *string {
	return s.Body
}

func (s *ReconnectAppChatResponse) SetHeaders(v map[string]*string) *ReconnectAppChatResponse {
	s.Headers = v
	return s
}

func (s *ReconnectAppChatResponse) SetStatusCode(v int32) *ReconnectAppChatResponse {
	s.StatusCode = &v
	return s
}

func (s *ReconnectAppChatResponse) SetId(v string) *ReconnectAppChatResponse {
	s.Id = &v
	return s
}

func (s *ReconnectAppChatResponse) SetEvent(v string) *ReconnectAppChatResponse {
	s.Event = &v
	return s
}

func (s *ReconnectAppChatResponse) SetBody(v string) *ReconnectAppChatResponse {
	s.Body = &v
	return s
}

func (s *ReconnectAppChatResponse) Validate() error {
	return dara.Validate(s)
}
