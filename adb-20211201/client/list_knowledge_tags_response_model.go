// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgeTagsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListKnowledgeTagsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListKnowledgeTagsResponse
	GetStatusCode() *int32
	SetBody(v *ListKnowledgeTagsResponseBody) *ListKnowledgeTagsResponse
	GetBody() *ListKnowledgeTagsResponseBody
}

type ListKnowledgeTagsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListKnowledgeTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListKnowledgeTagsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeTagsResponse) GoString() string {
	return s.String()
}

func (s *ListKnowledgeTagsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListKnowledgeTagsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListKnowledgeTagsResponse) GetBody() *ListKnowledgeTagsResponseBody {
	return s.Body
}

func (s *ListKnowledgeTagsResponse) SetHeaders(v map[string]*string) *ListKnowledgeTagsResponse {
	s.Headers = v
	return s
}

func (s *ListKnowledgeTagsResponse) SetStatusCode(v int32) *ListKnowledgeTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListKnowledgeTagsResponse) SetBody(v *ListKnowledgeTagsResponseBody) *ListKnowledgeTagsResponse {
	s.Body = v
	return s
}

func (s *ListKnowledgeTagsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
