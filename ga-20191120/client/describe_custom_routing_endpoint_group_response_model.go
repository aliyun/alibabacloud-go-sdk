// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomRoutingEndpointGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCustomRoutingEndpointGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCustomRoutingEndpointGroupResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCustomRoutingEndpointGroupResponseBody) *DescribeCustomRoutingEndpointGroupResponse
	GetBody() *DescribeCustomRoutingEndpointGroupResponseBody
}

type DescribeCustomRoutingEndpointGroupResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCustomRoutingEndpointGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCustomRoutingEndpointGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomRoutingEndpointGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustomRoutingEndpointGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCustomRoutingEndpointGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCustomRoutingEndpointGroupResponse) GetBody() *DescribeCustomRoutingEndpointGroupResponseBody {
	return s.Body
}

func (s *DescribeCustomRoutingEndpointGroupResponse) SetHeaders(v map[string]*string) *DescribeCustomRoutingEndpointGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupResponse) SetStatusCode(v int32) *DescribeCustomRoutingEndpointGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupResponse) SetBody(v *DescribeCustomRoutingEndpointGroupResponseBody) *DescribeCustomRoutingEndpointGroupResponse {
	s.Body = v
	return s
}

func (s *DescribeCustomRoutingEndpointGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
