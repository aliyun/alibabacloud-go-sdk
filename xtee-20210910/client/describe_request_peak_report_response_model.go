// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRequestPeakReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRequestPeakReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRequestPeakReportResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRequestPeakReportResponseBody) *DescribeRequestPeakReportResponse
	GetBody() *DescribeRequestPeakReportResponseBody
}

type DescribeRequestPeakReportResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRequestPeakReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRequestPeakReportResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRequestPeakReportResponse) GoString() string {
	return s.String()
}

func (s *DescribeRequestPeakReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRequestPeakReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRequestPeakReportResponse) GetBody() *DescribeRequestPeakReportResponseBody {
	return s.Body
}

func (s *DescribeRequestPeakReportResponse) SetHeaders(v map[string]*string) *DescribeRequestPeakReportResponse {
	s.Headers = v
	return s
}

func (s *DescribeRequestPeakReportResponse) SetStatusCode(v int32) *DescribeRequestPeakReportResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRequestPeakReportResponse) SetBody(v *DescribeRequestPeakReportResponseBody) *DescribeRequestPeakReportResponse {
	s.Body = v
	return s
}

func (s *DescribeRequestPeakReportResponse) Validate() error {
	return dara.Validate(s)
}
