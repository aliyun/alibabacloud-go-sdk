// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantKnowledgeBaseRetrievalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GrantKnowledgeBaseRetrievalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GrantKnowledgeBaseRetrievalResponse
	GetStatusCode() *int32
	SetBody(v *GrantKnowledgeBaseRetrievalResponseBody) *GrantKnowledgeBaseRetrievalResponse
	GetBody() *GrantKnowledgeBaseRetrievalResponseBody
}

type GrantKnowledgeBaseRetrievalResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GrantKnowledgeBaseRetrievalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GrantKnowledgeBaseRetrievalResponse) String() string {
	return dara.Prettify(s)
}

func (s GrantKnowledgeBaseRetrievalResponse) GoString() string {
	return s.String()
}

func (s *GrantKnowledgeBaseRetrievalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GrantKnowledgeBaseRetrievalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GrantKnowledgeBaseRetrievalResponse) GetBody() *GrantKnowledgeBaseRetrievalResponseBody {
	return s.Body
}

func (s *GrantKnowledgeBaseRetrievalResponse) SetHeaders(v map[string]*string) *GrantKnowledgeBaseRetrievalResponse {
	s.Headers = v
	return s
}

func (s *GrantKnowledgeBaseRetrievalResponse) SetStatusCode(v int32) *GrantKnowledgeBaseRetrievalResponse {
	s.StatusCode = &v
	return s
}

func (s *GrantKnowledgeBaseRetrievalResponse) SetBody(v *GrantKnowledgeBaseRetrievalResponseBody) *GrantKnowledgeBaseRetrievalResponse {
	s.Body = v
	return s
}

func (s *GrantKnowledgeBaseRetrievalResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
