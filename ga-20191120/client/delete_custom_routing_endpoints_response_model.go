// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomRoutingEndpointsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCustomRoutingEndpointsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCustomRoutingEndpointsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCustomRoutingEndpointsResponseBody) *DeleteCustomRoutingEndpointsResponse
	GetBody() *DeleteCustomRoutingEndpointsResponseBody
}

type DeleteCustomRoutingEndpointsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomRoutingEndpointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomRoutingEndpointsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomRoutingEndpointsResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomRoutingEndpointsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCustomRoutingEndpointsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCustomRoutingEndpointsResponse) GetBody() *DeleteCustomRoutingEndpointsResponseBody {
	return s.Body
}

func (s *DeleteCustomRoutingEndpointsResponse) SetHeaders(v map[string]*string) *DeleteCustomRoutingEndpointsResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomRoutingEndpointsResponse) SetStatusCode(v int32) *DeleteCustomRoutingEndpointsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomRoutingEndpointsResponse) SetBody(v *DeleteCustomRoutingEndpointsResponseBody) *DeleteCustomRoutingEndpointsResponse {
	s.Body = v
	return s
}

func (s *DeleteCustomRoutingEndpointsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
