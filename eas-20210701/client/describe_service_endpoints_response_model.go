// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceEndpointsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeServiceEndpointsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeServiceEndpointsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeServiceEndpointsResponseBody) *DescribeServiceEndpointsResponse
	GetBody() *DescribeServiceEndpointsResponseBody
}

type DescribeServiceEndpointsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeServiceEndpointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeServiceEndpointsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceEndpointsResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceEndpointsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeServiceEndpointsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeServiceEndpointsResponse) GetBody() *DescribeServiceEndpointsResponseBody {
	return s.Body
}

func (s *DescribeServiceEndpointsResponse) SetHeaders(v map[string]*string) *DescribeServiceEndpointsResponse {
	s.Headers = v
	return s
}

func (s *DescribeServiceEndpointsResponse) SetStatusCode(v int32) *DescribeServiceEndpointsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServiceEndpointsResponse) SetBody(v *DescribeServiceEndpointsResponseBody) *DescribeServiceEndpointsResponse {
	s.Body = v
	return s
}

func (s *DescribeServiceEndpointsResponse) Validate() error {
	return dara.Validate(s)
}
