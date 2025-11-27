// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBInstanceEndpointAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDBInstanceEndpointAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDBInstanceEndpointAddressResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDBInstanceEndpointAddressResponseBody) *DeleteDBInstanceEndpointAddressResponse
	GetBody() *DeleteDBInstanceEndpointAddressResponseBody
}

type DeleteDBInstanceEndpointAddressResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDBInstanceEndpointAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDBInstanceEndpointAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBInstanceEndpointAddressResponse) GoString() string {
	return s.String()
}

func (s *DeleteDBInstanceEndpointAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDBInstanceEndpointAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDBInstanceEndpointAddressResponse) GetBody() *DeleteDBInstanceEndpointAddressResponseBody {
	return s.Body
}

func (s *DeleteDBInstanceEndpointAddressResponse) SetHeaders(v map[string]*string) *DeleteDBInstanceEndpointAddressResponse {
	s.Headers = v
	return s
}

func (s *DeleteDBInstanceEndpointAddressResponse) SetStatusCode(v int32) *DeleteDBInstanceEndpointAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDBInstanceEndpointAddressResponse) SetBody(v *DeleteDBInstanceEndpointAddressResponseBody) *DeleteDBInstanceEndpointAddressResponse {
	s.Body = v
	return s
}

func (s *DeleteDBInstanceEndpointAddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
