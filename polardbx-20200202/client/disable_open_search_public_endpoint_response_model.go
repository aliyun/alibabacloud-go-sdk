// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableOpenSearchPublicEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableOpenSearchPublicEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableOpenSearchPublicEndpointResponse
	GetStatusCode() *int32
	SetBody(v *DisableOpenSearchPublicEndpointResponseBody) *DisableOpenSearchPublicEndpointResponse
	GetBody() *DisableOpenSearchPublicEndpointResponseBody
}

type DisableOpenSearchPublicEndpointResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableOpenSearchPublicEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableOpenSearchPublicEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableOpenSearchPublicEndpointResponse) GoString() string {
	return s.String()
}

func (s *DisableOpenSearchPublicEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableOpenSearchPublicEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableOpenSearchPublicEndpointResponse) GetBody() *DisableOpenSearchPublicEndpointResponseBody {
	return s.Body
}

func (s *DisableOpenSearchPublicEndpointResponse) SetHeaders(v map[string]*string) *DisableOpenSearchPublicEndpointResponse {
	s.Headers = v
	return s
}

func (s *DisableOpenSearchPublicEndpointResponse) SetStatusCode(v int32) *DisableOpenSearchPublicEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableOpenSearchPublicEndpointResponse) SetBody(v *DisableOpenSearchPublicEndpointResponseBody) *DisableOpenSearchPublicEndpointResponse {
	s.Body = v
	return s
}

func (s *DisableOpenSearchPublicEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
