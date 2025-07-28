// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResolverEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeResolverEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeResolverEndpointResponse
	GetStatusCode() *int32
	SetBody(v *DescribeResolverEndpointResponseBody) *DescribeResolverEndpointResponse
	GetBody() *DescribeResolverEndpointResponseBody
}

type DescribeResolverEndpointResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeResolverEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeResolverEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeResolverEndpointResponse) GoString() string {
	return s.String()
}

func (s *DescribeResolverEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeResolverEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeResolverEndpointResponse) GetBody() *DescribeResolverEndpointResponseBody {
	return s.Body
}

func (s *DescribeResolverEndpointResponse) SetHeaders(v map[string]*string) *DescribeResolverEndpointResponse {
	s.Headers = v
	return s
}

func (s *DescribeResolverEndpointResponse) SetStatusCode(v int32) *DescribeResolverEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResolverEndpointResponse) SetBody(v *DescribeResolverEndpointResponseBody) *DescribeResolverEndpointResponse {
	s.Body = v
	return s
}

func (s *DescribeResolverEndpointResponse) Validate() error {
	return dara.Validate(s)
}
