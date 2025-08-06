// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodUserUsageDetailDataExportTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodUserUsageDetailDataExportTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodUserUsageDetailDataExportTaskResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodUserUsageDetailDataExportTaskResponseBody) *DescribeVodUserUsageDetailDataExportTaskResponse
	GetBody() *DescribeVodUserUsageDetailDataExportTaskResponseBody
}

type DescribeVodUserUsageDetailDataExportTaskResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodUserUsageDetailDataExportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodUserUsageDetailDataExportTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodUserUsageDetailDataExportTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodUserUsageDetailDataExportTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodUserUsageDetailDataExportTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodUserUsageDetailDataExportTaskResponse) GetBody() *DescribeVodUserUsageDetailDataExportTaskResponseBody {
	return s.Body
}

func (s *DescribeVodUserUsageDetailDataExportTaskResponse) SetHeaders(v map[string]*string) *DescribeVodUserUsageDetailDataExportTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodUserUsageDetailDataExportTaskResponse) SetStatusCode(v int32) *DescribeVodUserUsageDetailDataExportTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodUserUsageDetailDataExportTaskResponse) SetBody(v *DescribeVodUserUsageDetailDataExportTaskResponseBody) *DescribeVodUserUsageDetailDataExportTaskResponse {
	s.Body = v
	return s
}

func (s *DescribeVodUserUsageDetailDataExportTaskResponse) Validate() error {
	return dara.Validate(s)
}
