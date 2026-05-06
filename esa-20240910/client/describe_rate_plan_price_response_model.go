// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRatePlanPriceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRatePlanPriceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRatePlanPriceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRatePlanPriceResponseBody) *DescribeRatePlanPriceResponse
	GetBody() *DescribeRatePlanPriceResponseBody
}

type DescribeRatePlanPriceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRatePlanPriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRatePlanPriceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRatePlanPriceResponse) GoString() string {
	return s.String()
}

func (s *DescribeRatePlanPriceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRatePlanPriceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRatePlanPriceResponse) GetBody() *DescribeRatePlanPriceResponseBody {
	return s.Body
}

func (s *DescribeRatePlanPriceResponse) SetHeaders(v map[string]*string) *DescribeRatePlanPriceResponse {
	s.Headers = v
	return s
}

func (s *DescribeRatePlanPriceResponse) SetStatusCode(v int32) *DescribeRatePlanPriceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRatePlanPriceResponse) SetBody(v *DescribeRatePlanPriceResponseBody) *DescribeRatePlanPriceResponse {
	s.Body = v
	return s
}

func (s *DescribeRatePlanPriceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
