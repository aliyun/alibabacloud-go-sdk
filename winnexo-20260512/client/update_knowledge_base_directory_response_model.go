// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKnowledgeBaseDirectoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateKnowledgeBaseDirectoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateKnowledgeBaseDirectoryResponse
	GetStatusCode() *int32
	SetBody(v *UpdateKnowledgeBaseDirectoryResponseBody) *UpdateKnowledgeBaseDirectoryResponse
	GetBody() *UpdateKnowledgeBaseDirectoryResponseBody
}

type UpdateKnowledgeBaseDirectoryResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateKnowledgeBaseDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateKnowledgeBaseDirectoryResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseDirectoryResponse) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseDirectoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateKnowledgeBaseDirectoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateKnowledgeBaseDirectoryResponse) GetBody() *UpdateKnowledgeBaseDirectoryResponseBody {
	return s.Body
}

func (s *UpdateKnowledgeBaseDirectoryResponse) SetHeaders(v map[string]*string) *UpdateKnowledgeBaseDirectoryResponse {
	s.Headers = v
	return s
}

func (s *UpdateKnowledgeBaseDirectoryResponse) SetStatusCode(v int32) *UpdateKnowledgeBaseDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateKnowledgeBaseDirectoryResponse) SetBody(v *UpdateKnowledgeBaseDirectoryResponseBody) *UpdateKnowledgeBaseDirectoryResponse {
	s.Body = v
	return s
}

func (s *UpdateKnowledgeBaseDirectoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
