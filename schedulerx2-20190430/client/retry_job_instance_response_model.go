// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryJobInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RetryJobInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RetryJobInstanceResponse
	GetStatusCode() *int32
	SetBody(v *RetryJobInstanceResponseBody) *RetryJobInstanceResponse
	GetBody() *RetryJobInstanceResponseBody
}

type RetryJobInstanceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RetryJobInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetryJobInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s RetryJobInstanceResponse) GoString() string {
	return s.String()
}

func (s *RetryJobInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RetryJobInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RetryJobInstanceResponse) GetBody() *RetryJobInstanceResponseBody {
	return s.Body
}

func (s *RetryJobInstanceResponse) SetHeaders(v map[string]*string) *RetryJobInstanceResponse {
	s.Headers = v
	return s
}

func (s *RetryJobInstanceResponse) SetStatusCode(v int32) *RetryJobInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RetryJobInstanceResponse) SetBody(v *RetryJobInstanceResponseBody) *RetryJobInstanceResponse {
	s.Body = v
	return s
}

func (s *RetryJobInstanceResponse) Validate() error {
	return dara.Validate(s)
}
