// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendMessageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendMessageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendMessageResponse
	GetStatusCode() *int32
	SetId(v string) *SendMessageResponse
	GetId() *string
	SetEvent(v string) *SendMessageResponse
	GetEvent() *string
	SetBody(v string) *SendMessageResponse
	GetBody() *string
}

type SendMessageResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string            `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string            `json:"event,omitempty" xml:"event,omitempty"`
	Body       *string            `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendMessageResponse) String() string {
	return dara.Prettify(s)
}

func (s SendMessageResponse) GoString() string {
	return s.String()
}

func (s *SendMessageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendMessageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendMessageResponse) GetId() *string {
	return s.Id
}

func (s *SendMessageResponse) GetEvent() *string {
	return s.Event
}

func (s *SendMessageResponse) GetBody() *string {
	return s.Body
}

func (s *SendMessageResponse) SetHeaders(v map[string]*string) *SendMessageResponse {
	s.Headers = v
	return s
}

func (s *SendMessageResponse) SetStatusCode(v int32) *SendMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *SendMessageResponse) SetId(v string) *SendMessageResponse {
	s.Id = &v
	return s
}

func (s *SendMessageResponse) SetEvent(v string) *SendMessageResponse {
	s.Event = &v
	return s
}

func (s *SendMessageResponse) SetBody(v string) *SendMessageResponse {
	s.Body = &v
	return s
}

func (s *SendMessageResponse) Validate() error {
	return dara.Validate(s)
}
