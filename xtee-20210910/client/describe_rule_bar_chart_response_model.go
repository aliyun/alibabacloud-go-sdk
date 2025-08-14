// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRuleBarChartResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRuleBarChartResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRuleBarChartResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRuleBarChartResponseBody) *DescribeRuleBarChartResponse
	GetBody() *DescribeRuleBarChartResponseBody
}

type DescribeRuleBarChartResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRuleBarChartResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRuleBarChartResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleBarChartResponse) GoString() string {
	return s.String()
}

func (s *DescribeRuleBarChartResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRuleBarChartResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRuleBarChartResponse) GetBody() *DescribeRuleBarChartResponseBody {
	return s.Body
}

func (s *DescribeRuleBarChartResponse) SetHeaders(v map[string]*string) *DescribeRuleBarChartResponse {
	s.Headers = v
	return s
}

func (s *DescribeRuleBarChartResponse) SetStatusCode(v int32) *DescribeRuleBarChartResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRuleBarChartResponse) SetBody(v *DescribeRuleBarChartResponseBody) *DescribeRuleBarChartResponse {
	s.Body = v
	return s
}

func (s *DescribeRuleBarChartResponse) Validate() error {
	return dara.Validate(s)
}
