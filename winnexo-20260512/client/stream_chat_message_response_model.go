// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStreamChatMessageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StreamChatMessageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StreamChatMessageResponse
	GetStatusCode() *int32
	SetId(v string) *StreamChatMessageResponse
	GetId() *string
	SetEvent(v string) *StreamChatMessageResponse
	GetEvent() *string
	SetBody(v *StreamChatMessageResponseBody) *StreamChatMessageResponse
	GetBody() *StreamChatMessageResponseBody
}

type StreamChatMessageResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                        `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                        `json:"event,omitempty" xml:"event,omitempty"`
	Body       *StreamChatMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StreamChatMessageResponse) String() string {
	return dara.Prettify(s)
}

func (s StreamChatMessageResponse) GoString() string {
	return s.String()
}

func (s *StreamChatMessageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StreamChatMessageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StreamChatMessageResponse) GetId() *string {
	return s.Id
}

func (s *StreamChatMessageResponse) GetEvent() *string {
	return s.Event
}

func (s *StreamChatMessageResponse) GetBody() *StreamChatMessageResponseBody {
	return s.Body
}

func (s *StreamChatMessageResponse) SetHeaders(v map[string]*string) *StreamChatMessageResponse {
	s.Headers = v
	return s
}

func (s *StreamChatMessageResponse) SetStatusCode(v int32) *StreamChatMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *StreamChatMessageResponse) SetId(v string) *StreamChatMessageResponse {
	s.Id = &v
	return s
}

func (s *StreamChatMessageResponse) SetEvent(v string) *StreamChatMessageResponse {
	s.Event = &v
	return s
}

func (s *StreamChatMessageResponse) SetBody(v *StreamChatMessageResponseBody) *StreamChatMessageResponse {
	s.Body = v
	return s
}

func (s *StreamChatMessageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
