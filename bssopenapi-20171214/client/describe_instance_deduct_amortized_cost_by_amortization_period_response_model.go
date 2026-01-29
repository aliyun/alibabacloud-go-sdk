// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceDeductAmortizedCostByAmortizationPeriodResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBody) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponse
	GetBody() *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBody
}

type DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponse struct {
	Headers    map[string]*string                                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponse) GetBody() *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBody {
	return s.Body
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponse) SetHeaders(v map[string]*string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponse) SetStatusCode(v int32) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponse) SetBody(v *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBody) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
