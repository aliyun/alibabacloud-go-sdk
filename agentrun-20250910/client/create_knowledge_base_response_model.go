// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeBaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateKnowledgeBaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateKnowledgeBaseResponse
	GetStatusCode() *int32
	SetBody(v *KnowledgeBaseResult) *CreateKnowledgeBaseResponse
	GetBody() *KnowledgeBaseResult
}

type CreateKnowledgeBaseResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *KnowledgeBaseResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateKnowledgeBaseResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseResponse) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateKnowledgeBaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateKnowledgeBaseResponse) GetBody() *KnowledgeBaseResult {
	return s.Body
}

func (s *CreateKnowledgeBaseResponse) SetHeaders(v map[string]*string) *CreateKnowledgeBaseResponse {
	s.Headers = v
	return s
}

func (s *CreateKnowledgeBaseResponse) SetStatusCode(v int32) *CreateKnowledgeBaseResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateKnowledgeBaseResponse) SetBody(v *KnowledgeBaseResult) *CreateKnowledgeBaseResponse {
	s.Body = v
	return s
}

func (s *CreateKnowledgeBaseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
