// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVisibleKnowledgeBaseContentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVisibleKnowledgeBaseContentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVisibleKnowledgeBaseContentsResponse
	GetStatusCode() *int32
	SetBody(v *ListVisibleKnowledgeBaseContentsResponseBody) *ListVisibleKnowledgeBaseContentsResponse
	GetBody() *ListVisibleKnowledgeBaseContentsResponseBody
}

type ListVisibleKnowledgeBaseContentsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVisibleKnowledgeBaseContentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVisibleKnowledgeBaseContentsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVisibleKnowledgeBaseContentsResponse) GoString() string {
	return s.String()
}

func (s *ListVisibleKnowledgeBaseContentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVisibleKnowledgeBaseContentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVisibleKnowledgeBaseContentsResponse) GetBody() *ListVisibleKnowledgeBaseContentsResponseBody {
	return s.Body
}

func (s *ListVisibleKnowledgeBaseContentsResponse) SetHeaders(v map[string]*string) *ListVisibleKnowledgeBaseContentsResponse {
	s.Headers = v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsResponse) SetStatusCode(v int32) *ListVisibleKnowledgeBaseContentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsResponse) SetBody(v *ListVisibleKnowledgeBaseContentsResponseBody) *ListVisibleKnowledgeBaseContentsResponse {
	s.Body = v
	return s
}

func (s *ListVisibleKnowledgeBaseContentsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
