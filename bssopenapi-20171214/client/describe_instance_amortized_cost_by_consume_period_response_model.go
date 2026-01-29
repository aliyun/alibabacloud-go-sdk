// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceAmortizedCostByConsumePeriodResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceAmortizedCostByConsumePeriodResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceAmortizedCostByConsumePeriodResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceAmortizedCostByConsumePeriodResponseBody) *DescribeInstanceAmortizedCostByConsumePeriodResponse
	GetBody() *DescribeInstanceAmortizedCostByConsumePeriodResponseBody
}

type DescribeInstanceAmortizedCostByConsumePeriodResponse struct {
	Headers    map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceAmortizedCostByConsumePeriodResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceAmortizedCostByConsumePeriodResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAmortizedCostByConsumePeriodResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponse) GetBody() *DescribeInstanceAmortizedCostByConsumePeriodResponseBody {
	return s.Body
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponse) SetHeaders(v map[string]*string) *DescribeInstanceAmortizedCostByConsumePeriodResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponse) SetStatusCode(v int32) *DescribeInstanceAmortizedCostByConsumePeriodResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponse) SetBody(v *DescribeInstanceAmortizedCostByConsumePeriodResponseBody) *DescribeInstanceAmortizedCostByConsumePeriodResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
