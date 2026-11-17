// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgeObjectsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListKnowledgeObjectsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListKnowledgeObjectsResponse
	GetStatusCode() *int32
	SetBody(v *ListKnowledgeObjectsResponseBody) *ListKnowledgeObjectsResponse
	GetBody() *ListKnowledgeObjectsResponseBody
}

type ListKnowledgeObjectsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListKnowledgeObjectsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListKnowledgeObjectsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeObjectsResponse) GoString() string {
	return s.String()
}

func (s *ListKnowledgeObjectsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListKnowledgeObjectsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListKnowledgeObjectsResponse) GetBody() *ListKnowledgeObjectsResponseBody {
	return s.Body
}

func (s *ListKnowledgeObjectsResponse) SetHeaders(v map[string]*string) *ListKnowledgeObjectsResponse {
	s.Headers = v
	return s
}

func (s *ListKnowledgeObjectsResponse) SetStatusCode(v int32) *ListKnowledgeObjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListKnowledgeObjectsResponse) SetBody(v *ListKnowledgeObjectsResponseBody) *ListKnowledgeObjectsResponse {
	s.Body = v
	return s
}

func (s *ListKnowledgeObjectsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
