// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomizeRuleTestHistogramResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCustomizeRuleTestHistogramResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCustomizeRuleTestHistogramResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCustomizeRuleTestHistogramResponseBody) *DescribeCustomizeRuleTestHistogramResponse
	GetBody() *DescribeCustomizeRuleTestHistogramResponseBody
}

type DescribeCustomizeRuleTestHistogramResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCustomizeRuleTestHistogramResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCustomizeRuleTestHistogramResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomizeRuleTestHistogramResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustomizeRuleTestHistogramResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCustomizeRuleTestHistogramResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCustomizeRuleTestHistogramResponse) GetBody() *DescribeCustomizeRuleTestHistogramResponseBody {
	return s.Body
}

func (s *DescribeCustomizeRuleTestHistogramResponse) SetHeaders(v map[string]*string) *DescribeCustomizeRuleTestHistogramResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustomizeRuleTestHistogramResponse) SetStatusCode(v int32) *DescribeCustomizeRuleTestHistogramResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCustomizeRuleTestHistogramResponse) SetBody(v *DescribeCustomizeRuleTestHistogramResponseBody) *DescribeCustomizeRuleTestHistogramResponse {
	s.Body = v
	return s
}

func (s *DescribeCustomizeRuleTestHistogramResponse) Validate() error {
	return dara.Validate(s)
}
