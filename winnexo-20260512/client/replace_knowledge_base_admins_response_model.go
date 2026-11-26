// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceKnowledgeBaseAdminsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReplaceKnowledgeBaseAdminsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReplaceKnowledgeBaseAdminsResponse
	GetStatusCode() *int32
	SetBody(v *ReplaceKnowledgeBaseAdminsResponseBody) *ReplaceKnowledgeBaseAdminsResponse
	GetBody() *ReplaceKnowledgeBaseAdminsResponseBody
}

type ReplaceKnowledgeBaseAdminsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReplaceKnowledgeBaseAdminsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReplaceKnowledgeBaseAdminsResponse) String() string {
	return dara.Prettify(s)
}

func (s ReplaceKnowledgeBaseAdminsResponse) GoString() string {
	return s.String()
}

func (s *ReplaceKnowledgeBaseAdminsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReplaceKnowledgeBaseAdminsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReplaceKnowledgeBaseAdminsResponse) GetBody() *ReplaceKnowledgeBaseAdminsResponseBody {
	return s.Body
}

func (s *ReplaceKnowledgeBaseAdminsResponse) SetHeaders(v map[string]*string) *ReplaceKnowledgeBaseAdminsResponse {
	s.Headers = v
	return s
}

func (s *ReplaceKnowledgeBaseAdminsResponse) SetStatusCode(v int32) *ReplaceKnowledgeBaseAdminsResponse {
	s.StatusCode = &v
	return s
}

func (s *ReplaceKnowledgeBaseAdminsResponse) SetBody(v *ReplaceKnowledgeBaseAdminsResponseBody) *ReplaceKnowledgeBaseAdminsResponse {
	s.Body = v
	return s
}

func (s *ReplaceKnowledgeBaseAdminsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
