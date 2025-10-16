// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrivateDnsEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePrivateDnsEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePrivateDnsEndpointResponse
	GetStatusCode() *int32
	SetBody(v *DeletePrivateDnsEndpointResponseBody) *DeletePrivateDnsEndpointResponse
	GetBody() *DeletePrivateDnsEndpointResponseBody
}

type DeletePrivateDnsEndpointResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePrivateDnsEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePrivateDnsEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePrivateDnsEndpointResponse) GoString() string {
	return s.String()
}

func (s *DeletePrivateDnsEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePrivateDnsEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePrivateDnsEndpointResponse) GetBody() *DeletePrivateDnsEndpointResponseBody {
	return s.Body
}

func (s *DeletePrivateDnsEndpointResponse) SetHeaders(v map[string]*string) *DeletePrivateDnsEndpointResponse {
	s.Headers = v
	return s
}

func (s *DeletePrivateDnsEndpointResponse) SetStatusCode(v int32) *DeletePrivateDnsEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePrivateDnsEndpointResponse) SetBody(v *DeletePrivateDnsEndpointResponseBody) *DeletePrivateDnsEndpointResponse {
	s.Body = v
	return s
}

func (s *DeletePrivateDnsEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
