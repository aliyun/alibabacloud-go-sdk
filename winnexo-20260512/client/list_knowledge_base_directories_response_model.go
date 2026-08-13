// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgeBaseDirectoriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListKnowledgeBaseDirectoriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListKnowledgeBaseDirectoriesResponse
	GetStatusCode() *int32
	SetBody(v *ListKnowledgeBaseDirectoriesResponseBody) *ListKnowledgeBaseDirectoriesResponse
	GetBody() *ListKnowledgeBaseDirectoriesResponseBody
}

type ListKnowledgeBaseDirectoriesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListKnowledgeBaseDirectoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListKnowledgeBaseDirectoriesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBaseDirectoriesResponse) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBaseDirectoriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListKnowledgeBaseDirectoriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListKnowledgeBaseDirectoriesResponse) GetBody() *ListKnowledgeBaseDirectoriesResponseBody {
	return s.Body
}

func (s *ListKnowledgeBaseDirectoriesResponse) SetHeaders(v map[string]*string) *ListKnowledgeBaseDirectoriesResponse {
	s.Headers = v
	return s
}

func (s *ListKnowledgeBaseDirectoriesResponse) SetStatusCode(v int32) *ListKnowledgeBaseDirectoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListKnowledgeBaseDirectoriesResponse) SetBody(v *ListKnowledgeBaseDirectoriesResponseBody) *ListKnowledgeBaseDirectoriesResponse {
	s.Body = v
	return s
}

func (s *ListKnowledgeBaseDirectoriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
