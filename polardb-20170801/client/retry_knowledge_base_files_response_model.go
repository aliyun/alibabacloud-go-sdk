// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryKnowledgeBaseFilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RetryKnowledgeBaseFilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RetryKnowledgeBaseFilesResponse
	GetStatusCode() *int32
	SetBody(v *RetryKnowledgeBaseFilesResponseBody) *RetryKnowledgeBaseFilesResponse
	GetBody() *RetryKnowledgeBaseFilesResponseBody
}

type RetryKnowledgeBaseFilesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RetryKnowledgeBaseFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetryKnowledgeBaseFilesResponse) String() string {
	return dara.Prettify(s)
}

func (s RetryKnowledgeBaseFilesResponse) GoString() string {
	return s.String()
}

func (s *RetryKnowledgeBaseFilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RetryKnowledgeBaseFilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RetryKnowledgeBaseFilesResponse) GetBody() *RetryKnowledgeBaseFilesResponseBody {
	return s.Body
}

func (s *RetryKnowledgeBaseFilesResponse) SetHeaders(v map[string]*string) *RetryKnowledgeBaseFilesResponse {
	s.Headers = v
	return s
}

func (s *RetryKnowledgeBaseFilesResponse) SetStatusCode(v int32) *RetryKnowledgeBaseFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *RetryKnowledgeBaseFilesResponse) SetBody(v *RetryKnowledgeBaseFilesResponseBody) *RetryKnowledgeBaseFilesResponse {
	s.Body = v
	return s
}

func (s *RetryKnowledgeBaseFilesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
