// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationEndpointAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateApplicationEndpointAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateApplicationEndpointAddressResponse
	GetStatusCode() *int32
	SetBody(v *CreateApplicationEndpointAddressResponseBody) *CreateApplicationEndpointAddressResponse
	GetBody() *CreateApplicationEndpointAddressResponseBody
}

type CreateApplicationEndpointAddressResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateApplicationEndpointAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateApplicationEndpointAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationEndpointAddressResponse) GoString() string {
	return s.String()
}

func (s *CreateApplicationEndpointAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateApplicationEndpointAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateApplicationEndpointAddressResponse) GetBody() *CreateApplicationEndpointAddressResponseBody {
	return s.Body
}

func (s *CreateApplicationEndpointAddressResponse) SetHeaders(v map[string]*string) *CreateApplicationEndpointAddressResponse {
	s.Headers = v
	return s
}

func (s *CreateApplicationEndpointAddressResponse) SetStatusCode(v int32) *CreateApplicationEndpointAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateApplicationEndpointAddressResponse) SetBody(v *CreateApplicationEndpointAddressResponseBody) *CreateApplicationEndpointAddressResponse {
	s.Body = v
	return s
}

func (s *CreateApplicationEndpointAddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
