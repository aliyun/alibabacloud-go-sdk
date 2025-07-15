// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceClusterInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceClusterInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceClusterInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceClusterInfoResponseBody) *DescribeInstanceClusterInfoResponse
	GetBody() *DescribeInstanceClusterInfoResponseBody
}

type DescribeInstanceClusterInfoResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceClusterInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceClusterInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceClusterInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceClusterInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceClusterInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceClusterInfoResponse) GetBody() *DescribeInstanceClusterInfoResponseBody {
	return s.Body
}

func (s *DescribeInstanceClusterInfoResponse) SetHeaders(v map[string]*string) *DescribeInstanceClusterInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceClusterInfoResponse) SetStatusCode(v int32) *DescribeInstanceClusterInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceClusterInfoResponse) SetBody(v *DescribeInstanceClusterInfoResponseBody) *DescribeInstanceClusterInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceClusterInfoResponse) Validate() error {
	return dara.Validate(s)
}
