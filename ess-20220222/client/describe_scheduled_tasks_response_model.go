// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScheduledTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeScheduledTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeScheduledTasksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeScheduledTasksResponseBody) *DescribeScheduledTasksResponse
	GetBody() *DescribeScheduledTasksResponseBody
}

type DescribeScheduledTasksResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeScheduledTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeScheduledTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeScheduledTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeScheduledTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeScheduledTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeScheduledTasksResponse) GetBody() *DescribeScheduledTasksResponseBody {
	return s.Body
}

func (s *DescribeScheduledTasksResponse) SetHeaders(v map[string]*string) *DescribeScheduledTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeScheduledTasksResponse) SetStatusCode(v int32) *DescribeScheduledTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScheduledTasksResponse) SetBody(v *DescribeScheduledTasksResponseBody) *DescribeScheduledTasksResponse {
	s.Body = v
	return s
}

func (s *DescribeScheduledTasksResponse) Validate() error {
	return dara.Validate(s)
}
