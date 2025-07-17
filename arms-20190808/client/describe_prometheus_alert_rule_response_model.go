// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePrometheusAlertRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePrometheusAlertRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePrometheusAlertRuleResponse
	GetStatusCode() *int32
	SetBody(v *DescribePrometheusAlertRuleResponseBody) *DescribePrometheusAlertRuleResponse
	GetBody() *DescribePrometheusAlertRuleResponseBody
}

type DescribePrometheusAlertRuleResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePrometheusAlertRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePrometheusAlertRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePrometheusAlertRuleResponse) GoString() string {
	return s.String()
}

func (s *DescribePrometheusAlertRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePrometheusAlertRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePrometheusAlertRuleResponse) GetBody() *DescribePrometheusAlertRuleResponseBody {
	return s.Body
}

func (s *DescribePrometheusAlertRuleResponse) SetHeaders(v map[string]*string) *DescribePrometheusAlertRuleResponse {
	s.Headers = v
	return s
}

func (s *DescribePrometheusAlertRuleResponse) SetStatusCode(v int32) *DescribePrometheusAlertRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponse) SetBody(v *DescribePrometheusAlertRuleResponseBody) *DescribePrometheusAlertRuleResponse {
	s.Body = v
	return s
}

func (s *DescribePrometheusAlertRuleResponse) Validate() error {
	return dara.Validate(s)
}
