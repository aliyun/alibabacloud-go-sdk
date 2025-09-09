// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRdsPerformanceSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRdsPerformanceSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRdsPerformanceSummaryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRdsPerformanceSummaryResponseBody) *DescribeRdsPerformanceSummaryResponse
	GetBody() *DescribeRdsPerformanceSummaryResponseBody
}

type DescribeRdsPerformanceSummaryResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRdsPerformanceSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRdsPerformanceSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRdsPerformanceSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeRdsPerformanceSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRdsPerformanceSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRdsPerformanceSummaryResponse) GetBody() *DescribeRdsPerformanceSummaryResponseBody {
	return s.Body
}

func (s *DescribeRdsPerformanceSummaryResponse) SetHeaders(v map[string]*string) *DescribeRdsPerformanceSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeRdsPerformanceSummaryResponse) SetStatusCode(v int32) *DescribeRdsPerformanceSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRdsPerformanceSummaryResponse) SetBody(v *DescribeRdsPerformanceSummaryResponseBody) *DescribeRdsPerformanceSummaryResponse {
	s.Body = v
	return s
}

func (s *DescribeRdsPerformanceSummaryResponse) Validate() error {
	return dara.Validate(s)
}
