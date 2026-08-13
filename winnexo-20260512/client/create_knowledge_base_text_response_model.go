// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeBaseTextResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateKnowledgeBaseTextResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateKnowledgeBaseTextResponse
	GetStatusCode() *int32
	SetBody(v *CreateKnowledgeBaseTextResponseBody) *CreateKnowledgeBaseTextResponse
	GetBody() *CreateKnowledgeBaseTextResponseBody
}

type CreateKnowledgeBaseTextResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateKnowledgeBaseTextResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateKnowledgeBaseTextResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseTextResponse) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseTextResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateKnowledgeBaseTextResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateKnowledgeBaseTextResponse) GetBody() *CreateKnowledgeBaseTextResponseBody {
	return s.Body
}

func (s *CreateKnowledgeBaseTextResponse) SetHeaders(v map[string]*string) *CreateKnowledgeBaseTextResponse {
	s.Headers = v
	return s
}

func (s *CreateKnowledgeBaseTextResponse) SetStatusCode(v int32) *CreateKnowledgeBaseTextResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateKnowledgeBaseTextResponse) SetBody(v *CreateKnowledgeBaseTextResponseBody) *CreateKnowledgeBaseTextResponse {
	s.Body = v
	return s
}

func (s *CreateKnowledgeBaseTextResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
