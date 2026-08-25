// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMultiModelKnowledgeBaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMultiModelKnowledgeBaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMultiModelKnowledgeBaseResponse
	GetStatusCode() *int32
	SetBody(v *CreateMultiModelKnowledgeBaseResponseBody) *CreateMultiModelKnowledgeBaseResponse
	GetBody() *CreateMultiModelKnowledgeBaseResponseBody
}

type CreateMultiModelKnowledgeBaseResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMultiModelKnowledgeBaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMultiModelKnowledgeBaseResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMultiModelKnowledgeBaseResponse) GoString() string {
	return s.String()
}

func (s *CreateMultiModelKnowledgeBaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMultiModelKnowledgeBaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMultiModelKnowledgeBaseResponse) GetBody() *CreateMultiModelKnowledgeBaseResponseBody {
	return s.Body
}

func (s *CreateMultiModelKnowledgeBaseResponse) SetHeaders(v map[string]*string) *CreateMultiModelKnowledgeBaseResponse {
	s.Headers = v
	return s
}

func (s *CreateMultiModelKnowledgeBaseResponse) SetStatusCode(v int32) *CreateMultiModelKnowledgeBaseResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMultiModelKnowledgeBaseResponse) SetBody(v *CreateMultiModelKnowledgeBaseResponseBody) *CreateMultiModelKnowledgeBaseResponse {
	s.Body = v
	return s
}

func (s *CreateMultiModelKnowledgeBaseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
