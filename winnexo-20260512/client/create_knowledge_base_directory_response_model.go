// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeBaseDirectoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateKnowledgeBaseDirectoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateKnowledgeBaseDirectoryResponse
	GetStatusCode() *int32
	SetBody(v *CreateKnowledgeBaseDirectoryResponseBody) *CreateKnowledgeBaseDirectoryResponse
	GetBody() *CreateKnowledgeBaseDirectoryResponseBody
}

type CreateKnowledgeBaseDirectoryResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateKnowledgeBaseDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateKnowledgeBaseDirectoryResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseDirectoryResponse) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseDirectoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateKnowledgeBaseDirectoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateKnowledgeBaseDirectoryResponse) GetBody() *CreateKnowledgeBaseDirectoryResponseBody {
	return s.Body
}

func (s *CreateKnowledgeBaseDirectoryResponse) SetHeaders(v map[string]*string) *CreateKnowledgeBaseDirectoryResponse {
	s.Headers = v
	return s
}

func (s *CreateKnowledgeBaseDirectoryResponse) SetStatusCode(v int32) *CreateKnowledgeBaseDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateKnowledgeBaseDirectoryResponse) SetBody(v *CreateKnowledgeBaseDirectoryResponseBody) *CreateKnowledgeBaseDirectoryResponse {
	s.Body = v
	return s
}

func (s *CreateKnowledgeBaseDirectoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
