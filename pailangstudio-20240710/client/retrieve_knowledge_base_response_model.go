// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetrieveKnowledgeBaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RetrieveKnowledgeBaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RetrieveKnowledgeBaseResponse
	GetStatusCode() *int32
	SetBody(v *RetrieveKnowledgeBaseResponseBody) *RetrieveKnowledgeBaseResponse
	GetBody() *RetrieveKnowledgeBaseResponseBody
}

type RetrieveKnowledgeBaseResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RetrieveKnowledgeBaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetrieveKnowledgeBaseResponse) String() string {
	return dara.Prettify(s)
}

func (s RetrieveKnowledgeBaseResponse) GoString() string {
	return s.String()
}

func (s *RetrieveKnowledgeBaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RetrieveKnowledgeBaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RetrieveKnowledgeBaseResponse) GetBody() *RetrieveKnowledgeBaseResponseBody {
	return s.Body
}

func (s *RetrieveKnowledgeBaseResponse) SetHeaders(v map[string]*string) *RetrieveKnowledgeBaseResponse {
	s.Headers = v
	return s
}

func (s *RetrieveKnowledgeBaseResponse) SetStatusCode(v int32) *RetrieveKnowledgeBaseResponse {
	s.StatusCode = &v
	return s
}

func (s *RetrieveKnowledgeBaseResponse) SetBody(v *RetrieveKnowledgeBaseResponseBody) *RetrieveKnowledgeBaseResponse {
	s.Body = v
	return s
}

func (s *RetrieveKnowledgeBaseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
