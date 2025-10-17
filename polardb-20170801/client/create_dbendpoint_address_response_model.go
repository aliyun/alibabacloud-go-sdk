// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBEndpointAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDBEndpointAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDBEndpointAddressResponse
	GetStatusCode() *int32
	SetBody(v *CreateDBEndpointAddressResponseBody) *CreateDBEndpointAddressResponse
	GetBody() *CreateDBEndpointAddressResponseBody
}

type CreateDBEndpointAddressResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDBEndpointAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDBEndpointAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDBEndpointAddressResponse) GoString() string {
	return s.String()
}

func (s *CreateDBEndpointAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDBEndpointAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDBEndpointAddressResponse) GetBody() *CreateDBEndpointAddressResponseBody {
	return s.Body
}

func (s *CreateDBEndpointAddressResponse) SetHeaders(v map[string]*string) *CreateDBEndpointAddressResponse {
	s.Headers = v
	return s
}

func (s *CreateDBEndpointAddressResponse) SetStatusCode(v int32) *CreateDBEndpointAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDBEndpointAddressResponse) SetBody(v *CreateDBEndpointAddressResponseBody) *CreateDBEndpointAddressResponse {
	s.Body = v
	return s
}

func (s *CreateDBEndpointAddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
