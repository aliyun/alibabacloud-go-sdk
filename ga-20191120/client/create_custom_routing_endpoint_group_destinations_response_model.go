// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomRoutingEndpointGroupDestinationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCustomRoutingEndpointGroupDestinationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCustomRoutingEndpointGroupDestinationsResponse
	GetStatusCode() *int32
	SetBody(v *CreateCustomRoutingEndpointGroupDestinationsResponseBody) *CreateCustomRoutingEndpointGroupDestinationsResponse
	GetBody() *CreateCustomRoutingEndpointGroupDestinationsResponseBody
}

type CreateCustomRoutingEndpointGroupDestinationsResponse struct {
	Headers    map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCustomRoutingEndpointGroupDestinationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCustomRoutingEndpointGroupDestinationsResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomRoutingEndpointGroupDestinationsResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomRoutingEndpointGroupDestinationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCustomRoutingEndpointGroupDestinationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCustomRoutingEndpointGroupDestinationsResponse) GetBody() *CreateCustomRoutingEndpointGroupDestinationsResponseBody {
	return s.Body
}

func (s *CreateCustomRoutingEndpointGroupDestinationsResponse) SetHeaders(v map[string]*string) *CreateCustomRoutingEndpointGroupDestinationsResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomRoutingEndpointGroupDestinationsResponse) SetStatusCode(v int32) *CreateCustomRoutingEndpointGroupDestinationsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustomRoutingEndpointGroupDestinationsResponse) SetBody(v *CreateCustomRoutingEndpointGroupDestinationsResponseBody) *CreateCustomRoutingEndpointGroupDestinationsResponse {
	s.Body = v
	return s
}

func (s *CreateCustomRoutingEndpointGroupDestinationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
