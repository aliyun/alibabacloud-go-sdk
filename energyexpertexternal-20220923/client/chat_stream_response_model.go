// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatStreamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChatStreamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChatStreamResponse
	GetStatusCode() *int32
	SetId(v string) *ChatStreamResponse
	GetId() *string
	SetEvent(v string) *ChatStreamResponse
	GetEvent() *string
	SetBody(v *ChatStreamResponseBody) *ChatStreamResponse
	GetBody() *ChatStreamResponseBody
}

type ChatStreamResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                 `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                 `json:"event,omitempty" xml:"event,omitempty"`
	Body       *ChatStreamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChatStreamResponse) String() string {
	return dara.Prettify(s)
}

func (s ChatStreamResponse) GoString() string {
	return s.String()
}

func (s *ChatStreamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChatStreamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChatStreamResponse) GetId() *string {
	return s.Id
}

func (s *ChatStreamResponse) GetEvent() *string {
	return s.Event
}

func (s *ChatStreamResponse) GetBody() *ChatStreamResponseBody {
	return s.Body
}

func (s *ChatStreamResponse) SetHeaders(v map[string]*string) *ChatStreamResponse {
	s.Headers = v
	return s
}

func (s *ChatStreamResponse) SetStatusCode(v int32) *ChatStreamResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatStreamResponse) SetId(v string) *ChatStreamResponse {
	s.Id = &v
	return s
}

func (s *ChatStreamResponse) SetEvent(v string) *ChatStreamResponse {
	s.Event = &v
	return s
}

func (s *ChatStreamResponse) SetBody(v *ChatStreamResponseBody) *ChatStreamResponse {
	s.Body = v
	return s
}

func (s *ChatStreamResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
