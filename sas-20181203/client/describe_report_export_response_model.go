// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeReportExportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeReportExportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeReportExportResponse
	GetStatusCode() *int32
	SetBody(v *DescribeReportExportResponseBody) *DescribeReportExportResponse
	GetBody() *DescribeReportExportResponseBody
}

type DescribeReportExportResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeReportExportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeReportExportResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeReportExportResponse) GoString() string {
	return s.String()
}

func (s *DescribeReportExportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeReportExportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeReportExportResponse) GetBody() *DescribeReportExportResponseBody {
	return s.Body
}

func (s *DescribeReportExportResponse) SetHeaders(v map[string]*string) *DescribeReportExportResponse {
	s.Headers = v
	return s
}

func (s *DescribeReportExportResponse) SetStatusCode(v int32) *DescribeReportExportResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeReportExportResponse) SetBody(v *DescribeReportExportResponseBody) *DescribeReportExportResponse {
	s.Body = v
	return s
}

func (s *DescribeReportExportResponse) Validate() error {
	return dara.Validate(s)
}
