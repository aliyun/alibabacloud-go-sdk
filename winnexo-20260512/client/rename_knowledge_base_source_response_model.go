// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenameKnowledgeBaseSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RenameKnowledgeBaseSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RenameKnowledgeBaseSourceResponse
	GetStatusCode() *int32
	SetBody(v *RenameKnowledgeBaseSourceResponseBody) *RenameKnowledgeBaseSourceResponse
	GetBody() *RenameKnowledgeBaseSourceResponseBody
}

type RenameKnowledgeBaseSourceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenameKnowledgeBaseSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenameKnowledgeBaseSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s RenameKnowledgeBaseSourceResponse) GoString() string {
	return s.String()
}

func (s *RenameKnowledgeBaseSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RenameKnowledgeBaseSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RenameKnowledgeBaseSourceResponse) GetBody() *RenameKnowledgeBaseSourceResponseBody {
	return s.Body
}

func (s *RenameKnowledgeBaseSourceResponse) SetHeaders(v map[string]*string) *RenameKnowledgeBaseSourceResponse {
	s.Headers = v
	return s
}

func (s *RenameKnowledgeBaseSourceResponse) SetStatusCode(v int32) *RenameKnowledgeBaseSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *RenameKnowledgeBaseSourceResponse) SetBody(v *RenameKnowledgeBaseSourceResponseBody) *RenameKnowledgeBaseSourceResponse {
	s.Body = v
	return s
}

func (s *RenameKnowledgeBaseSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
