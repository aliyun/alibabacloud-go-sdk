// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewKnowledgeBaseSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PreviewKnowledgeBaseSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PreviewKnowledgeBaseSourceResponse
	GetStatusCode() *int32
	SetBody(v *PreviewKnowledgeBaseSourceResponseBody) *PreviewKnowledgeBaseSourceResponse
	GetBody() *PreviewKnowledgeBaseSourceResponseBody
}

type PreviewKnowledgeBaseSourceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PreviewKnowledgeBaseSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PreviewKnowledgeBaseSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s PreviewKnowledgeBaseSourceResponse) GoString() string {
	return s.String()
}

func (s *PreviewKnowledgeBaseSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PreviewKnowledgeBaseSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PreviewKnowledgeBaseSourceResponse) GetBody() *PreviewKnowledgeBaseSourceResponseBody {
	return s.Body
}

func (s *PreviewKnowledgeBaseSourceResponse) SetHeaders(v map[string]*string) *PreviewKnowledgeBaseSourceResponse {
	s.Headers = v
	return s
}

func (s *PreviewKnowledgeBaseSourceResponse) SetStatusCode(v int32) *PreviewKnowledgeBaseSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *PreviewKnowledgeBaseSourceResponse) SetBody(v *PreviewKnowledgeBaseSourceResponseBody) *PreviewKnowledgeBaseSourceResponse {
	s.Body = v
	return s
}

func (s *PreviewKnowledgeBaseSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
