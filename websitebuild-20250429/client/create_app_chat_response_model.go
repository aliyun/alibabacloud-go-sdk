// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppChatResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAppChatResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAppChatResponse
	GetStatusCode() *int32
	SetId(v string) *CreateAppChatResponse
	GetId() *string
	SetEvent(v string) *CreateAppChatResponse
	GetEvent() *string
	SetBody(v string) *CreateAppChatResponse
	GetBody() *string
}

type CreateAppChatResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string            `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string            `json:"event,omitempty" xml:"event,omitempty"`
	Body       *string            `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAppChatResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAppChatResponse) GoString() string {
	return s.String()
}

func (s *CreateAppChatResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAppChatResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAppChatResponse) GetId() *string {
	return s.Id
}

func (s *CreateAppChatResponse) GetEvent() *string {
	return s.Event
}

func (s *CreateAppChatResponse) GetBody() *string {
	return s.Body
}

func (s *CreateAppChatResponse) SetHeaders(v map[string]*string) *CreateAppChatResponse {
	s.Headers = v
	return s
}

func (s *CreateAppChatResponse) SetStatusCode(v int32) *CreateAppChatResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppChatResponse) SetId(v string) *CreateAppChatResponse {
	s.Id = &v
	return s
}

func (s *CreateAppChatResponse) SetEvent(v string) *CreateAppChatResponse {
	s.Event = &v
	return s
}

func (s *CreateAppChatResponse) SetBody(v string) *CreateAppChatResponse {
	s.Body = &v
	return s
}

func (s *CreateAppChatResponse) Validate() error {
	return dara.Validate(s)
}
