// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryKnowledgeBaseFailedSourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RetryKnowledgeBaseFailedSourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RetryKnowledgeBaseFailedSourcesResponse
	GetStatusCode() *int32
	SetBody(v *RetryKnowledgeBaseFailedSourcesResponseBody) *RetryKnowledgeBaseFailedSourcesResponse
	GetBody() *RetryKnowledgeBaseFailedSourcesResponseBody
}

type RetryKnowledgeBaseFailedSourcesResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RetryKnowledgeBaseFailedSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetryKnowledgeBaseFailedSourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s RetryKnowledgeBaseFailedSourcesResponse) GoString() string {
	return s.String()
}

func (s *RetryKnowledgeBaseFailedSourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RetryKnowledgeBaseFailedSourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RetryKnowledgeBaseFailedSourcesResponse) GetBody() *RetryKnowledgeBaseFailedSourcesResponseBody {
	return s.Body
}

func (s *RetryKnowledgeBaseFailedSourcesResponse) SetHeaders(v map[string]*string) *RetryKnowledgeBaseFailedSourcesResponse {
	s.Headers = v
	return s
}

func (s *RetryKnowledgeBaseFailedSourcesResponse) SetStatusCode(v int32) *RetryKnowledgeBaseFailedSourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *RetryKnowledgeBaseFailedSourcesResponse) SetBody(v *RetryKnowledgeBaseFailedSourcesResponseBody) *RetryKnowledgeBaseFailedSourcesResponse {
	s.Body = v
	return s
}

func (s *RetryKnowledgeBaseFailedSourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
