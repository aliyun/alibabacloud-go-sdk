// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseIpv6AddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReleaseIpv6AddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReleaseIpv6AddressResponse
	GetStatusCode() *int32
	SetBody(v *ReleaseIpv6AddressResponseBody) *ReleaseIpv6AddressResponse
	GetBody() *ReleaseIpv6AddressResponseBody
}

type ReleaseIpv6AddressResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseIpv6AddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseIpv6AddressResponse) String() string {
	return dara.Prettify(s)
}

func (s ReleaseIpv6AddressResponse) GoString() string {
	return s.String()
}

func (s *ReleaseIpv6AddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReleaseIpv6AddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReleaseIpv6AddressResponse) GetBody() *ReleaseIpv6AddressResponseBody {
	return s.Body
}

func (s *ReleaseIpv6AddressResponse) SetHeaders(v map[string]*string) *ReleaseIpv6AddressResponse {
	s.Headers = v
	return s
}

func (s *ReleaseIpv6AddressResponse) SetStatusCode(v int32) *ReleaseIpv6AddressResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseIpv6AddressResponse) SetBody(v *ReleaseIpv6AddressResponseBody) *ReleaseIpv6AddressResponse {
	s.Body = v
	return s
}

func (s *ReleaseIpv6AddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
