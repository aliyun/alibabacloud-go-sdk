// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecondaryPublicIpAddressesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSecondaryPublicIpAddressesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSecondaryPublicIpAddressesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSecondaryPublicIpAddressesResponseBody) *DescribeSecondaryPublicIpAddressesResponse
	GetBody() *DescribeSecondaryPublicIpAddressesResponseBody
}

type DescribeSecondaryPublicIpAddressesResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSecondaryPublicIpAddressesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSecondaryPublicIpAddressesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecondaryPublicIpAddressesResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecondaryPublicIpAddressesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSecondaryPublicIpAddressesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSecondaryPublicIpAddressesResponse) GetBody() *DescribeSecondaryPublicIpAddressesResponseBody {
	return s.Body
}

func (s *DescribeSecondaryPublicIpAddressesResponse) SetHeaders(v map[string]*string) *DescribeSecondaryPublicIpAddressesResponse {
	s.Headers = v
	return s
}

func (s *DescribeSecondaryPublicIpAddressesResponse) SetStatusCode(v int32) *DescribeSecondaryPublicIpAddressesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSecondaryPublicIpAddressesResponse) SetBody(v *DescribeSecondaryPublicIpAddressesResponseBody) *DescribeSecondaryPublicIpAddressesResponse {
	s.Body = v
	return s
}

func (s *DescribeSecondaryPublicIpAddressesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
