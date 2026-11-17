// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgeBaseRetrievalGrantsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListKnowledgeBaseRetrievalGrantsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListKnowledgeBaseRetrievalGrantsResponse
	GetStatusCode() *int32
	SetBody(v *ListKnowledgeBaseRetrievalGrantsResponseBody) *ListKnowledgeBaseRetrievalGrantsResponse
	GetBody() *ListKnowledgeBaseRetrievalGrantsResponseBody
}

type ListKnowledgeBaseRetrievalGrantsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListKnowledgeBaseRetrievalGrantsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListKnowledgeBaseRetrievalGrantsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBaseRetrievalGrantsResponse) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBaseRetrievalGrantsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListKnowledgeBaseRetrievalGrantsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListKnowledgeBaseRetrievalGrantsResponse) GetBody() *ListKnowledgeBaseRetrievalGrantsResponseBody {
	return s.Body
}

func (s *ListKnowledgeBaseRetrievalGrantsResponse) SetHeaders(v map[string]*string) *ListKnowledgeBaseRetrievalGrantsResponse {
	s.Headers = v
	return s
}

func (s *ListKnowledgeBaseRetrievalGrantsResponse) SetStatusCode(v int32) *ListKnowledgeBaseRetrievalGrantsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListKnowledgeBaseRetrievalGrantsResponse) SetBody(v *ListKnowledgeBaseRetrievalGrantsResponseBody) *ListKnowledgeBaseRetrievalGrantsResponse {
	s.Body = v
	return s
}

func (s *ListKnowledgeBaseRetrievalGrantsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
