// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParameterTimedScheduleTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeParameterTimedScheduleTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeParameterTimedScheduleTaskResponse
	GetStatusCode() *int32
	SetBody(v *DescribeParameterTimedScheduleTaskResponseBody) *DescribeParameterTimedScheduleTaskResponse
	GetBody() *DescribeParameterTimedScheduleTaskResponseBody
}

type DescribeParameterTimedScheduleTaskResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeParameterTimedScheduleTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeParameterTimedScheduleTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterTimedScheduleTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeParameterTimedScheduleTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeParameterTimedScheduleTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeParameterTimedScheduleTaskResponse) GetBody() *DescribeParameterTimedScheduleTaskResponseBody {
	return s.Body
}

func (s *DescribeParameterTimedScheduleTaskResponse) SetHeaders(v map[string]*string) *DescribeParameterTimedScheduleTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeParameterTimedScheduleTaskResponse) SetStatusCode(v int32) *DescribeParameterTimedScheduleTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeParameterTimedScheduleTaskResponse) SetBody(v *DescribeParameterTimedScheduleTaskResponseBody) *DescribeParameterTimedScheduleTaskResponse {
	s.Body = v
	return s
}

func (s *DescribeParameterTimedScheduleTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
