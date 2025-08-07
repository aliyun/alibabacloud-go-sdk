// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBEndpointAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDBEndpointAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDBEndpointAddressResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDBEndpointAddressResponseBody) *DeleteDBEndpointAddressResponse
	GetBody() *DeleteDBEndpointAddressResponseBody
}

type DeleteDBEndpointAddressResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDBEndpointAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDBEndpointAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBEndpointAddressResponse) GoString() string {
	return s.String()
}

func (s *DeleteDBEndpointAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDBEndpointAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDBEndpointAddressResponse) GetBody() *DeleteDBEndpointAddressResponseBody {
	return s.Body
}

func (s *DeleteDBEndpointAddressResponse) SetHeaders(v map[string]*string) *DeleteDBEndpointAddressResponse {
	s.Headers = v
	return s
}

func (s *DeleteDBEndpointAddressResponse) SetStatusCode(v int32) *DeleteDBEndpointAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDBEndpointAddressResponse) SetBody(v *DeleteDBEndpointAddressResponseBody) *DeleteDBEndpointAddressResponse {
	s.Body = v
	return s
}

func (s *DeleteDBEndpointAddressResponse) Validate() error {
	return dara.Validate(s)
}
