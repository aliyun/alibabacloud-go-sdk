// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchLumaKnowledgeBaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchLumaKnowledgeBaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchLumaKnowledgeBaseResponse
	GetStatusCode() *int32
	SetBody(v *SearchLumaKnowledgeBaseResponseBody) *SearchLumaKnowledgeBaseResponse
	GetBody() *SearchLumaKnowledgeBaseResponseBody
}

type SearchLumaKnowledgeBaseResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchLumaKnowledgeBaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchLumaKnowledgeBaseResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchLumaKnowledgeBaseResponse) GoString() string {
	return s.String()
}

func (s *SearchLumaKnowledgeBaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchLumaKnowledgeBaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchLumaKnowledgeBaseResponse) GetBody() *SearchLumaKnowledgeBaseResponseBody {
	return s.Body
}

func (s *SearchLumaKnowledgeBaseResponse) SetHeaders(v map[string]*string) *SearchLumaKnowledgeBaseResponse {
	s.Headers = v
	return s
}

func (s *SearchLumaKnowledgeBaseResponse) SetStatusCode(v int32) *SearchLumaKnowledgeBaseResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchLumaKnowledgeBaseResponse) SetBody(v *SearchLumaKnowledgeBaseResponseBody) *SearchLumaKnowledgeBaseResponse {
	s.Body = v
	return s
}

func (s *SearchLumaKnowledgeBaseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
