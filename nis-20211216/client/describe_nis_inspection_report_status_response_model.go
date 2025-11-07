// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNisInspectionReportStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNisInspectionReportStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNisInspectionReportStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNisInspectionReportStatusResponseBody) *DescribeNisInspectionReportStatusResponse
	GetBody() *DescribeNisInspectionReportStatusResponseBody
}

type DescribeNisInspectionReportStatusResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNisInspectionReportStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNisInspectionReportStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNisInspectionReportStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeNisInspectionReportStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNisInspectionReportStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNisInspectionReportStatusResponse) GetBody() *DescribeNisInspectionReportStatusResponseBody {
	return s.Body
}

func (s *DescribeNisInspectionReportStatusResponse) SetHeaders(v map[string]*string) *DescribeNisInspectionReportStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeNisInspectionReportStatusResponse) SetStatusCode(v int32) *DescribeNisInspectionReportStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNisInspectionReportStatusResponse) SetBody(v *DescribeNisInspectionReportStatusResponseBody) *DescribeNisInspectionReportStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeNisInspectionReportStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
