// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateIpAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AllocateIpAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AllocateIpAddressResponse
	GetStatusCode() *int32
	SetBody(v *AllocateIpAddressResponseBody) *AllocateIpAddressResponse
	GetBody() *AllocateIpAddressResponseBody
}

type AllocateIpAddressResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AllocateIpAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AllocateIpAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s AllocateIpAddressResponse) GoString() string {
	return s.String()
}

func (s *AllocateIpAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AllocateIpAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AllocateIpAddressResponse) GetBody() *AllocateIpAddressResponseBody {
	return s.Body
}

func (s *AllocateIpAddressResponse) SetHeaders(v map[string]*string) *AllocateIpAddressResponse {
	s.Headers = v
	return s
}

func (s *AllocateIpAddressResponse) SetStatusCode(v int32) *AllocateIpAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *AllocateIpAddressResponse) SetBody(v *AllocateIpAddressResponseBody) *AllocateIpAddressResponse {
	s.Body = v
	return s
}

func (s *AllocateIpAddressResponse) Validate() error {
	return dara.Validate(s)
}
