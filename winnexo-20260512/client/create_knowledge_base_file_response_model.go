// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeBaseFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateKnowledgeBaseFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateKnowledgeBaseFileResponse
	GetStatusCode() *int32
	SetBody(v *CreateKnowledgeBaseFileResponseBody) *CreateKnowledgeBaseFileResponse
	GetBody() *CreateKnowledgeBaseFileResponseBody
}

type CreateKnowledgeBaseFileResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateKnowledgeBaseFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateKnowledgeBaseFileResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseFileResponse) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateKnowledgeBaseFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateKnowledgeBaseFileResponse) GetBody() *CreateKnowledgeBaseFileResponseBody {
	return s.Body
}

func (s *CreateKnowledgeBaseFileResponse) SetHeaders(v map[string]*string) *CreateKnowledgeBaseFileResponse {
	s.Headers = v
	return s
}

func (s *CreateKnowledgeBaseFileResponse) SetStatusCode(v int32) *CreateKnowledgeBaseFileResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateKnowledgeBaseFileResponse) SetBody(v *CreateKnowledgeBaseFileResponseBody) *CreateKnowledgeBaseFileResponse {
	s.Body = v
	return s
}

func (s *CreateKnowledgeBaseFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
