// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleasePublicNetworkAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReleasePublicNetworkAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReleasePublicNetworkAddressResponse
	GetStatusCode() *int32
	SetBody(v *ReleasePublicNetworkAddressResponseBody) *ReleasePublicNetworkAddressResponse
	GetBody() *ReleasePublicNetworkAddressResponseBody
}

type ReleasePublicNetworkAddressResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleasePublicNetworkAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleasePublicNetworkAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s ReleasePublicNetworkAddressResponse) GoString() string {
	return s.String()
}

func (s *ReleasePublicNetworkAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReleasePublicNetworkAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReleasePublicNetworkAddressResponse) GetBody() *ReleasePublicNetworkAddressResponseBody {
	return s.Body
}

func (s *ReleasePublicNetworkAddressResponse) SetHeaders(v map[string]*string) *ReleasePublicNetworkAddressResponse {
	s.Headers = v
	return s
}

func (s *ReleasePublicNetworkAddressResponse) SetStatusCode(v int32) *ReleasePublicNetworkAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleasePublicNetworkAddressResponse) SetBody(v *ReleasePublicNetworkAddressResponseBody) *ReleasePublicNetworkAddressResponse {
	s.Body = v
	return s
}

func (s *ReleasePublicNetworkAddressResponse) Validate() error {
	return dara.Validate(s)
}
