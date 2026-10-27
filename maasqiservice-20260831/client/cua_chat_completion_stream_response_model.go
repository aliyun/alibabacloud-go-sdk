// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCuaChatCompletionStreamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CuaChatCompletionStreamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CuaChatCompletionStreamResponse
	GetStatusCode() *int32
	SetId(v string) *CuaChatCompletionStreamResponse
	GetId() *string
	SetEvent(v string) *CuaChatCompletionStreamResponse
	GetEvent() *string
	SetBody(v *CuaChatCompletionStreamResponseBody) *CuaChatCompletionStreamResponse
	GetBody() *CuaChatCompletionStreamResponseBody
}

type CuaChatCompletionStreamResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                              `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                              `json:"event,omitempty" xml:"event,omitempty"`
	Body       *CuaChatCompletionStreamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CuaChatCompletionStreamResponse) String() string {
	return dara.Prettify(s)
}

func (s CuaChatCompletionStreamResponse) GoString() string {
	return s.String()
}

func (s *CuaChatCompletionStreamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CuaChatCompletionStreamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CuaChatCompletionStreamResponse) GetId() *string {
	return s.Id
}

func (s *CuaChatCompletionStreamResponse) GetEvent() *string {
	return s.Event
}

func (s *CuaChatCompletionStreamResponse) GetBody() *CuaChatCompletionStreamResponseBody {
	return s.Body
}

func (s *CuaChatCompletionStreamResponse) SetHeaders(v map[string]*string) *CuaChatCompletionStreamResponse {
	s.Headers = v
	return s
}

func (s *CuaChatCompletionStreamResponse) SetStatusCode(v int32) *CuaChatCompletionStreamResponse {
	s.StatusCode = &v
	return s
}

func (s *CuaChatCompletionStreamResponse) SetId(v string) *CuaChatCompletionStreamResponse {
	s.Id = &v
	return s
}

func (s *CuaChatCompletionStreamResponse) SetEvent(v string) *CuaChatCompletionStreamResponse {
	s.Event = &v
	return s
}

func (s *CuaChatCompletionStreamResponse) SetBody(v *CuaChatCompletionStreamResponseBody) *CuaChatCompletionStreamResponse {
	s.Body = v
	return s
}

func (s *CuaChatCompletionStreamResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
