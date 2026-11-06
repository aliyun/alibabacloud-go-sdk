// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInspectionReportListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInspectionReportListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInspectionReportListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInspectionReportListResponseBody) *DescribeInspectionReportListResponse
	GetBody() *DescribeInspectionReportListResponseBody
}

type DescribeInspectionReportListResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInspectionReportListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInspectionReportListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInspectionReportListResponse) GoString() string {
	return s.String()
}

func (s *DescribeInspectionReportListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInspectionReportListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInspectionReportListResponse) GetBody() *DescribeInspectionReportListResponseBody {
	return s.Body
}

func (s *DescribeInspectionReportListResponse) SetHeaders(v map[string]*string) *DescribeInspectionReportListResponse {
	s.Headers = v
	return s
}

func (s *DescribeInspectionReportListResponse) SetStatusCode(v int32) *DescribeInspectionReportListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInspectionReportListResponse) SetBody(v *DescribeInspectionReportListResponseBody) *DescribeInspectionReportListResponse {
	s.Body = v
	return s
}

func (s *DescribeInspectionReportListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
