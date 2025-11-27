// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBProxyEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBProxyEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBProxyEndpointResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBProxyEndpointResponseBody) *DescribeDBProxyEndpointResponse
	GetBody() *DescribeDBProxyEndpointResponseBody
}

type DescribeDBProxyEndpointResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBProxyEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBProxyEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBProxyEndpointResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBProxyEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBProxyEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBProxyEndpointResponse) GetBody() *DescribeDBProxyEndpointResponseBody {
	return s.Body
}

func (s *DescribeDBProxyEndpointResponse) SetHeaders(v map[string]*string) *DescribeDBProxyEndpointResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBProxyEndpointResponse) SetStatusCode(v int32) *DescribeDBProxyEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBProxyEndpointResponse) SetBody(v *DescribeDBProxyEndpointResponseBody) *DescribeDBProxyEndpointResponse {
	s.Body = v
	return s
}

func (s *DescribeDBProxyEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
