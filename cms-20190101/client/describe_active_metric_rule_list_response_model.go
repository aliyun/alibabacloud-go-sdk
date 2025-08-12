// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeActiveMetricRuleListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeActiveMetricRuleListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeActiveMetricRuleListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeActiveMetricRuleListResponseBody) *DescribeActiveMetricRuleListResponse
	GetBody() *DescribeActiveMetricRuleListResponseBody
}

type DescribeActiveMetricRuleListResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeActiveMetricRuleListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeActiveMetricRuleListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeActiveMetricRuleListResponse) GoString() string {
	return s.String()
}

func (s *DescribeActiveMetricRuleListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeActiveMetricRuleListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeActiveMetricRuleListResponse) GetBody() *DescribeActiveMetricRuleListResponseBody {
	return s.Body
}

func (s *DescribeActiveMetricRuleListResponse) SetHeaders(v map[string]*string) *DescribeActiveMetricRuleListResponse {
	s.Headers = v
	return s
}

func (s *DescribeActiveMetricRuleListResponse) SetStatusCode(v int32) *DescribeActiveMetricRuleListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponse) SetBody(v *DescribeActiveMetricRuleListResponseBody) *DescribeActiveMetricRuleListResponse {
	s.Body = v
	return s
}

func (s *DescribeActiveMetricRuleListResponse) Validate() error {
	return dara.Validate(s)
}
