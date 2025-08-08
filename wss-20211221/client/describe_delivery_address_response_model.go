// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeliveryAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDeliveryAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDeliveryAddressResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDeliveryAddressResponseBody) *DescribeDeliveryAddressResponse
	GetBody() *DescribeDeliveryAddressResponseBody
}

type DescribeDeliveryAddressResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDeliveryAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDeliveryAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeliveryAddressResponse) GoString() string {
	return s.String()
}

func (s *DescribeDeliveryAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDeliveryAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDeliveryAddressResponse) GetBody() *DescribeDeliveryAddressResponseBody {
	return s.Body
}

func (s *DescribeDeliveryAddressResponse) SetHeaders(v map[string]*string) *DescribeDeliveryAddressResponse {
	s.Headers = v
	return s
}

func (s *DescribeDeliveryAddressResponse) SetStatusCode(v int32) *DescribeDeliveryAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDeliveryAddressResponse) SetBody(v *DescribeDeliveryAddressResponseBody) *DescribeDeliveryAddressResponse {
	s.Body = v
	return s
}

func (s *DescribeDeliveryAddressResponse) Validate() error {
	return dara.Validate(s)
}
