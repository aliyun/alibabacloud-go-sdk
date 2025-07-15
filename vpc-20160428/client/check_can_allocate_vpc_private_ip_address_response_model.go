// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckCanAllocateVpcPrivateIpAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckCanAllocateVpcPrivateIpAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckCanAllocateVpcPrivateIpAddressResponse
	GetStatusCode() *int32
	SetBody(v *CheckCanAllocateVpcPrivateIpAddressResponseBody) *CheckCanAllocateVpcPrivateIpAddressResponse
	GetBody() *CheckCanAllocateVpcPrivateIpAddressResponseBody
}

type CheckCanAllocateVpcPrivateIpAddressResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckCanAllocateVpcPrivateIpAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckCanAllocateVpcPrivateIpAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckCanAllocateVpcPrivateIpAddressResponse) GoString() string {
	return s.String()
}

func (s *CheckCanAllocateVpcPrivateIpAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckCanAllocateVpcPrivateIpAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckCanAllocateVpcPrivateIpAddressResponse) GetBody() *CheckCanAllocateVpcPrivateIpAddressResponseBody {
	return s.Body
}

func (s *CheckCanAllocateVpcPrivateIpAddressResponse) SetHeaders(v map[string]*string) *CheckCanAllocateVpcPrivateIpAddressResponse {
	s.Headers = v
	return s
}

func (s *CheckCanAllocateVpcPrivateIpAddressResponse) SetStatusCode(v int32) *CheckCanAllocateVpcPrivateIpAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckCanAllocateVpcPrivateIpAddressResponse) SetBody(v *CheckCanAllocateVpcPrivateIpAddressResponseBody) *CheckCanAllocateVpcPrivateIpAddressResponse {
	s.Body = v
	return s
}

func (s *CheckCanAllocateVpcPrivateIpAddressResponse) Validate() error {
	return dara.Validate(s)
}
