// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebReportTopIpResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWebReportTopIpResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWebReportTopIpResponse
	GetStatusCode() *int32
	SetBody(v *DescribeWebReportTopIpResponseBody) *DescribeWebReportTopIpResponse
	GetBody() *DescribeWebReportTopIpResponseBody
}

type DescribeWebReportTopIpResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWebReportTopIpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWebReportTopIpResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebReportTopIpResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebReportTopIpResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWebReportTopIpResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWebReportTopIpResponse) GetBody() *DescribeWebReportTopIpResponseBody {
	return s.Body
}

func (s *DescribeWebReportTopIpResponse) SetHeaders(v map[string]*string) *DescribeWebReportTopIpResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebReportTopIpResponse) SetStatusCode(v int32) *DescribeWebReportTopIpResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWebReportTopIpResponse) SetBody(v *DescribeWebReportTopIpResponseBody) *DescribeWebReportTopIpResponse {
	s.Body = v
	return s
}

func (s *DescribeWebReportTopIpResponse) Validate() error {
	return dara.Validate(s)
}
