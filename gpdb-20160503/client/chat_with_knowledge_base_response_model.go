// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatWithKnowledgeBaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChatWithKnowledgeBaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChatWithKnowledgeBaseResponse
	GetStatusCode() *int32
	SetBody(v *ChatWithKnowledgeBaseResponseBody) *ChatWithKnowledgeBaseResponse
	GetBody() *ChatWithKnowledgeBaseResponseBody
}

type ChatWithKnowledgeBaseResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChatWithKnowledgeBaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChatWithKnowledgeBaseResponse) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseResponse) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChatWithKnowledgeBaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChatWithKnowledgeBaseResponse) GetBody() *ChatWithKnowledgeBaseResponseBody {
	return s.Body
}

func (s *ChatWithKnowledgeBaseResponse) SetHeaders(v map[string]*string) *ChatWithKnowledgeBaseResponse {
	s.Headers = v
	return s
}

func (s *ChatWithKnowledgeBaseResponse) SetStatusCode(v int32) *ChatWithKnowledgeBaseResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatWithKnowledgeBaseResponse) SetBody(v *ChatWithKnowledgeBaseResponseBody) *ChatWithKnowledgeBaseResponse {
	s.Body = v
	return s
}

func (s *ChatWithKnowledgeBaseResponse) Validate() error {
	return dara.Validate(s)
}
