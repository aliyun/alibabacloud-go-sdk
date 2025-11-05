// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserUsageDataExportTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserUsageDataExportTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserUsageDataExportTaskResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserUsageDataExportTaskResponseBody) *DescribeUserUsageDataExportTaskResponse
	GetBody() *DescribeUserUsageDataExportTaskResponseBody
}

type DescribeUserUsageDataExportTaskResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserUsageDataExportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserUsageDataExportTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserUsageDataExportTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserUsageDataExportTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserUsageDataExportTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserUsageDataExportTaskResponse) GetBody() *DescribeUserUsageDataExportTaskResponseBody {
	return s.Body
}

func (s *DescribeUserUsageDataExportTaskResponse) SetHeaders(v map[string]*string) *DescribeUserUsageDataExportTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserUsageDataExportTaskResponse) SetStatusCode(v int32) *DescribeUserUsageDataExportTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserUsageDataExportTaskResponse) SetBody(v *DescribeUserUsageDataExportTaskResponseBody) *DescribeUserUsageDataExportTaskResponse {
	s.Body = v
	return s
}

func (s *DescribeUserUsageDataExportTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
