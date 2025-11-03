// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePublicIpAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePublicIpAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePublicIpAddressResponse
	GetStatusCode() *int32
	SetBody(v *DescribePublicIpAddressResponseBody) *DescribePublicIpAddressResponse
	GetBody() *DescribePublicIpAddressResponseBody
}

type DescribePublicIpAddressResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePublicIpAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePublicIpAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePublicIpAddressResponse) GoString() string {
	return s.String()
}

func (s *DescribePublicIpAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePublicIpAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePublicIpAddressResponse) GetBody() *DescribePublicIpAddressResponseBody {
	return s.Body
}

func (s *DescribePublicIpAddressResponse) SetHeaders(v map[string]*string) *DescribePublicIpAddressResponse {
	s.Headers = v
	return s
}

func (s *DescribePublicIpAddressResponse) SetStatusCode(v int32) *DescribePublicIpAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePublicIpAddressResponse) SetBody(v *DescribePublicIpAddressResponseBody) *DescribePublicIpAddressResponse {
	s.Body = v
	return s
}

func (s *DescribePublicIpAddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
