// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAigcChatCompletionStreamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AigcChatCompletionStreamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AigcChatCompletionStreamResponse
	GetStatusCode() *int32
	SetId(v string) *AigcChatCompletionStreamResponse
	GetId() *string
	SetEvent(v string) *AigcChatCompletionStreamResponse
	GetEvent() *string
	SetBody(v *AigcChatCompletionStreamResponseBody) *AigcChatCompletionStreamResponse
	GetBody() *AigcChatCompletionStreamResponseBody
}

type AigcChatCompletionStreamResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                               `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                               `json:"event,omitempty" xml:"event,omitempty"`
	Body       *AigcChatCompletionStreamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AigcChatCompletionStreamResponse) String() string {
	return dara.Prettify(s)
}

func (s AigcChatCompletionStreamResponse) GoString() string {
	return s.String()
}

func (s *AigcChatCompletionStreamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AigcChatCompletionStreamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AigcChatCompletionStreamResponse) GetId() *string {
	return s.Id
}

func (s *AigcChatCompletionStreamResponse) GetEvent() *string {
	return s.Event
}

func (s *AigcChatCompletionStreamResponse) GetBody() *AigcChatCompletionStreamResponseBody {
	return s.Body
}

func (s *AigcChatCompletionStreamResponse) SetHeaders(v map[string]*string) *AigcChatCompletionStreamResponse {
	s.Headers = v
	return s
}

func (s *AigcChatCompletionStreamResponse) SetStatusCode(v int32) *AigcChatCompletionStreamResponse {
	s.StatusCode = &v
	return s
}

func (s *AigcChatCompletionStreamResponse) SetId(v string) *AigcChatCompletionStreamResponse {
	s.Id = &v
	return s
}

func (s *AigcChatCompletionStreamResponse) SetEvent(v string) *AigcChatCompletionStreamResponse {
	s.Event = &v
	return s
}

func (s *AigcChatCompletionStreamResponse) SetBody(v *AigcChatCompletionStreamResponseBody) *AigcChatCompletionStreamResponse {
	s.Body = v
	return s
}

func (s *AigcChatCompletionStreamResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
