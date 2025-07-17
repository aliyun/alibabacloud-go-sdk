// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnvDropMetricsRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEnvDropMetricsRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEnvDropMetricsRuleResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEnvDropMetricsRuleResponseBody) *DescribeEnvDropMetricsRuleResponse
	GetBody() *DescribeEnvDropMetricsRuleResponseBody
}

type DescribeEnvDropMetricsRuleResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEnvDropMetricsRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEnvDropMetricsRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnvDropMetricsRuleResponse) GoString() string {
	return s.String()
}

func (s *DescribeEnvDropMetricsRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEnvDropMetricsRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEnvDropMetricsRuleResponse) GetBody() *DescribeEnvDropMetricsRuleResponseBody {
	return s.Body
}

func (s *DescribeEnvDropMetricsRuleResponse) SetHeaders(v map[string]*string) *DescribeEnvDropMetricsRuleResponse {
	s.Headers = v
	return s
}

func (s *DescribeEnvDropMetricsRuleResponse) SetStatusCode(v int32) *DescribeEnvDropMetricsRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEnvDropMetricsRuleResponse) SetBody(v *DescribeEnvDropMetricsRuleResponseBody) *DescribeEnvDropMetricsRuleResponse {
	s.Body = v
	return s
}

func (s *DescribeEnvDropMetricsRuleResponse) Validate() error {
	return dara.Validate(s)
}
