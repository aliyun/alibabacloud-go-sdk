// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKnowledgeBaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateKnowledgeBaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateKnowledgeBaseResponse
	GetStatusCode() *int32
	SetBody(v *KnowledgeBaseResult) *UpdateKnowledgeBaseResponse
	GetBody() *KnowledgeBaseResult
}

type UpdateKnowledgeBaseResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *KnowledgeBaseResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateKnowledgeBaseResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseResponse) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateKnowledgeBaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateKnowledgeBaseResponse) GetBody() *KnowledgeBaseResult {
	return s.Body
}

func (s *UpdateKnowledgeBaseResponse) SetHeaders(v map[string]*string) *UpdateKnowledgeBaseResponse {
	s.Headers = v
	return s
}

func (s *UpdateKnowledgeBaseResponse) SetStatusCode(v int32) *UpdateKnowledgeBaseResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateKnowledgeBaseResponse) SetBody(v *KnowledgeBaseResult) *UpdateKnowledgeBaseResponse {
	s.Body = v
	return s
}

func (s *UpdateKnowledgeBaseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
