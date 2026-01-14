// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomRoutingEndpointGroupDestinationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCustomRoutingEndpointGroupDestinationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCustomRoutingEndpointGroupDestinationsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCustomRoutingEndpointGroupDestinationsResponseBody) *DeleteCustomRoutingEndpointGroupDestinationsResponse
	GetBody() *DeleteCustomRoutingEndpointGroupDestinationsResponseBody
}

type DeleteCustomRoutingEndpointGroupDestinationsResponse struct {
	Headers    map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomRoutingEndpointGroupDestinationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomRoutingEndpointGroupDestinationsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomRoutingEndpointGroupDestinationsResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomRoutingEndpointGroupDestinationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCustomRoutingEndpointGroupDestinationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCustomRoutingEndpointGroupDestinationsResponse) GetBody() *DeleteCustomRoutingEndpointGroupDestinationsResponseBody {
	return s.Body
}

func (s *DeleteCustomRoutingEndpointGroupDestinationsResponse) SetHeaders(v map[string]*string) *DeleteCustomRoutingEndpointGroupDestinationsResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomRoutingEndpointGroupDestinationsResponse) SetStatusCode(v int32) *DeleteCustomRoutingEndpointGroupDestinationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomRoutingEndpointGroupDestinationsResponse) SetBody(v *DeleteCustomRoutingEndpointGroupDestinationsResponseBody) *DeleteCustomRoutingEndpointGroupDestinationsResponse {
	s.Body = v
	return s
}

func (s *DeleteCustomRoutingEndpointGroupDestinationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
