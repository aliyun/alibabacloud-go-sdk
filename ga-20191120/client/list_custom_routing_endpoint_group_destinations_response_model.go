// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomRoutingEndpointGroupDestinationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCustomRoutingEndpointGroupDestinationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCustomRoutingEndpointGroupDestinationsResponse
	GetStatusCode() *int32
	SetBody(v *ListCustomRoutingEndpointGroupDestinationsResponseBody) *ListCustomRoutingEndpointGroupDestinationsResponse
	GetBody() *ListCustomRoutingEndpointGroupDestinationsResponseBody
}

type ListCustomRoutingEndpointGroupDestinationsResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCustomRoutingEndpointGroupDestinationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCustomRoutingEndpointGroupDestinationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCustomRoutingEndpointGroupDestinationsResponse) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponse) GetBody() *ListCustomRoutingEndpointGroupDestinationsResponseBody {
	return s.Body
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponse) SetHeaders(v map[string]*string) *ListCustomRoutingEndpointGroupDestinationsResponse {
	s.Headers = v
	return s
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponse) SetStatusCode(v int32) *ListCustomRoutingEndpointGroupDestinationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponse) SetBody(v *ListCustomRoutingEndpointGroupDestinationsResponseBody) *ListCustomRoutingEndpointGroupDestinationsResponse {
	s.Body = v
	return s
}

func (s *ListCustomRoutingEndpointGroupDestinationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
