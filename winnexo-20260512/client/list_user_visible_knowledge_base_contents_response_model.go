// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserVisibleKnowledgeBaseContentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserVisibleKnowledgeBaseContentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserVisibleKnowledgeBaseContentsResponse
	GetStatusCode() *int32
	SetBody(v *ListUserVisibleKnowledgeBaseContentsResponseBody) *ListUserVisibleKnowledgeBaseContentsResponse
	GetBody() *ListUserVisibleKnowledgeBaseContentsResponseBody
}

type ListUserVisibleKnowledgeBaseContentsResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserVisibleKnowledgeBaseContentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserVisibleKnowledgeBaseContentsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserVisibleKnowledgeBaseContentsResponse) GoString() string {
	return s.String()
}

func (s *ListUserVisibleKnowledgeBaseContentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserVisibleKnowledgeBaseContentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserVisibleKnowledgeBaseContentsResponse) GetBody() *ListUserVisibleKnowledgeBaseContentsResponseBody {
	return s.Body
}

func (s *ListUserVisibleKnowledgeBaseContentsResponse) SetHeaders(v map[string]*string) *ListUserVisibleKnowledgeBaseContentsResponse {
	s.Headers = v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponse) SetStatusCode(v int32) *ListUserVisibleKnowledgeBaseContentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponse) SetBody(v *ListUserVisibleKnowledgeBaseContentsResponseBody) *ListUserVisibleKnowledgeBaseContentsResponse {
	s.Body = v
	return s
}

func (s *ListUserVisibleKnowledgeBaseContentsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
