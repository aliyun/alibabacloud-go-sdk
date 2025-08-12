// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetricRuleBlackListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMetricRuleBlackListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMetricRuleBlackListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMetricRuleBlackListResponseBody) *DescribeMetricRuleBlackListResponse
	GetBody() *DescribeMetricRuleBlackListResponseBody
}

type DescribeMetricRuleBlackListResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMetricRuleBlackListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMetricRuleBlackListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleBlackListResponse) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleBlackListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMetricRuleBlackListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMetricRuleBlackListResponse) GetBody() *DescribeMetricRuleBlackListResponseBody {
	return s.Body
}

func (s *DescribeMetricRuleBlackListResponse) SetHeaders(v map[string]*string) *DescribeMetricRuleBlackListResponse {
	s.Headers = v
	return s
}

func (s *DescribeMetricRuleBlackListResponse) SetStatusCode(v int32) *DescribeMetricRuleBlackListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMetricRuleBlackListResponse) SetBody(v *DescribeMetricRuleBlackListResponseBody) *DescribeMetricRuleBlackListResponse {
	s.Body = v
	return s
}

func (s *DescribeMetricRuleBlackListResponse) Validate() error {
	return dara.Validate(s)
}
