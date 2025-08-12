// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetricRuleListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMetricRuleListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMetricRuleListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMetricRuleListResponseBody) *DescribeMetricRuleListResponse
	GetBody() *DescribeMetricRuleListResponseBody
}

type DescribeMetricRuleListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMetricRuleListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMetricRuleListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleListResponse) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMetricRuleListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMetricRuleListResponse) GetBody() *DescribeMetricRuleListResponseBody {
	return s.Body
}

func (s *DescribeMetricRuleListResponse) SetHeaders(v map[string]*string) *DescribeMetricRuleListResponse {
	s.Headers = v
	return s
}

func (s *DescribeMetricRuleListResponse) SetStatusCode(v int32) *DescribeMetricRuleListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMetricRuleListResponse) SetBody(v *DescribeMetricRuleListResponseBody) *DescribeMetricRuleListResponse {
	s.Body = v
	return s
}

func (s *DescribeMetricRuleListResponse) Validate() error {
	return dara.Validate(s)
}
