// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryAgentlessTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RetryAgentlessTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RetryAgentlessTaskResponse
	GetStatusCode() *int32
	SetBody(v *RetryAgentlessTaskResponseBody) *RetryAgentlessTaskResponse
	GetBody() *RetryAgentlessTaskResponseBody
}

type RetryAgentlessTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RetryAgentlessTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetryAgentlessTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s RetryAgentlessTaskResponse) GoString() string {
	return s.String()
}

func (s *RetryAgentlessTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RetryAgentlessTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RetryAgentlessTaskResponse) GetBody() *RetryAgentlessTaskResponseBody {
	return s.Body
}

func (s *RetryAgentlessTaskResponse) SetHeaders(v map[string]*string) *RetryAgentlessTaskResponse {
	s.Headers = v
	return s
}

func (s *RetryAgentlessTaskResponse) SetStatusCode(v int32) *RetryAgentlessTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *RetryAgentlessTaskResponse) SetBody(v *RetryAgentlessTaskResponseBody) *RetryAgentlessTaskResponse {
	s.Body = v
	return s
}

func (s *RetryAgentlessTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
