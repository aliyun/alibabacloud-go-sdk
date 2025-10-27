// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomizeReportConfigDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCustomizeReportConfigDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCustomizeReportConfigDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCustomizeReportConfigDetailResponseBody) *DescribeCustomizeReportConfigDetailResponse
	GetBody() *DescribeCustomizeReportConfigDetailResponseBody
}

type DescribeCustomizeReportConfigDetailResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCustomizeReportConfigDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCustomizeReportConfigDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomizeReportConfigDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustomizeReportConfigDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCustomizeReportConfigDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCustomizeReportConfigDetailResponse) GetBody() *DescribeCustomizeReportConfigDetailResponseBody {
	return s.Body
}

func (s *DescribeCustomizeReportConfigDetailResponse) SetHeaders(v map[string]*string) *DescribeCustomizeReportConfigDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustomizeReportConfigDetailResponse) SetStatusCode(v int32) *DescribeCustomizeReportConfigDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCustomizeReportConfigDetailResponse) SetBody(v *DescribeCustomizeReportConfigDetailResponseBody) *DescribeCustomizeReportConfigDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeCustomizeReportConfigDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
