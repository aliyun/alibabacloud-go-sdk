// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNisInspectionReportSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNisInspectionReportSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNisInspectionReportSummaryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNisInspectionReportSummaryResponseBody) *DescribeNisInspectionReportSummaryResponse
	GetBody() *DescribeNisInspectionReportSummaryResponseBody
}

type DescribeNisInspectionReportSummaryResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNisInspectionReportSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNisInspectionReportSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNisInspectionReportSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeNisInspectionReportSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNisInspectionReportSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNisInspectionReportSummaryResponse) GetBody() *DescribeNisInspectionReportSummaryResponseBody {
	return s.Body
}

func (s *DescribeNisInspectionReportSummaryResponse) SetHeaders(v map[string]*string) *DescribeNisInspectionReportSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeNisInspectionReportSummaryResponse) SetStatusCode(v int32) *DescribeNisInspectionReportSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNisInspectionReportSummaryResponse) SetBody(v *DescribeNisInspectionReportSummaryResponseBody) *DescribeNisInspectionReportSummaryResponse {
	s.Body = v
	return s
}

func (s *DescribeNisInspectionReportSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
