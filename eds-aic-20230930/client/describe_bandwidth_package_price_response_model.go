// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBandwidthPackagePriceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBandwidthPackagePriceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBandwidthPackagePriceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBandwidthPackagePriceResponseBody) *DescribeBandwidthPackagePriceResponse
	GetBody() *DescribeBandwidthPackagePriceResponseBody
}

type DescribeBandwidthPackagePriceResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBandwidthPackagePriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBandwidthPackagePriceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBandwidthPackagePriceResponse) GoString() string {
	return s.String()
}

func (s *DescribeBandwidthPackagePriceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBandwidthPackagePriceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBandwidthPackagePriceResponse) GetBody() *DescribeBandwidthPackagePriceResponseBody {
	return s.Body
}

func (s *DescribeBandwidthPackagePriceResponse) SetHeaders(v map[string]*string) *DescribeBandwidthPackagePriceResponse {
	s.Headers = v
	return s
}

func (s *DescribeBandwidthPackagePriceResponse) SetStatusCode(v int32) *DescribeBandwidthPackagePriceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBandwidthPackagePriceResponse) SetBody(v *DescribeBandwidthPackagePriceResponseBody) *DescribeBandwidthPackagePriceResponse {
	s.Body = v
	return s
}

func (s *DescribeBandwidthPackagePriceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
