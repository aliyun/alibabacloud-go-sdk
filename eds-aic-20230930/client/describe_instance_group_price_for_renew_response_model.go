// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceGroupPriceForRenewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceGroupPriceForRenewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceGroupPriceForRenewResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceGroupPriceForRenewResponseBody) *DescribeInstanceGroupPriceForRenewResponse
	GetBody() *DescribeInstanceGroupPriceForRenewResponseBody
}

type DescribeInstanceGroupPriceForRenewResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceGroupPriceForRenewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceGroupPriceForRenewResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceGroupPriceForRenewResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceGroupPriceForRenewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceGroupPriceForRenewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceGroupPriceForRenewResponse) GetBody() *DescribeInstanceGroupPriceForRenewResponseBody {
	return s.Body
}

func (s *DescribeInstanceGroupPriceForRenewResponse) SetHeaders(v map[string]*string) *DescribeInstanceGroupPriceForRenewResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceGroupPriceForRenewResponse) SetStatusCode(v int32) *DescribeInstanceGroupPriceForRenewResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceGroupPriceForRenewResponse) SetBody(v *DescribeInstanceGroupPriceForRenewResponseBody) *DescribeInstanceGroupPriceForRenewResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceGroupPriceForRenewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
