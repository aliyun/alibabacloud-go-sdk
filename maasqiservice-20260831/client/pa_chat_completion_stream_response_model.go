// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPaChatCompletionStreamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PaChatCompletionStreamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PaChatCompletionStreamResponse
	GetStatusCode() *int32
	SetId(v string) *PaChatCompletionStreamResponse
	GetId() *string
	SetEvent(v string) *PaChatCompletionStreamResponse
	GetEvent() *string
	SetBody(v *PaChatCompletionStreamResponseBody) *PaChatCompletionStreamResponse
	GetBody() *PaChatCompletionStreamResponseBody
}

type PaChatCompletionStreamResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                             `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                             `json:"event,omitempty" xml:"event,omitempty"`
	Body       *PaChatCompletionStreamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PaChatCompletionStreamResponse) String() string {
	return dara.Prettify(s)
}

func (s PaChatCompletionStreamResponse) GoString() string {
	return s.String()
}

func (s *PaChatCompletionStreamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PaChatCompletionStreamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PaChatCompletionStreamResponse) GetId() *string {
	return s.Id
}

func (s *PaChatCompletionStreamResponse) GetEvent() *string {
	return s.Event
}

func (s *PaChatCompletionStreamResponse) GetBody() *PaChatCompletionStreamResponseBody {
	return s.Body
}

func (s *PaChatCompletionStreamResponse) SetHeaders(v map[string]*string) *PaChatCompletionStreamResponse {
	s.Headers = v
	return s
}

func (s *PaChatCompletionStreamResponse) SetStatusCode(v int32) *PaChatCompletionStreamResponse {
	s.StatusCode = &v
	return s
}

func (s *PaChatCompletionStreamResponse) SetId(v string) *PaChatCompletionStreamResponse {
	s.Id = &v
	return s
}

func (s *PaChatCompletionStreamResponse) SetEvent(v string) *PaChatCompletionStreamResponse {
	s.Event = &v
	return s
}

func (s *PaChatCompletionStreamResponse) SetBody(v *PaChatCompletionStreamResponseBody) *PaChatCompletionStreamResponse {
	s.Body = v
	return s
}

func (s *PaChatCompletionStreamResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
