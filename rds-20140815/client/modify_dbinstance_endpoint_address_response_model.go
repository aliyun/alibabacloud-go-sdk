// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceEndpointAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBInstanceEndpointAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBInstanceEndpointAddressResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBInstanceEndpointAddressResponseBody) *ModifyDBInstanceEndpointAddressResponse
	GetBody() *ModifyDBInstanceEndpointAddressResponseBody
}

type ModifyDBInstanceEndpointAddressResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBInstanceEndpointAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBInstanceEndpointAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceEndpointAddressResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceEndpointAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBInstanceEndpointAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBInstanceEndpointAddressResponse) GetBody() *ModifyDBInstanceEndpointAddressResponseBody {
	return s.Body
}

func (s *ModifyDBInstanceEndpointAddressResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceEndpointAddressResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceEndpointAddressResponse) SetStatusCode(v int32) *ModifyDBInstanceEndpointAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBInstanceEndpointAddressResponse) SetBody(v *ModifyDBInstanceEndpointAddressResponseBody) *ModifyDBInstanceEndpointAddressResponse {
	s.Body = v
	return s
}

func (s *ModifyDBInstanceEndpointAddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
