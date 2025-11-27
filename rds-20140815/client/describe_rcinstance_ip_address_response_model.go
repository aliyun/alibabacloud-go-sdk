// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCInstanceIpAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRCInstanceIpAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRCInstanceIpAddressResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRCInstanceIpAddressResponseBody) *DescribeRCInstanceIpAddressResponse
	GetBody() *DescribeRCInstanceIpAddressResponseBody
}

type DescribeRCInstanceIpAddressResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRCInstanceIpAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRCInstanceIpAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceIpAddressResponse) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceIpAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRCInstanceIpAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRCInstanceIpAddressResponse) GetBody() *DescribeRCInstanceIpAddressResponseBody {
	return s.Body
}

func (s *DescribeRCInstanceIpAddressResponse) SetHeaders(v map[string]*string) *DescribeRCInstanceIpAddressResponse {
	s.Headers = v
	return s
}

func (s *DescribeRCInstanceIpAddressResponse) SetStatusCode(v int32) *DescribeRCInstanceIpAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRCInstanceIpAddressResponse) SetBody(v *DescribeRCInstanceIpAddressResponseBody) *DescribeRCInstanceIpAddressResponse {
	s.Body = v
	return s
}

func (s *DescribeRCInstanceIpAddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
