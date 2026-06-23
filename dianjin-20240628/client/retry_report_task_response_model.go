// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryReportTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RetryReportTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RetryReportTaskResponse
	GetStatusCode() *int32
	SetBody(v *RetryReportTaskResponseBody) *RetryReportTaskResponse
	GetBody() *RetryReportTaskResponseBody
}

type RetryReportTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RetryReportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetryReportTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s RetryReportTaskResponse) GoString() string {
	return s.String()
}

func (s *RetryReportTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RetryReportTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RetryReportTaskResponse) GetBody() *RetryReportTaskResponseBody {
	return s.Body
}

func (s *RetryReportTaskResponse) SetHeaders(v map[string]*string) *RetryReportTaskResponse {
	s.Headers = v
	return s
}

func (s *RetryReportTaskResponse) SetStatusCode(v int32) *RetryReportTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *RetryReportTaskResponse) SetBody(v *RetryReportTaskResponseBody) *RetryReportTaskResponse {
	s.Body = v
	return s
}

func (s *RetryReportTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
