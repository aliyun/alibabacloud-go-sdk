// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomRoutingEndpointGroupDestinationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCustomRoutingEndpointGroupDestinationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCustomRoutingEndpointGroupDestinationsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCustomRoutingEndpointGroupDestinationsResponseBody) *DescribeCustomRoutingEndpointGroupDestinationsResponse
	GetBody() *DescribeCustomRoutingEndpointGroupDestinationsResponseBody
}

type DescribeCustomRoutingEndpointGroupDestinationsResponse struct {
	Headers    map[string]*string                                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCustomRoutingEndpointGroupDestinationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCustomRoutingEndpointGroupDestinationsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomRoutingEndpointGroupDestinationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsResponse) GetBody() *DescribeCustomRoutingEndpointGroupDestinationsResponseBody {
	return s.Body
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsResponse) SetHeaders(v map[string]*string) *DescribeCustomRoutingEndpointGroupDestinationsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsResponse) SetStatusCode(v int32) *DescribeCustomRoutingEndpointGroupDestinationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsResponse) SetBody(v *DescribeCustomRoutingEndpointGroupDestinationsResponseBody) *DescribeCustomRoutingEndpointGroupDestinationsResponse {
	s.Body = v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupDestinationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
