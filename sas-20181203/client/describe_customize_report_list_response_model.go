// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomizeReportListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCustomizeReportListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCustomizeReportListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCustomizeReportListResponseBody) *DescribeCustomizeReportListResponse
	GetBody() *DescribeCustomizeReportListResponseBody
}

type DescribeCustomizeReportListResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCustomizeReportListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCustomizeReportListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomizeReportListResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustomizeReportListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCustomizeReportListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCustomizeReportListResponse) GetBody() *DescribeCustomizeReportListResponseBody {
	return s.Body
}

func (s *DescribeCustomizeReportListResponse) SetHeaders(v map[string]*string) *DescribeCustomizeReportListResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustomizeReportListResponse) SetStatusCode(v int32) *DescribeCustomizeReportListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCustomizeReportListResponse) SetBody(v *DescribeCustomizeReportListResponseBody) *DescribeCustomizeReportListResponse {
	s.Body = v
	return s
}

func (s *DescribeCustomizeReportListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
