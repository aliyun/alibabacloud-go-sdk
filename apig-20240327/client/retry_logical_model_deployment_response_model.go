// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryLogicalModelDeploymentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RetryLogicalModelDeploymentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RetryLogicalModelDeploymentResponse
	GetStatusCode() *int32
	SetBody(v *RetryLogicalModelDeploymentResponseBody) *RetryLogicalModelDeploymentResponse
	GetBody() *RetryLogicalModelDeploymentResponseBody
}

type RetryLogicalModelDeploymentResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RetryLogicalModelDeploymentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetryLogicalModelDeploymentResponse) String() string {
	return dara.Prettify(s)
}

func (s RetryLogicalModelDeploymentResponse) GoString() string {
	return s.String()
}

func (s *RetryLogicalModelDeploymentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RetryLogicalModelDeploymentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RetryLogicalModelDeploymentResponse) GetBody() *RetryLogicalModelDeploymentResponseBody {
	return s.Body
}

func (s *RetryLogicalModelDeploymentResponse) SetHeaders(v map[string]*string) *RetryLogicalModelDeploymentResponse {
	s.Headers = v
	return s
}

func (s *RetryLogicalModelDeploymentResponse) SetStatusCode(v int32) *RetryLogicalModelDeploymentResponse {
	s.StatusCode = &v
	return s
}

func (s *RetryLogicalModelDeploymentResponse) SetBody(v *RetryLogicalModelDeploymentResponseBody) *RetryLogicalModelDeploymentResponse {
	s.Body = v
	return s
}

func (s *RetryLogicalModelDeploymentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
