// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgeSourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListKnowledgeSourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListKnowledgeSourcesResponse
	GetStatusCode() *int32
	SetBody(v *ListKnowledgeSourcesResponseBody) *ListKnowledgeSourcesResponse
	GetBody() *ListKnowledgeSourcesResponseBody
}

type ListKnowledgeSourcesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListKnowledgeSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListKnowledgeSourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeSourcesResponse) GoString() string {
	return s.String()
}

func (s *ListKnowledgeSourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListKnowledgeSourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListKnowledgeSourcesResponse) GetBody() *ListKnowledgeSourcesResponseBody {
	return s.Body
}

func (s *ListKnowledgeSourcesResponse) SetHeaders(v map[string]*string) *ListKnowledgeSourcesResponse {
	s.Headers = v
	return s
}

func (s *ListKnowledgeSourcesResponse) SetStatusCode(v int32) *ListKnowledgeSourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListKnowledgeSourcesResponse) SetBody(v *ListKnowledgeSourcesResponseBody) *ListKnowledgeSourcesResponse {
	s.Body = v
	return s
}

func (s *ListKnowledgeSourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
