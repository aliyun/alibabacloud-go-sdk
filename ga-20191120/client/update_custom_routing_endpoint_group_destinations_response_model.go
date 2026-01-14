// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomRoutingEndpointGroupDestinationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCustomRoutingEndpointGroupDestinationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCustomRoutingEndpointGroupDestinationsResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCustomRoutingEndpointGroupDestinationsResponseBody) *UpdateCustomRoutingEndpointGroupDestinationsResponse
	GetBody() *UpdateCustomRoutingEndpointGroupDestinationsResponseBody
}

type UpdateCustomRoutingEndpointGroupDestinationsResponse struct {
	Headers    map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCustomRoutingEndpointGroupDestinationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCustomRoutingEndpointGroupDestinationsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomRoutingEndpointGroupDestinationsResponse) GoString() string {
	return s.String()
}

func (s *UpdateCustomRoutingEndpointGroupDestinationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCustomRoutingEndpointGroupDestinationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCustomRoutingEndpointGroupDestinationsResponse) GetBody() *UpdateCustomRoutingEndpointGroupDestinationsResponseBody {
	return s.Body
}

func (s *UpdateCustomRoutingEndpointGroupDestinationsResponse) SetHeaders(v map[string]*string) *UpdateCustomRoutingEndpointGroupDestinationsResponse {
	s.Headers = v
	return s
}

func (s *UpdateCustomRoutingEndpointGroupDestinationsResponse) SetStatusCode(v int32) *UpdateCustomRoutingEndpointGroupDestinationsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCustomRoutingEndpointGroupDestinationsResponse) SetBody(v *UpdateCustomRoutingEndpointGroupDestinationsResponseBody) *UpdateCustomRoutingEndpointGroupDestinationsResponse {
	s.Body = v
	return s
}

func (s *UpdateCustomRoutingEndpointGroupDestinationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
