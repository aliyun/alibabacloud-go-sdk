// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateNodePrivateNetworkAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AllocateNodePrivateNetworkAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AllocateNodePrivateNetworkAddressResponse
	GetStatusCode() *int32
	SetBody(v *AllocateNodePrivateNetworkAddressResponseBody) *AllocateNodePrivateNetworkAddressResponse
	GetBody() *AllocateNodePrivateNetworkAddressResponseBody
}

type AllocateNodePrivateNetworkAddressResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AllocateNodePrivateNetworkAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AllocateNodePrivateNetworkAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s AllocateNodePrivateNetworkAddressResponse) GoString() string {
	return s.String()
}

func (s *AllocateNodePrivateNetworkAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AllocateNodePrivateNetworkAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AllocateNodePrivateNetworkAddressResponse) GetBody() *AllocateNodePrivateNetworkAddressResponseBody {
	return s.Body
}

func (s *AllocateNodePrivateNetworkAddressResponse) SetHeaders(v map[string]*string) *AllocateNodePrivateNetworkAddressResponse {
	s.Headers = v
	return s
}

func (s *AllocateNodePrivateNetworkAddressResponse) SetStatusCode(v int32) *AllocateNodePrivateNetworkAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *AllocateNodePrivateNetworkAddressResponse) SetBody(v *AllocateNodePrivateNetworkAddressResponseBody) *AllocateNodePrivateNetworkAddressResponse {
	s.Body = v
	return s
}

func (s *AllocateNodePrivateNetworkAddressResponse) Validate() error {
	return dara.Validate(s)
}
