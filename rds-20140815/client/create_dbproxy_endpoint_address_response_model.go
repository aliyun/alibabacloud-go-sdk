// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBProxyEndpointAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDBProxyEndpointAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDBProxyEndpointAddressResponse
	GetStatusCode() *int32
	SetBody(v *CreateDBProxyEndpointAddressResponseBody) *CreateDBProxyEndpointAddressResponse
	GetBody() *CreateDBProxyEndpointAddressResponseBody
}

type CreateDBProxyEndpointAddressResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDBProxyEndpointAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDBProxyEndpointAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDBProxyEndpointAddressResponse) GoString() string {
	return s.String()
}

func (s *CreateDBProxyEndpointAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDBProxyEndpointAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDBProxyEndpointAddressResponse) GetBody() *CreateDBProxyEndpointAddressResponseBody {
	return s.Body
}

func (s *CreateDBProxyEndpointAddressResponse) SetHeaders(v map[string]*string) *CreateDBProxyEndpointAddressResponse {
	s.Headers = v
	return s
}

func (s *CreateDBProxyEndpointAddressResponse) SetStatusCode(v int32) *CreateDBProxyEndpointAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDBProxyEndpointAddressResponse) SetBody(v *CreateDBProxyEndpointAddressResponseBody) *CreateDBProxyEndpointAddressResponse {
	s.Body = v
	return s
}

func (s *CreateDBProxyEndpointAddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
