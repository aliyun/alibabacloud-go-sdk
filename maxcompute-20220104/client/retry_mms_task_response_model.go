// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryMmsTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RetryMmsTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RetryMmsTaskResponse
	GetStatusCode() *int32
	SetBody(v *RetryMmsTaskResponseBody) *RetryMmsTaskResponse
	GetBody() *RetryMmsTaskResponseBody
}

type RetryMmsTaskResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RetryMmsTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetryMmsTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s RetryMmsTaskResponse) GoString() string {
	return s.String()
}

func (s *RetryMmsTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RetryMmsTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RetryMmsTaskResponse) GetBody() *RetryMmsTaskResponseBody {
	return s.Body
}

func (s *RetryMmsTaskResponse) SetHeaders(v map[string]*string) *RetryMmsTaskResponse {
	s.Headers = v
	return s
}

func (s *RetryMmsTaskResponse) SetStatusCode(v int32) *RetryMmsTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *RetryMmsTaskResponse) SetBody(v *RetryMmsTaskResponseBody) *RetryMmsTaskResponse {
	s.Body = v
	return s
}

func (s *RetryMmsTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
