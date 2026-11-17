// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryKnowledgeBaseBuildTaskInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RetryKnowledgeBaseBuildTaskInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RetryKnowledgeBaseBuildTaskInstanceResponse
	GetStatusCode() *int32
	SetBody(v *RetryKnowledgeBaseBuildTaskInstanceResponseBody) *RetryKnowledgeBaseBuildTaskInstanceResponse
	GetBody() *RetryKnowledgeBaseBuildTaskInstanceResponseBody
}

type RetryKnowledgeBaseBuildTaskInstanceResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RetryKnowledgeBaseBuildTaskInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetryKnowledgeBaseBuildTaskInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s RetryKnowledgeBaseBuildTaskInstanceResponse) GoString() string {
	return s.String()
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponse) GetBody() *RetryKnowledgeBaseBuildTaskInstanceResponseBody {
	return s.Body
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponse) SetHeaders(v map[string]*string) *RetryKnowledgeBaseBuildTaskInstanceResponse {
	s.Headers = v
	return s
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponse) SetStatusCode(v int32) *RetryKnowledgeBaseBuildTaskInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponse) SetBody(v *RetryKnowledgeBaseBuildTaskInstanceResponseBody) *RetryKnowledgeBaseBuildTaskInstanceResponse {
	s.Body = v
	return s
}

func (s *RetryKnowledgeBaseBuildTaskInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
