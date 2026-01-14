// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgeBaseChunksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListKnowledgeBaseChunksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListKnowledgeBaseChunksResponse
	GetStatusCode() *int32
	SetBody(v *ListKnowledgeBaseChunksResponseBody) *ListKnowledgeBaseChunksResponse
	GetBody() *ListKnowledgeBaseChunksResponseBody
}

type ListKnowledgeBaseChunksResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListKnowledgeBaseChunksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListKnowledgeBaseChunksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBaseChunksResponse) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBaseChunksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListKnowledgeBaseChunksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListKnowledgeBaseChunksResponse) GetBody() *ListKnowledgeBaseChunksResponseBody {
	return s.Body
}

func (s *ListKnowledgeBaseChunksResponse) SetHeaders(v map[string]*string) *ListKnowledgeBaseChunksResponse {
	s.Headers = v
	return s
}

func (s *ListKnowledgeBaseChunksResponse) SetStatusCode(v int32) *ListKnowledgeBaseChunksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListKnowledgeBaseChunksResponse) SetBody(v *ListKnowledgeBaseChunksResponseBody) *ListKnowledgeBaseChunksResponse {
	s.Body = v
	return s
}

func (s *ListKnowledgeBaseChunksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
