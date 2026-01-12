// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKnowledgeBaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetKnowledgeBaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetKnowledgeBaseResponse
	GetStatusCode() *int32
	SetBody(v *KnowledgeBaseResult) *GetKnowledgeBaseResponse
	GetBody() *KnowledgeBaseResult
}

type GetKnowledgeBaseResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *KnowledgeBaseResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetKnowledgeBaseResponse) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseResponse) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetKnowledgeBaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetKnowledgeBaseResponse) GetBody() *KnowledgeBaseResult {
	return s.Body
}

func (s *GetKnowledgeBaseResponse) SetHeaders(v map[string]*string) *GetKnowledgeBaseResponse {
	s.Headers = v
	return s
}

func (s *GetKnowledgeBaseResponse) SetStatusCode(v int32) *GetKnowledgeBaseResponse {
	s.StatusCode = &v
	return s
}

func (s *GetKnowledgeBaseResponse) SetBody(v *KnowledgeBaseResult) *GetKnowledgeBaseResponse {
	s.Body = v
	return s
}

func (s *GetKnowledgeBaseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
