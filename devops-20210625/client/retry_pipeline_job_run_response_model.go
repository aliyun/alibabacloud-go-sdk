// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryPipelineJobRunResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RetryPipelineJobRunResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RetryPipelineJobRunResponse
	GetStatusCode() *int32
	SetBody(v *RetryPipelineJobRunResponseBody) *RetryPipelineJobRunResponse
	GetBody() *RetryPipelineJobRunResponseBody
}

type RetryPipelineJobRunResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RetryPipelineJobRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetryPipelineJobRunResponse) String() string {
	return dara.Prettify(s)
}

func (s RetryPipelineJobRunResponse) GoString() string {
	return s.String()
}

func (s *RetryPipelineJobRunResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RetryPipelineJobRunResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RetryPipelineJobRunResponse) GetBody() *RetryPipelineJobRunResponseBody {
	return s.Body
}

func (s *RetryPipelineJobRunResponse) SetHeaders(v map[string]*string) *RetryPipelineJobRunResponse {
	s.Headers = v
	return s
}

func (s *RetryPipelineJobRunResponse) SetStatusCode(v int32) *RetryPipelineJobRunResponse {
	s.StatusCode = &v
	return s
}

func (s *RetryPipelineJobRunResponse) SetBody(v *RetryPipelineJobRunResponseBody) *RetryPipelineJobRunResponse {
	s.Body = v
	return s
}

func (s *RetryPipelineJobRunResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
