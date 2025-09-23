// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUpgradeReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUpgradeReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUpgradeReportResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUpgradeReportResponseBody) *DescribeUpgradeReportResponse
	GetBody() *DescribeUpgradeReportResponseBody
}

type DescribeUpgradeReportResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUpgradeReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUpgradeReportResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUpgradeReportResponse) GoString() string {
	return s.String()
}

func (s *DescribeUpgradeReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUpgradeReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUpgradeReportResponse) GetBody() *DescribeUpgradeReportResponseBody {
	return s.Body
}

func (s *DescribeUpgradeReportResponse) SetHeaders(v map[string]*string) *DescribeUpgradeReportResponse {
	s.Headers = v
	return s
}

func (s *DescribeUpgradeReportResponse) SetStatusCode(v int32) *DescribeUpgradeReportResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUpgradeReportResponse) SetBody(v *DescribeUpgradeReportResponseBody) *DescribeUpgradeReportResponse {
	s.Body = v
	return s
}

func (s *DescribeUpgradeReportResponse) Validate() error {
	return dara.Validate(s)
}
