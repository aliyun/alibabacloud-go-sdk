// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgeBaseJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListKnowledgeBaseJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListKnowledgeBaseJobsResponse
	GetStatusCode() *int32
	SetBody(v *ListKnowledgeBaseJobsResponseBody) *ListKnowledgeBaseJobsResponse
	GetBody() *ListKnowledgeBaseJobsResponseBody
}

type ListKnowledgeBaseJobsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListKnowledgeBaseJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListKnowledgeBaseJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBaseJobsResponse) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBaseJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListKnowledgeBaseJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListKnowledgeBaseJobsResponse) GetBody() *ListKnowledgeBaseJobsResponseBody {
	return s.Body
}

func (s *ListKnowledgeBaseJobsResponse) SetHeaders(v map[string]*string) *ListKnowledgeBaseJobsResponse {
	s.Headers = v
	return s
}

func (s *ListKnowledgeBaseJobsResponse) SetStatusCode(v int32) *ListKnowledgeBaseJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListKnowledgeBaseJobsResponse) SetBody(v *ListKnowledgeBaseJobsResponseBody) *ListKnowledgeBaseJobsResponse {
	s.Body = v
	return s
}

func (s *ListKnowledgeBaseJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
