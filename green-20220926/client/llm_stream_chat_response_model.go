// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLlmStreamChatResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *LlmStreamChatResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *LlmStreamChatResponse
	GetStatusCode() *int32
	SetBody(v *LlmStreamChatResponseBody) *LlmStreamChatResponse
	GetBody() *LlmStreamChatResponseBody
}

type LlmStreamChatResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LlmStreamChatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LlmStreamChatResponse) String() string {
	return dara.Prettify(s)
}

func (s LlmStreamChatResponse) GoString() string {
	return s.String()
}

func (s *LlmStreamChatResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *LlmStreamChatResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *LlmStreamChatResponse) GetBody() *LlmStreamChatResponseBody {
	return s.Body
}

func (s *LlmStreamChatResponse) SetHeaders(v map[string]*string) *LlmStreamChatResponse {
	s.Headers = v
	return s
}

func (s *LlmStreamChatResponse) SetStatusCode(v int32) *LlmStreamChatResponse {
	s.StatusCode = &v
	return s
}

func (s *LlmStreamChatResponse) SetBody(v *LlmStreamChatResponseBody) *LlmStreamChatResponse {
	s.Body = v
	return s
}

func (s *LlmStreamChatResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
