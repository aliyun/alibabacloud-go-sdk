// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChatResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChatResponse
	GetStatusCode() *int32
	SetId(v string) *ChatResponse
	GetId() *string
	SetEvent(v string) *ChatResponse
	GetEvent() *string
	SetBody(v *ChatResponseBody) *ChatResponse
	GetBody() *ChatResponseBody
}

type ChatResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string            `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string            `json:"event,omitempty" xml:"event,omitempty"`
	Body       *ChatResponseBody  `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChatResponse) String() string {
	return dara.Prettify(s)
}

func (s ChatResponse) GoString() string {
	return s.String()
}

func (s *ChatResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChatResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChatResponse) GetId() *string {
	return s.Id
}

func (s *ChatResponse) GetEvent() *string {
	return s.Event
}

func (s *ChatResponse) GetBody() *ChatResponseBody {
	return s.Body
}

func (s *ChatResponse) SetHeaders(v map[string]*string) *ChatResponse {
	s.Headers = v
	return s
}

func (s *ChatResponse) SetStatusCode(v int32) *ChatResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatResponse) SetId(v string) *ChatResponse {
	s.Id = &v
	return s
}

func (s *ChatResponse) SetEvent(v string) *ChatResponse {
	s.Event = &v
	return s
}

func (s *ChatResponse) SetBody(v *ChatResponseBody) *ChatResponse {
	s.Body = v
	return s
}

func (s *ChatResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
