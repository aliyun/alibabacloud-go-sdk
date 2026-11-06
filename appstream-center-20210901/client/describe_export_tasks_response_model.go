// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExportTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeExportTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeExportTasksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeExportTasksResponseBody) *DescribeExportTasksResponse
	GetBody() *DescribeExportTasksResponseBody
}

type DescribeExportTasksResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExportTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExportTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeExportTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeExportTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeExportTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeExportTasksResponse) GetBody() *DescribeExportTasksResponseBody {
	return s.Body
}

func (s *DescribeExportTasksResponse) SetHeaders(v map[string]*string) *DescribeExportTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeExportTasksResponse) SetStatusCode(v int32) *DescribeExportTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExportTasksResponse) SetBody(v *DescribeExportTasksResponseBody) *DescribeExportTasksResponse {
	s.Body = v
	return s
}

func (s *DescribeExportTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
