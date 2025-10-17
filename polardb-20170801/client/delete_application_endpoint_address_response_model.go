// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApplicationEndpointAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteApplicationEndpointAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteApplicationEndpointAddressResponse
	GetStatusCode() *int32
	SetBody(v *DeleteApplicationEndpointAddressResponseBody) *DeleteApplicationEndpointAddressResponse
	GetBody() *DeleteApplicationEndpointAddressResponseBody
}

type DeleteApplicationEndpointAddressResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteApplicationEndpointAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteApplicationEndpointAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteApplicationEndpointAddressResponse) GoString() string {
	return s.String()
}

func (s *DeleteApplicationEndpointAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteApplicationEndpointAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteApplicationEndpointAddressResponse) GetBody() *DeleteApplicationEndpointAddressResponseBody {
	return s.Body
}

func (s *DeleteApplicationEndpointAddressResponse) SetHeaders(v map[string]*string) *DeleteApplicationEndpointAddressResponse {
	s.Headers = v
	return s
}

func (s *DeleteApplicationEndpointAddressResponse) SetStatusCode(v int32) *DeleteApplicationEndpointAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteApplicationEndpointAddressResponse) SetBody(v *DeleteApplicationEndpointAddressResponseBody) *DeleteApplicationEndpointAddressResponse {
	s.Body = v
	return s
}

func (s *DeleteApplicationEndpointAddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
