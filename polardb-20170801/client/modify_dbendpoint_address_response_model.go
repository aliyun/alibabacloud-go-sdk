// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBEndpointAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBEndpointAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBEndpointAddressResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBEndpointAddressResponseBody) *ModifyDBEndpointAddressResponse
	GetBody() *ModifyDBEndpointAddressResponseBody
}

type ModifyDBEndpointAddressResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBEndpointAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBEndpointAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBEndpointAddressResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBEndpointAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBEndpointAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBEndpointAddressResponse) GetBody() *ModifyDBEndpointAddressResponseBody {
	return s.Body
}

func (s *ModifyDBEndpointAddressResponse) SetHeaders(v map[string]*string) *ModifyDBEndpointAddressResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBEndpointAddressResponse) SetStatusCode(v int32) *ModifyDBEndpointAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBEndpointAddressResponse) SetBody(v *ModifyDBEndpointAddressResponseBody) *ModifyDBEndpointAddressResponse {
	s.Body = v
	return s
}

func (s *ModifyDBEndpointAddressResponse) Validate() error {
	return dara.Validate(s)
}
