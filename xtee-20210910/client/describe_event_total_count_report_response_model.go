// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventTotalCountReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEventTotalCountReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEventTotalCountReportResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEventTotalCountReportResponseBody) *DescribeEventTotalCountReportResponse
	GetBody() *DescribeEventTotalCountReportResponseBody
}

type DescribeEventTotalCountReportResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEventTotalCountReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEventTotalCountReportResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventTotalCountReportResponse) GoString() string {
	return s.String()
}

func (s *DescribeEventTotalCountReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEventTotalCountReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEventTotalCountReportResponse) GetBody() *DescribeEventTotalCountReportResponseBody {
	return s.Body
}

func (s *DescribeEventTotalCountReportResponse) SetHeaders(v map[string]*string) *DescribeEventTotalCountReportResponse {
	s.Headers = v
	return s
}

func (s *DescribeEventTotalCountReportResponse) SetStatusCode(v int32) *DescribeEventTotalCountReportResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEventTotalCountReportResponse) SetBody(v *DescribeEventTotalCountReportResponseBody) *DescribeEventTotalCountReportResponse {
	s.Body = v
	return s
}

func (s *DescribeEventTotalCountReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
