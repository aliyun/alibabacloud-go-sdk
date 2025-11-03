// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateIpv6AddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AllocateIpv6AddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AllocateIpv6AddressResponse
	GetStatusCode() *int32
	SetBody(v *AllocateIpv6AddressResponseBody) *AllocateIpv6AddressResponse
	GetBody() *AllocateIpv6AddressResponseBody
}

type AllocateIpv6AddressResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AllocateIpv6AddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AllocateIpv6AddressResponse) String() string {
	return dara.Prettify(s)
}

func (s AllocateIpv6AddressResponse) GoString() string {
	return s.String()
}

func (s *AllocateIpv6AddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AllocateIpv6AddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AllocateIpv6AddressResponse) GetBody() *AllocateIpv6AddressResponseBody {
	return s.Body
}

func (s *AllocateIpv6AddressResponse) SetHeaders(v map[string]*string) *AllocateIpv6AddressResponse {
	s.Headers = v
	return s
}

func (s *AllocateIpv6AddressResponse) SetStatusCode(v int32) *AllocateIpv6AddressResponse {
	s.StatusCode = &v
	return s
}

func (s *AllocateIpv6AddressResponse) SetBody(v *AllocateIpv6AddressResponseBody) *AllocateIpv6AddressResponse {
	s.Body = v
	return s
}

func (s *AllocateIpv6AddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
