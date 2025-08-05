// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryMmsJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RetryMmsJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RetryMmsJobResponse
	GetStatusCode() *int32
	SetBody(v *RetryMmsJobResponseBody) *RetryMmsJobResponse
	GetBody() *RetryMmsJobResponseBody
}

type RetryMmsJobResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RetryMmsJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetryMmsJobResponse) String() string {
	return dara.Prettify(s)
}

func (s RetryMmsJobResponse) GoString() string {
	return s.String()
}

func (s *RetryMmsJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RetryMmsJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RetryMmsJobResponse) GetBody() *RetryMmsJobResponseBody {
	return s.Body
}

func (s *RetryMmsJobResponse) SetHeaders(v map[string]*string) *RetryMmsJobResponse {
	s.Headers = v
	return s
}

func (s *RetryMmsJobResponse) SetStatusCode(v int32) *RetryMmsJobResponse {
	s.StatusCode = &v
	return s
}

func (s *RetryMmsJobResponse) SetBody(v *RetryMmsJobResponseBody) *RetryMmsJobResponse {
	s.Body = v
	return s
}

func (s *RetryMmsJobResponse) Validate() error {
	return dara.Validate(s)
}
