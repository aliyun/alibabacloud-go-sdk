// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatWithKnowledgeBaseStreamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChatWithKnowledgeBaseStreamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChatWithKnowledgeBaseStreamResponse
	GetStatusCode() *int32
	SetBody(v *ChatWithKnowledgeBaseStreamResponseBody) *ChatWithKnowledgeBaseStreamResponse
	GetBody() *ChatWithKnowledgeBaseStreamResponseBody
}

type ChatWithKnowledgeBaseStreamResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChatWithKnowledgeBaseStreamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChatWithKnowledgeBaseStreamResponse) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseStreamResponse) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseStreamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChatWithKnowledgeBaseStreamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChatWithKnowledgeBaseStreamResponse) GetBody() *ChatWithKnowledgeBaseStreamResponseBody {
	return s.Body
}

func (s *ChatWithKnowledgeBaseStreamResponse) SetHeaders(v map[string]*string) *ChatWithKnowledgeBaseStreamResponse {
	s.Headers = v
	return s
}

func (s *ChatWithKnowledgeBaseStreamResponse) SetStatusCode(v int32) *ChatWithKnowledgeBaseStreamResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamResponse) SetBody(v *ChatWithKnowledgeBaseStreamResponseBody) *ChatWithKnowledgeBaseStreamResponse {
	s.Body = v
	return s
}

func (s *ChatWithKnowledgeBaseStreamResponse) Validate() error {
	return dara.Validate(s)
}
