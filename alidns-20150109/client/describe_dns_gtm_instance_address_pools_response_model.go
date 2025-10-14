// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsGtmInstanceAddressPoolsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDnsGtmInstanceAddressPoolsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDnsGtmInstanceAddressPoolsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDnsGtmInstanceAddressPoolsResponseBody) *DescribeDnsGtmInstanceAddressPoolsResponse
	GetBody() *DescribeDnsGtmInstanceAddressPoolsResponseBody
}

type DescribeDnsGtmInstanceAddressPoolsResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDnsGtmInstanceAddressPoolsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDnsGtmInstanceAddressPoolsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmInstanceAddressPoolsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponse) GetBody() *DescribeDnsGtmInstanceAddressPoolsResponseBody {
	return s.Body
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponse) SetHeaders(v map[string]*string) *DescribeDnsGtmInstanceAddressPoolsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponse) SetStatusCode(v int32) *DescribeDnsGtmInstanceAddressPoolsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponse) SetBody(v *DescribeDnsGtmInstanceAddressPoolsResponseBody) *DescribeDnsGtmInstanceAddressPoolsResponse {
	s.Body = v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
