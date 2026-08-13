// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceKnowledgeBaseSourceFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReplaceKnowledgeBaseSourceFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReplaceKnowledgeBaseSourceFileResponse
	GetStatusCode() *int32
	SetBody(v *ReplaceKnowledgeBaseSourceFileResponseBody) *ReplaceKnowledgeBaseSourceFileResponse
	GetBody() *ReplaceKnowledgeBaseSourceFileResponseBody
}

type ReplaceKnowledgeBaseSourceFileResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReplaceKnowledgeBaseSourceFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReplaceKnowledgeBaseSourceFileResponse) String() string {
	return dara.Prettify(s)
}

func (s ReplaceKnowledgeBaseSourceFileResponse) GoString() string {
	return s.String()
}

func (s *ReplaceKnowledgeBaseSourceFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReplaceKnowledgeBaseSourceFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReplaceKnowledgeBaseSourceFileResponse) GetBody() *ReplaceKnowledgeBaseSourceFileResponseBody {
	return s.Body
}

func (s *ReplaceKnowledgeBaseSourceFileResponse) SetHeaders(v map[string]*string) *ReplaceKnowledgeBaseSourceFileResponse {
	s.Headers = v
	return s
}

func (s *ReplaceKnowledgeBaseSourceFileResponse) SetStatusCode(v int32) *ReplaceKnowledgeBaseSourceFileResponse {
	s.StatusCode = &v
	return s
}

func (s *ReplaceKnowledgeBaseSourceFileResponse) SetBody(v *ReplaceKnowledgeBaseSourceFileResponseBody) *ReplaceKnowledgeBaseSourceFileResponse {
	s.Body = v
	return s
}

func (s *ReplaceKnowledgeBaseSourceFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
