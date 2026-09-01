// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAnswerKnowledgeBaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AnswerKnowledgeBaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AnswerKnowledgeBaseResponse
	GetStatusCode() *int32
	SetBody(v *AnswerKnowledgeBaseResponseBody) *AnswerKnowledgeBaseResponse
	GetBody() *AnswerKnowledgeBaseResponseBody
}

type AnswerKnowledgeBaseResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AnswerKnowledgeBaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AnswerKnowledgeBaseResponse) String() string {
	return dara.Prettify(s)
}

func (s AnswerKnowledgeBaseResponse) GoString() string {
	return s.String()
}

func (s *AnswerKnowledgeBaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AnswerKnowledgeBaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AnswerKnowledgeBaseResponse) GetBody() *AnswerKnowledgeBaseResponseBody {
	return s.Body
}

func (s *AnswerKnowledgeBaseResponse) SetHeaders(v map[string]*string) *AnswerKnowledgeBaseResponse {
	s.Headers = v
	return s
}

func (s *AnswerKnowledgeBaseResponse) SetStatusCode(v int32) *AnswerKnowledgeBaseResponse {
	s.StatusCode = &v
	return s
}

func (s *AnswerKnowledgeBaseResponse) SetBody(v *AnswerKnowledgeBaseResponseBody) *AnswerKnowledgeBaseResponse {
	s.Body = v
	return s
}

func (s *AnswerKnowledgeBaseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
