// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePurchasedDeviceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePurchasedDeviceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePurchasedDeviceResponse
	GetStatusCode() *int32
	SetBody(v *DescribePurchasedDeviceResponseBody) *DescribePurchasedDeviceResponse
	GetBody() *DescribePurchasedDeviceResponseBody
}

type DescribePurchasedDeviceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePurchasedDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePurchasedDeviceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePurchasedDeviceResponse) GoString() string {
	return s.String()
}

func (s *DescribePurchasedDeviceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePurchasedDeviceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePurchasedDeviceResponse) GetBody() *DescribePurchasedDeviceResponseBody {
	return s.Body
}

func (s *DescribePurchasedDeviceResponse) SetHeaders(v map[string]*string) *DescribePurchasedDeviceResponse {
	s.Headers = v
	return s
}

func (s *DescribePurchasedDeviceResponse) SetStatusCode(v int32) *DescribePurchasedDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePurchasedDeviceResponse) SetBody(v *DescribePurchasedDeviceResponseBody) *DescribePurchasedDeviceResponse {
	s.Body = v
	return s
}

func (s *DescribePurchasedDeviceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
