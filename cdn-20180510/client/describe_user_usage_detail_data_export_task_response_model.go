// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserUsageDetailDataExportTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserUsageDetailDataExportTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserUsageDetailDataExportTaskResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserUsageDetailDataExportTaskResponseBody) *DescribeUserUsageDetailDataExportTaskResponse
	GetBody() *DescribeUserUsageDetailDataExportTaskResponseBody
}

type DescribeUserUsageDetailDataExportTaskResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserUsageDetailDataExportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserUsageDetailDataExportTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserUsageDetailDataExportTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserUsageDetailDataExportTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserUsageDetailDataExportTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserUsageDetailDataExportTaskResponse) GetBody() *DescribeUserUsageDetailDataExportTaskResponseBody {
	return s.Body
}

func (s *DescribeUserUsageDetailDataExportTaskResponse) SetHeaders(v map[string]*string) *DescribeUserUsageDetailDataExportTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserUsageDetailDataExportTaskResponse) SetStatusCode(v int32) *DescribeUserUsageDetailDataExportTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserUsageDetailDataExportTaskResponse) SetBody(v *DescribeUserUsageDetailDataExportTaskResponseBody) *DescribeUserUsageDetailDataExportTaskResponse {
	s.Body = v
	return s
}

func (s *DescribeUserUsageDetailDataExportTaskResponse) Validate() error {
	return dara.Validate(s)
}
