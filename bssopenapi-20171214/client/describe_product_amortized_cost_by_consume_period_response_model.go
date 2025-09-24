// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProductAmortizedCostByConsumePeriodResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeProductAmortizedCostByConsumePeriodResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeProductAmortizedCostByConsumePeriodResponse
	GetStatusCode() *int32
	SetBody(v *DescribeProductAmortizedCostByConsumePeriodResponseBody) *DescribeProductAmortizedCostByConsumePeriodResponse
	GetBody() *DescribeProductAmortizedCostByConsumePeriodResponseBody
}

type DescribeProductAmortizedCostByConsumePeriodResponse struct {
	Headers    map[string]*string                                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeProductAmortizedCostByConsumePeriodResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeProductAmortizedCostByConsumePeriodResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductAmortizedCostByConsumePeriodResponse) GoString() string {
	return s.String()
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponse) GetBody() *DescribeProductAmortizedCostByConsumePeriodResponseBody {
	return s.Body
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponse) SetHeaders(v map[string]*string) *DescribeProductAmortizedCostByConsumePeriodResponse {
	s.Headers = v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponse) SetStatusCode(v int32) *DescribeProductAmortizedCostByConsumePeriodResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponse) SetBody(v *DescribeProductAmortizedCostByConsumePeriodResponseBody) *DescribeProductAmortizedCostByConsumePeriodResponse {
	s.Body = v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponse) Validate() error {
	return dara.Validate(s)
}
