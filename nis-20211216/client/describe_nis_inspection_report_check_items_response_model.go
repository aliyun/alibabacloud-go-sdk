// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNisInspectionReportCheckItemsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNisInspectionReportCheckItemsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNisInspectionReportCheckItemsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNisInspectionReportCheckItemsResponseBody) *DescribeNisInspectionReportCheckItemsResponse
	GetBody() *DescribeNisInspectionReportCheckItemsResponseBody
}

type DescribeNisInspectionReportCheckItemsResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNisInspectionReportCheckItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNisInspectionReportCheckItemsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNisInspectionReportCheckItemsResponse) GoString() string {
	return s.String()
}

func (s *DescribeNisInspectionReportCheckItemsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNisInspectionReportCheckItemsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNisInspectionReportCheckItemsResponse) GetBody() *DescribeNisInspectionReportCheckItemsResponseBody {
	return s.Body
}

func (s *DescribeNisInspectionReportCheckItemsResponse) SetHeaders(v map[string]*string) *DescribeNisInspectionReportCheckItemsResponse {
	s.Headers = v
	return s
}

func (s *DescribeNisInspectionReportCheckItemsResponse) SetStatusCode(v int32) *DescribeNisInspectionReportCheckItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNisInspectionReportCheckItemsResponse) SetBody(v *DescribeNisInspectionReportCheckItemsResponseBody) *DescribeNisInspectionReportCheckItemsResponse {
	s.Body = v
	return s
}

func (s *DescribeNisInspectionReportCheckItemsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
