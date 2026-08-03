// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryInspectionTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RetryInspectionTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RetryInspectionTaskResponse
	GetStatusCode() *int32
	SetBody(v *RetryInspectionTaskResponseBody) *RetryInspectionTaskResponse
	GetBody() *RetryInspectionTaskResponseBody
}

type RetryInspectionTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RetryInspectionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetryInspectionTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s RetryInspectionTaskResponse) GoString() string {
	return s.String()
}

func (s *RetryInspectionTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RetryInspectionTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RetryInspectionTaskResponse) GetBody() *RetryInspectionTaskResponseBody {
	return s.Body
}

func (s *RetryInspectionTaskResponse) SetHeaders(v map[string]*string) *RetryInspectionTaskResponse {
	s.Headers = v
	return s
}

func (s *RetryInspectionTaskResponse) SetStatusCode(v int32) *RetryInspectionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *RetryInspectionTaskResponse) SetBody(v *RetryInspectionTaskResponseBody) *RetryInspectionTaskResponse {
	s.Body = v
	return s
}

func (s *RetryInspectionTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
