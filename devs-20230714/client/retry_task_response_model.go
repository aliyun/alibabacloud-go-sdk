// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RetryTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RetryTaskResponse
	GetStatusCode() *int32
	SetBody(v *Task) *RetryTaskResponse
	GetBody() *Task
}

type RetryTaskResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Task              `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetryTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s RetryTaskResponse) GoString() string {
	return s.String()
}

func (s *RetryTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RetryTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RetryTaskResponse) GetBody() *Task {
	return s.Body
}

func (s *RetryTaskResponse) SetHeaders(v map[string]*string) *RetryTaskResponse {
	s.Headers = v
	return s
}

func (s *RetryTaskResponse) SetStatusCode(v int32) *RetryTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *RetryTaskResponse) SetBody(v *Task) *RetryTaskResponse {
	s.Body = v
	return s
}

func (s *RetryTaskResponse) Validate() error {
	return dara.Validate(s)
}
