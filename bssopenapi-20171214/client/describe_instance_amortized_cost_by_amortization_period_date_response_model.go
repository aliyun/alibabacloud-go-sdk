// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceAmortizedCostByAmortizationPeriodDateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBody) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponse
	GetBody() *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBody
}

type DescribeInstanceAmortizedCostByAmortizationPeriodDateResponse struct {
	Headers    map[string]*string                                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceAmortizedCostByAmortizationPeriodDateResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAmortizedCostByAmortizationPeriodDateResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponse) GetBody() *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBody {
	return s.Body
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponse) SetHeaders(v map[string]*string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponse) SetStatusCode(v int32) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponse) SetBody(v *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBody) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
