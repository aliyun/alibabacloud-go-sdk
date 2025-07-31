// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseNodePrivateNetworkAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReleaseNodePrivateNetworkAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReleaseNodePrivateNetworkAddressResponse
	GetStatusCode() *int32
	SetBody(v *ReleaseNodePrivateNetworkAddressResponseBody) *ReleaseNodePrivateNetworkAddressResponse
	GetBody() *ReleaseNodePrivateNetworkAddressResponseBody
}

type ReleaseNodePrivateNetworkAddressResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseNodePrivateNetworkAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseNodePrivateNetworkAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s ReleaseNodePrivateNetworkAddressResponse) GoString() string {
	return s.String()
}

func (s *ReleaseNodePrivateNetworkAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReleaseNodePrivateNetworkAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReleaseNodePrivateNetworkAddressResponse) GetBody() *ReleaseNodePrivateNetworkAddressResponseBody {
	return s.Body
}

func (s *ReleaseNodePrivateNetworkAddressResponse) SetHeaders(v map[string]*string) *ReleaseNodePrivateNetworkAddressResponse {
	s.Headers = v
	return s
}

func (s *ReleaseNodePrivateNetworkAddressResponse) SetStatusCode(v int32) *ReleaseNodePrivateNetworkAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseNodePrivateNetworkAddressResponse) SetBody(v *ReleaseNodePrivateNetworkAddressResponseBody) *ReleaseNodePrivateNetworkAddressResponse {
	s.Body = v
	return s
}

func (s *ReleaseNodePrivateNetworkAddressResponse) Validate() error {
	return dara.Validate(s)
}
