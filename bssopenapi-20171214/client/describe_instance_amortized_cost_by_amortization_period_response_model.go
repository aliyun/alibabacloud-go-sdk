// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceAmortizedCostByAmortizationPeriodResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceAmortizedCostByAmortizationPeriodResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceAmortizedCostByAmortizationPeriodResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBody) *DescribeInstanceAmortizedCostByAmortizationPeriodResponse
	GetBody() *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBody
}

type DescribeInstanceAmortizedCostByAmortizationPeriodResponse struct {
	Headers    map[string]*string                                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceAmortizedCostByAmortizationPeriodResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAmortizedCostByAmortizationPeriodResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponse) GetBody() *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBody {
	return s.Body
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponse) SetHeaders(v map[string]*string) *DescribeInstanceAmortizedCostByAmortizationPeriodResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponse) SetStatusCode(v int32) *DescribeInstanceAmortizedCostByAmortizationPeriodResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponse) SetBody(v *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBody) *DescribeInstanceAmortizedCostByAmortizationPeriodResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponse) Validate() error {
	return dara.Validate(s)
}
