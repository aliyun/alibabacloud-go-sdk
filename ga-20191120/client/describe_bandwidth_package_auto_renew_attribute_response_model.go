// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBandwidthPackageAutoRenewAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBandwidthPackageAutoRenewAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBandwidthPackageAutoRenewAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBandwidthPackageAutoRenewAttributeResponseBody) *DescribeBandwidthPackageAutoRenewAttributeResponse
	GetBody() *DescribeBandwidthPackageAutoRenewAttributeResponseBody
}

type DescribeBandwidthPackageAutoRenewAttributeResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBandwidthPackageAutoRenewAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBandwidthPackageAutoRenewAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBandwidthPackageAutoRenewAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeBandwidthPackageAutoRenewAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBandwidthPackageAutoRenewAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBandwidthPackageAutoRenewAttributeResponse) GetBody() *DescribeBandwidthPackageAutoRenewAttributeResponseBody {
	return s.Body
}

func (s *DescribeBandwidthPackageAutoRenewAttributeResponse) SetHeaders(v map[string]*string) *DescribeBandwidthPackageAutoRenewAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeBandwidthPackageAutoRenewAttributeResponse) SetStatusCode(v int32) *DescribeBandwidthPackageAutoRenewAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBandwidthPackageAutoRenewAttributeResponse) SetBody(v *DescribeBandwidthPackageAutoRenewAttributeResponseBody) *DescribeBandwidthPackageAutoRenewAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeBandwidthPackageAutoRenewAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
