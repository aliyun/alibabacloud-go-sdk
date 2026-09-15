// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGuiChatCompletionStreamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GuiChatCompletionStreamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GuiChatCompletionStreamResponse
	GetStatusCode() *int32
	SetId(v string) *GuiChatCompletionStreamResponse
	GetId() *string
	SetEvent(v string) *GuiChatCompletionStreamResponse
	GetEvent() *string
	SetBody(v *GuiChatCompletionStreamResponseBody) *GuiChatCompletionStreamResponse
	GetBody() *GuiChatCompletionStreamResponseBody
}

type GuiChatCompletionStreamResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                              `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                              `json:"event,omitempty" xml:"event,omitempty"`
	Body       *GuiChatCompletionStreamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GuiChatCompletionStreamResponse) String() string {
	return dara.Prettify(s)
}

func (s GuiChatCompletionStreamResponse) GoString() string {
	return s.String()
}

func (s *GuiChatCompletionStreamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GuiChatCompletionStreamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GuiChatCompletionStreamResponse) GetId() *string {
	return s.Id
}

func (s *GuiChatCompletionStreamResponse) GetEvent() *string {
	return s.Event
}

func (s *GuiChatCompletionStreamResponse) GetBody() *GuiChatCompletionStreamResponseBody {
	return s.Body
}

func (s *GuiChatCompletionStreamResponse) SetHeaders(v map[string]*string) *GuiChatCompletionStreamResponse {
	s.Headers = v
	return s
}

func (s *GuiChatCompletionStreamResponse) SetStatusCode(v int32) *GuiChatCompletionStreamResponse {
	s.StatusCode = &v
	return s
}

func (s *GuiChatCompletionStreamResponse) SetId(v string) *GuiChatCompletionStreamResponse {
	s.Id = &v
	return s
}

func (s *GuiChatCompletionStreamResponse) SetEvent(v string) *GuiChatCompletionStreamResponse {
	s.Event = &v
	return s
}

func (s *GuiChatCompletionStreamResponse) SetBody(v *GuiChatCompletionStreamResponseBody) *GuiChatCompletionStreamResponse {
	s.Body = v
	return s
}

func (s *GuiChatCompletionStreamResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
