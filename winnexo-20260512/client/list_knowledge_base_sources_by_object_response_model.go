// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgeBaseSourcesByObjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListKnowledgeBaseSourcesByObjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListKnowledgeBaseSourcesByObjectResponse
	GetStatusCode() *int32
	SetBody(v *ListKnowledgeBaseSourcesByObjectResponseBody) *ListKnowledgeBaseSourcesByObjectResponse
	GetBody() *ListKnowledgeBaseSourcesByObjectResponseBody
}

type ListKnowledgeBaseSourcesByObjectResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListKnowledgeBaseSourcesByObjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListKnowledgeBaseSourcesByObjectResponse) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBaseSourcesByObjectResponse) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBaseSourcesByObjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListKnowledgeBaseSourcesByObjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListKnowledgeBaseSourcesByObjectResponse) GetBody() *ListKnowledgeBaseSourcesByObjectResponseBody {
	return s.Body
}

func (s *ListKnowledgeBaseSourcesByObjectResponse) SetHeaders(v map[string]*string) *ListKnowledgeBaseSourcesByObjectResponse {
	s.Headers = v
	return s
}

func (s *ListKnowledgeBaseSourcesByObjectResponse) SetStatusCode(v int32) *ListKnowledgeBaseSourcesByObjectResponse {
	s.StatusCode = &v
	return s
}

func (s *ListKnowledgeBaseSourcesByObjectResponse) SetBody(v *ListKnowledgeBaseSourcesByObjectResponseBody) *ListKnowledgeBaseSourcesByObjectResponse {
	s.Body = v
	return s
}

func (s *ListKnowledgeBaseSourcesByObjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
