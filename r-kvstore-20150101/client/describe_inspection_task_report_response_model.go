// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInspectionTaskReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInspectionTaskReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInspectionTaskReportResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInspectionTaskReportResponseBody) *DescribeInspectionTaskReportResponse
	GetBody() *DescribeInspectionTaskReportResponseBody
}

type DescribeInspectionTaskReportResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInspectionTaskReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInspectionTaskReportResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInspectionTaskReportResponse) GoString() string {
	return s.String()
}

func (s *DescribeInspectionTaskReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInspectionTaskReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInspectionTaskReportResponse) GetBody() *DescribeInspectionTaskReportResponseBody {
	return s.Body
}

func (s *DescribeInspectionTaskReportResponse) SetHeaders(v map[string]*string) *DescribeInspectionTaskReportResponse {
	s.Headers = v
	return s
}

func (s *DescribeInspectionTaskReportResponse) SetStatusCode(v int32) *DescribeInspectionTaskReportResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInspectionTaskReportResponse) SetBody(v *DescribeInspectionTaskReportResponseBody) *DescribeInspectionTaskReportResponse {
	s.Body = v
	return s
}

func (s *DescribeInspectionTaskReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
