// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBProxyEndpointAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDBProxyEndpointAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDBProxyEndpointAddressResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDBProxyEndpointAddressResponseBody) *DeleteDBProxyEndpointAddressResponse
	GetBody() *DeleteDBProxyEndpointAddressResponseBody
}

type DeleteDBProxyEndpointAddressResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDBProxyEndpointAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDBProxyEndpointAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBProxyEndpointAddressResponse) GoString() string {
	return s.String()
}

func (s *DeleteDBProxyEndpointAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDBProxyEndpointAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDBProxyEndpointAddressResponse) GetBody() *DeleteDBProxyEndpointAddressResponseBody {
	return s.Body
}

func (s *DeleteDBProxyEndpointAddressResponse) SetHeaders(v map[string]*string) *DeleteDBProxyEndpointAddressResponse {
	s.Headers = v
	return s
}

func (s *DeleteDBProxyEndpointAddressResponse) SetStatusCode(v int32) *DeleteDBProxyEndpointAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDBProxyEndpointAddressResponse) SetBody(v *DeleteDBProxyEndpointAddressResponseBody) *DeleteDBProxyEndpointAddressResponse {
	s.Body = v
	return s
}

func (s *DeleteDBProxyEndpointAddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
