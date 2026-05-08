// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgeBaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListKnowledgeBaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListKnowledgeBaseResponse
	GetStatusCode() *int32
	SetBody(v *KnowledgeBaseListResult) *ListKnowledgeBaseResponse
	GetBody() *KnowledgeBaseListResult
}

type ListKnowledgeBaseResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *KnowledgeBaseListResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListKnowledgeBaseResponse) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBaseResponse) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListKnowledgeBaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListKnowledgeBaseResponse) GetBody() *KnowledgeBaseListResult {
	return s.Body
}

func (s *ListKnowledgeBaseResponse) SetHeaders(v map[string]*string) *ListKnowledgeBaseResponse {
	s.Headers = v
	return s
}

func (s *ListKnowledgeBaseResponse) SetStatusCode(v int32) *ListKnowledgeBaseResponse {
	s.StatusCode = &v
	return s
}

func (s *ListKnowledgeBaseResponse) SetBody(v *KnowledgeBaseListResult) *ListKnowledgeBaseResponse {
	s.Body = v
	return s
}

func (s *ListKnowledgeBaseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
