// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeJobGroupExportTaskProgressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeJobGroupExportTaskProgressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeJobGroupExportTaskProgressResponse
	GetStatusCode() *int32
	SetBody(v *DescribeJobGroupExportTaskProgressResponseBody) *DescribeJobGroupExportTaskProgressResponse
	GetBody() *DescribeJobGroupExportTaskProgressResponseBody
}

type DescribeJobGroupExportTaskProgressResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeJobGroupExportTaskProgressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeJobGroupExportTaskProgressResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobGroupExportTaskProgressResponse) GoString() string {
	return s.String()
}

func (s *DescribeJobGroupExportTaskProgressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeJobGroupExportTaskProgressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeJobGroupExportTaskProgressResponse) GetBody() *DescribeJobGroupExportTaskProgressResponseBody {
	return s.Body
}

func (s *DescribeJobGroupExportTaskProgressResponse) SetHeaders(v map[string]*string) *DescribeJobGroupExportTaskProgressResponse {
	s.Headers = v
	return s
}

func (s *DescribeJobGroupExportTaskProgressResponse) SetStatusCode(v int32) *DescribeJobGroupExportTaskProgressResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeJobGroupExportTaskProgressResponse) SetBody(v *DescribeJobGroupExportTaskProgressResponseBody) *DescribeJobGroupExportTaskProgressResponse {
	s.Body = v
	return s
}

func (s *DescribeJobGroupExportTaskProgressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
