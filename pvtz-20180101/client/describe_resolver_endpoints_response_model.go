// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResolverEndpointsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeResolverEndpointsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeResolverEndpointsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeResolverEndpointsResponseBody) *DescribeResolverEndpointsResponse
	GetBody() *DescribeResolverEndpointsResponseBody
}

type DescribeResolverEndpointsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeResolverEndpointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeResolverEndpointsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeResolverEndpointsResponse) GoString() string {
	return s.String()
}

func (s *DescribeResolverEndpointsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeResolverEndpointsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeResolverEndpointsResponse) GetBody() *DescribeResolverEndpointsResponseBody {
	return s.Body
}

func (s *DescribeResolverEndpointsResponse) SetHeaders(v map[string]*string) *DescribeResolverEndpointsResponse {
	s.Headers = v
	return s
}

func (s *DescribeResolverEndpointsResponse) SetStatusCode(v int32) *DescribeResolverEndpointsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResolverEndpointsResponse) SetBody(v *DescribeResolverEndpointsResponseBody) *DescribeResolverEndpointsResponse {
	s.Body = v
	return s
}

func (s *DescribeResolverEndpointsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
