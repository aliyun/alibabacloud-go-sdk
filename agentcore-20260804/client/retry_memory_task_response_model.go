// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryMemoryTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RetryMemoryTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RetryMemoryTaskResponse
	GetStatusCode() *int32
	SetBody(v *RetryMemoryTaskResponseBody) *RetryMemoryTaskResponse
	GetBody() *RetryMemoryTaskResponseBody
}

type RetryMemoryTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RetryMemoryTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetryMemoryTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s RetryMemoryTaskResponse) GoString() string {
	return s.String()
}

func (s *RetryMemoryTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RetryMemoryTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RetryMemoryTaskResponse) GetBody() *RetryMemoryTaskResponseBody {
	return s.Body
}

func (s *RetryMemoryTaskResponse) SetHeaders(v map[string]*string) *RetryMemoryTaskResponse {
	s.Headers = v
	return s
}

func (s *RetryMemoryTaskResponse) SetStatusCode(v int32) *RetryMemoryTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *RetryMemoryTaskResponse) SetBody(v *RetryMemoryTaskResponseBody) *RetryMemoryTaskResponse {
	s.Body = v
	return s
}

func (s *RetryMemoryTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
