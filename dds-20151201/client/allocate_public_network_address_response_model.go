// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocatePublicNetworkAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AllocatePublicNetworkAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AllocatePublicNetworkAddressResponse
	GetStatusCode() *int32
	SetBody(v *AllocatePublicNetworkAddressResponseBody) *AllocatePublicNetworkAddressResponse
	GetBody() *AllocatePublicNetworkAddressResponseBody
}

type AllocatePublicNetworkAddressResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AllocatePublicNetworkAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AllocatePublicNetworkAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s AllocatePublicNetworkAddressResponse) GoString() string {
	return s.String()
}

func (s *AllocatePublicNetworkAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AllocatePublicNetworkAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AllocatePublicNetworkAddressResponse) GetBody() *AllocatePublicNetworkAddressResponseBody {
	return s.Body
}

func (s *AllocatePublicNetworkAddressResponse) SetHeaders(v map[string]*string) *AllocatePublicNetworkAddressResponse {
	s.Headers = v
	return s
}

func (s *AllocatePublicNetworkAddressResponse) SetStatusCode(v int32) *AllocatePublicNetworkAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *AllocatePublicNetworkAddressResponse) SetBody(v *AllocatePublicNetworkAddressResponseBody) *AllocatePublicNetworkAddressResponse {
	s.Body = v
	return s
}

func (s *AllocatePublicNetworkAddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
