// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceDropConnectionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceDropConnectionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceDropConnectionsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceDropConnectionsResponseBody) *DescribeInstanceDropConnectionsResponse
	GetBody() *DescribeInstanceDropConnectionsResponseBody
}

type DescribeInstanceDropConnectionsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceDropConnectionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceDropConnectionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceDropConnectionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDropConnectionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceDropConnectionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceDropConnectionsResponse) GetBody() *DescribeInstanceDropConnectionsResponseBody {
	return s.Body
}

func (s *DescribeInstanceDropConnectionsResponse) SetHeaders(v map[string]*string) *DescribeInstanceDropConnectionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceDropConnectionsResponse) SetStatusCode(v int32) *DescribeInstanceDropConnectionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceDropConnectionsResponse) SetBody(v *DescribeInstanceDropConnectionsResponseBody) *DescribeInstanceDropConnectionsResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceDropConnectionsResponse) Validate() error {
	return dara.Validate(s)
}
