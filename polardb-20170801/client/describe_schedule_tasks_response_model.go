// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScheduleTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeScheduleTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeScheduleTasksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeScheduleTasksResponseBody) *DescribeScheduleTasksResponse
	GetBody() *DescribeScheduleTasksResponseBody
}

type DescribeScheduleTasksResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeScheduleTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeScheduleTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeScheduleTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeScheduleTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeScheduleTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeScheduleTasksResponse) GetBody() *DescribeScheduleTasksResponseBody {
	return s.Body
}

func (s *DescribeScheduleTasksResponse) SetHeaders(v map[string]*string) *DescribeScheduleTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeScheduleTasksResponse) SetStatusCode(v int32) *DescribeScheduleTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScheduleTasksResponse) SetBody(v *DescribeScheduleTasksResponseBody) *DescribeScheduleTasksResponse {
	s.Body = v
	return s
}

func (s *DescribeScheduleTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
