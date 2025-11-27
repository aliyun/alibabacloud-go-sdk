// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBProxyEndpointAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBProxyEndpointAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBProxyEndpointAddressResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBProxyEndpointAddressResponseBody) *ModifyDBProxyEndpointAddressResponse
	GetBody() *ModifyDBProxyEndpointAddressResponseBody
}

type ModifyDBProxyEndpointAddressResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBProxyEndpointAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBProxyEndpointAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBProxyEndpointAddressResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBProxyEndpointAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBProxyEndpointAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBProxyEndpointAddressResponse) GetBody() *ModifyDBProxyEndpointAddressResponseBody {
	return s.Body
}

func (s *ModifyDBProxyEndpointAddressResponse) SetHeaders(v map[string]*string) *ModifyDBProxyEndpointAddressResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBProxyEndpointAddressResponse) SetStatusCode(v int32) *ModifyDBProxyEndpointAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBProxyEndpointAddressResponse) SetBody(v *ModifyDBProxyEndpointAddressResponseBody) *ModifyDBProxyEndpointAddressResponse {
	s.Body = v
	return s
}

func (s *ModifyDBProxyEndpointAddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
