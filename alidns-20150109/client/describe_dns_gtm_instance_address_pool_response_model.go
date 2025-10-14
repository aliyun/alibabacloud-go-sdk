// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsGtmInstanceAddressPoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDnsGtmInstanceAddressPoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDnsGtmInstanceAddressPoolResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDnsGtmInstanceAddressPoolResponseBody) *DescribeDnsGtmInstanceAddressPoolResponse
	GetBody() *DescribeDnsGtmInstanceAddressPoolResponseBody
}

type DescribeDnsGtmInstanceAddressPoolResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDnsGtmInstanceAddressPoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDnsGtmInstanceAddressPoolResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmInstanceAddressPoolResponse) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstanceAddressPoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDnsGtmInstanceAddressPoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDnsGtmInstanceAddressPoolResponse) GetBody() *DescribeDnsGtmInstanceAddressPoolResponseBody {
	return s.Body
}

func (s *DescribeDnsGtmInstanceAddressPoolResponse) SetHeaders(v map[string]*string) *DescribeDnsGtmInstanceAddressPoolResponse {
	s.Headers = v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolResponse) SetStatusCode(v int32) *DescribeDnsGtmInstanceAddressPoolResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolResponse) SetBody(v *DescribeDnsGtmInstanceAddressPoolResponseBody) *DescribeDnsGtmInstanceAddressPoolResponse {
	s.Body = v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
