// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProductsOfActiveMetricRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeProductsOfActiveMetricRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeProductsOfActiveMetricRuleResponse
	GetStatusCode() *int32
	SetBody(v *DescribeProductsOfActiveMetricRuleResponseBody) *DescribeProductsOfActiveMetricRuleResponse
	GetBody() *DescribeProductsOfActiveMetricRuleResponseBody
}

type DescribeProductsOfActiveMetricRuleResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeProductsOfActiveMetricRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeProductsOfActiveMetricRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductsOfActiveMetricRuleResponse) GoString() string {
	return s.String()
}

func (s *DescribeProductsOfActiveMetricRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeProductsOfActiveMetricRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeProductsOfActiveMetricRuleResponse) GetBody() *DescribeProductsOfActiveMetricRuleResponseBody {
	return s.Body
}

func (s *DescribeProductsOfActiveMetricRuleResponse) SetHeaders(v map[string]*string) *DescribeProductsOfActiveMetricRuleResponse {
	s.Headers = v
	return s
}

func (s *DescribeProductsOfActiveMetricRuleResponse) SetStatusCode(v int32) *DescribeProductsOfActiveMetricRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeProductsOfActiveMetricRuleResponse) SetBody(v *DescribeProductsOfActiveMetricRuleResponseBody) *DescribeProductsOfActiveMetricRuleResponse {
	s.Body = v
	return s
}

func (s *DescribeProductsOfActiveMetricRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
