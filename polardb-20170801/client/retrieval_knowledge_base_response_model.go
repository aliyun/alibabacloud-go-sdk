// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetrievalKnowledgeBaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RetrievalKnowledgeBaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RetrievalKnowledgeBaseResponse
	GetStatusCode() *int32
	SetBody(v *RetrievalKnowledgeBaseResponseBody) *RetrievalKnowledgeBaseResponse
	GetBody() *RetrievalKnowledgeBaseResponseBody
}

type RetrievalKnowledgeBaseResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RetrievalKnowledgeBaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetrievalKnowledgeBaseResponse) String() string {
	return dara.Prettify(s)
}

func (s RetrievalKnowledgeBaseResponse) GoString() string {
	return s.String()
}

func (s *RetrievalKnowledgeBaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RetrievalKnowledgeBaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RetrievalKnowledgeBaseResponse) GetBody() *RetrievalKnowledgeBaseResponseBody {
	return s.Body
}

func (s *RetrievalKnowledgeBaseResponse) SetHeaders(v map[string]*string) *RetrievalKnowledgeBaseResponse {
	s.Headers = v
	return s
}

func (s *RetrievalKnowledgeBaseResponse) SetStatusCode(v int32) *RetrievalKnowledgeBaseResponse {
	s.StatusCode = &v
	return s
}

func (s *RetrievalKnowledgeBaseResponse) SetBody(v *RetrievalKnowledgeBaseResponseBody) *RetrievalKnowledgeBaseResponse {
	s.Body = v
	return s
}

func (s *RetrievalKnowledgeBaseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
