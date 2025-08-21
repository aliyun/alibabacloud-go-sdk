// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePurchasedDevicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePurchasedDevicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePurchasedDevicesResponse
	GetStatusCode() *int32
	SetBody(v *DescribePurchasedDevicesResponseBody) *DescribePurchasedDevicesResponse
	GetBody() *DescribePurchasedDevicesResponseBody
}

type DescribePurchasedDevicesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePurchasedDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePurchasedDevicesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePurchasedDevicesResponse) GoString() string {
	return s.String()
}

func (s *DescribePurchasedDevicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePurchasedDevicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePurchasedDevicesResponse) GetBody() *DescribePurchasedDevicesResponseBody {
	return s.Body
}

func (s *DescribePurchasedDevicesResponse) SetHeaders(v map[string]*string) *DescribePurchasedDevicesResponse {
	s.Headers = v
	return s
}

func (s *DescribePurchasedDevicesResponse) SetStatusCode(v int32) *DescribePurchasedDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePurchasedDevicesResponse) SetBody(v *DescribePurchasedDevicesResponseBody) *DescribePurchasedDevicesResponse {
	s.Body = v
	return s
}

func (s *DescribePurchasedDevicesResponse) Validate() error {
	return dara.Validate(s)
}
