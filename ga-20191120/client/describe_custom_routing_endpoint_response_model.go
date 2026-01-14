// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomRoutingEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCustomRoutingEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCustomRoutingEndpointResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCustomRoutingEndpointResponseBody) *DescribeCustomRoutingEndpointResponse
	GetBody() *DescribeCustomRoutingEndpointResponseBody
}

type DescribeCustomRoutingEndpointResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCustomRoutingEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCustomRoutingEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomRoutingEndpointResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustomRoutingEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCustomRoutingEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCustomRoutingEndpointResponse) GetBody() *DescribeCustomRoutingEndpointResponseBody {
	return s.Body
}

func (s *DescribeCustomRoutingEndpointResponse) SetHeaders(v map[string]*string) *DescribeCustomRoutingEndpointResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustomRoutingEndpointResponse) SetStatusCode(v int32) *DescribeCustomRoutingEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCustomRoutingEndpointResponse) SetBody(v *DescribeCustomRoutingEndpointResponseBody) *DescribeCustomRoutingEndpointResponse {
	s.Body = v
	return s
}

func (s *DescribeCustomRoutingEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
