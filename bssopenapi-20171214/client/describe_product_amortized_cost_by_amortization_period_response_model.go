// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProductAmortizedCostByAmortizationPeriodResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeProductAmortizedCostByAmortizationPeriodResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeProductAmortizedCostByAmortizationPeriodResponse
	GetStatusCode() *int32
	SetBody(v *DescribeProductAmortizedCostByAmortizationPeriodResponseBody) *DescribeProductAmortizedCostByAmortizationPeriodResponse
	GetBody() *DescribeProductAmortizedCostByAmortizationPeriodResponseBody
}

type DescribeProductAmortizedCostByAmortizationPeriodResponse struct {
	Headers    map[string]*string                                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeProductAmortizedCostByAmortizationPeriodResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeProductAmortizedCostByAmortizationPeriodResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductAmortizedCostByAmortizationPeriodResponse) GoString() string {
	return s.String()
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponse) GetBody() *DescribeProductAmortizedCostByAmortizationPeriodResponseBody {
	return s.Body
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponse) SetHeaders(v map[string]*string) *DescribeProductAmortizedCostByAmortizationPeriodResponse {
	s.Headers = v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponse) SetStatusCode(v int32) *DescribeProductAmortizedCostByAmortizationPeriodResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponse) SetBody(v *DescribeProductAmortizedCostByAmortizationPeriodResponseBody) *DescribeProductAmortizedCostByAmortizationPeriodResponse {
	s.Body = v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponse) Validate() error {
	return dara.Validate(s)
}
