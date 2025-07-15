// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEcdReportTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEcdReportTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEcdReportTasksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEcdReportTasksResponseBody) *DescribeEcdReportTasksResponse
	GetBody() *DescribeEcdReportTasksResponseBody
}

type DescribeEcdReportTasksResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEcdReportTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEcdReportTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEcdReportTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeEcdReportTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEcdReportTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEcdReportTasksResponse) GetBody() *DescribeEcdReportTasksResponseBody {
	return s.Body
}

func (s *DescribeEcdReportTasksResponse) SetHeaders(v map[string]*string) *DescribeEcdReportTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeEcdReportTasksResponse) SetStatusCode(v int32) *DescribeEcdReportTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEcdReportTasksResponse) SetBody(v *DescribeEcdReportTasksResponseBody) *DescribeEcdReportTasksResponse {
	s.Body = v
	return s
}

func (s *DescribeEcdReportTasksResponse) Validate() error {
	return dara.Validate(s)
}
