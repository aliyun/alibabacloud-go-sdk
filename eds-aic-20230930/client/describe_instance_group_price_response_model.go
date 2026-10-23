// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceGroupPriceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceGroupPriceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceGroupPriceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceGroupPriceResponseBody) *DescribeInstanceGroupPriceResponse
	GetBody() *DescribeInstanceGroupPriceResponseBody
}

type DescribeInstanceGroupPriceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceGroupPriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceGroupPriceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceGroupPriceResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceGroupPriceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceGroupPriceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceGroupPriceResponse) GetBody() *DescribeInstanceGroupPriceResponseBody {
	return s.Body
}

func (s *DescribeInstanceGroupPriceResponse) SetHeaders(v map[string]*string) *DescribeInstanceGroupPriceResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceGroupPriceResponse) SetStatusCode(v int32) *DescribeInstanceGroupPriceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceGroupPriceResponse) SetBody(v *DescribeInstanceGroupPriceResponseBody) *DescribeInstanceGroupPriceResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceGroupPriceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
