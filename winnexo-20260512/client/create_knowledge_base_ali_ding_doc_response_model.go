// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeBaseAliDingDocResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateKnowledgeBaseAliDingDocResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateKnowledgeBaseAliDingDocResponse
	GetStatusCode() *int32
	SetBody(v *CreateKnowledgeBaseAliDingDocResponseBody) *CreateKnowledgeBaseAliDingDocResponse
	GetBody() *CreateKnowledgeBaseAliDingDocResponseBody
}

type CreateKnowledgeBaseAliDingDocResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateKnowledgeBaseAliDingDocResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateKnowledgeBaseAliDingDocResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseAliDingDocResponse) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseAliDingDocResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateKnowledgeBaseAliDingDocResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateKnowledgeBaseAliDingDocResponse) GetBody() *CreateKnowledgeBaseAliDingDocResponseBody {
	return s.Body
}

func (s *CreateKnowledgeBaseAliDingDocResponse) SetHeaders(v map[string]*string) *CreateKnowledgeBaseAliDingDocResponse {
	s.Headers = v
	return s
}

func (s *CreateKnowledgeBaseAliDingDocResponse) SetStatusCode(v int32) *CreateKnowledgeBaseAliDingDocResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateKnowledgeBaseAliDingDocResponse) SetBody(v *CreateKnowledgeBaseAliDingDocResponseBody) *CreateKnowledgeBaseAliDingDocResponse {
	s.Body = v
	return s
}

func (s *CreateKnowledgeBaseAliDingDocResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
