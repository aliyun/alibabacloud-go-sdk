// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBInstanceEndpointAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDBInstanceEndpointAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDBInstanceEndpointAddressResponse
	GetStatusCode() *int32
	SetBody(v *CreateDBInstanceEndpointAddressResponseBody) *CreateDBInstanceEndpointAddressResponse
	GetBody() *CreateDBInstanceEndpointAddressResponseBody
}

type CreateDBInstanceEndpointAddressResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDBInstanceEndpointAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDBInstanceEndpointAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstanceEndpointAddressResponse) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceEndpointAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDBInstanceEndpointAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDBInstanceEndpointAddressResponse) GetBody() *CreateDBInstanceEndpointAddressResponseBody {
	return s.Body
}

func (s *CreateDBInstanceEndpointAddressResponse) SetHeaders(v map[string]*string) *CreateDBInstanceEndpointAddressResponse {
	s.Headers = v
	return s
}

func (s *CreateDBInstanceEndpointAddressResponse) SetStatusCode(v int32) *CreateDBInstanceEndpointAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDBInstanceEndpointAddressResponse) SetBody(v *CreateDBInstanceEndpointAddressResponseBody) *CreateDBInstanceEndpointAddressResponse {
	s.Body = v
	return s
}

func (s *CreateDBInstanceEndpointAddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
