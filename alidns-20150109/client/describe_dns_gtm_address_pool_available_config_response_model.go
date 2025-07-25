// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsGtmAddressPoolAvailableConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDnsGtmAddressPoolAvailableConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDnsGtmAddressPoolAvailableConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDnsGtmAddressPoolAvailableConfigResponseBody) *DescribeDnsGtmAddressPoolAvailableConfigResponse
	GetBody() *DescribeDnsGtmAddressPoolAvailableConfigResponseBody
}

type DescribeDnsGtmAddressPoolAvailableConfigResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDnsGtmAddressPoolAvailableConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDnsGtmAddressPoolAvailableConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmAddressPoolAvailableConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAddressPoolAvailableConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDnsGtmAddressPoolAvailableConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDnsGtmAddressPoolAvailableConfigResponse) GetBody() *DescribeDnsGtmAddressPoolAvailableConfigResponseBody {
	return s.Body
}

func (s *DescribeDnsGtmAddressPoolAvailableConfigResponse) SetHeaders(v map[string]*string) *DescribeDnsGtmAddressPoolAvailableConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeDnsGtmAddressPoolAvailableConfigResponse) SetStatusCode(v int32) *DescribeDnsGtmAddressPoolAvailableConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDnsGtmAddressPoolAvailableConfigResponse) SetBody(v *DescribeDnsGtmAddressPoolAvailableConfigResponseBody) *DescribeDnsGtmAddressPoolAvailableConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeDnsGtmAddressPoolAvailableConfigResponse) Validate() error {
	return dara.Validate(s)
}
