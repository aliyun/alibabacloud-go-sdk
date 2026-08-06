// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchKnowledgeBaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchKnowledgeBaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchKnowledgeBaseResponse
	GetStatusCode() *int32
	SetBody(v *SearchKnowledgeBaseResponseBody) *SearchKnowledgeBaseResponse
	GetBody() *SearchKnowledgeBaseResponseBody
}

type SearchKnowledgeBaseResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchKnowledgeBaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchKnowledgeBaseResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchKnowledgeBaseResponse) GoString() string {
	return s.String()
}

func (s *SearchKnowledgeBaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchKnowledgeBaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchKnowledgeBaseResponse) GetBody() *SearchKnowledgeBaseResponseBody {
	return s.Body
}

func (s *SearchKnowledgeBaseResponse) SetHeaders(v map[string]*string) *SearchKnowledgeBaseResponse {
	s.Headers = v
	return s
}

func (s *SearchKnowledgeBaseResponse) SetStatusCode(v int32) *SearchKnowledgeBaseResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchKnowledgeBaseResponse) SetBody(v *SearchKnowledgeBaseResponseBody) *SearchKnowledgeBaseResponse {
	s.Body = v
	return s
}

func (s *SearchKnowledgeBaseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
