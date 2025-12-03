// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceNewConnectionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceNewConnectionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceNewConnectionsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceNewConnectionsResponseBody) *DescribeInstanceNewConnectionsResponse
	GetBody() *DescribeInstanceNewConnectionsResponseBody
}

type DescribeInstanceNewConnectionsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceNewConnectionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceNewConnectionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceNewConnectionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceNewConnectionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceNewConnectionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceNewConnectionsResponse) GetBody() *DescribeInstanceNewConnectionsResponseBody {
	return s.Body
}

func (s *DescribeInstanceNewConnectionsResponse) SetHeaders(v map[string]*string) *DescribeInstanceNewConnectionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceNewConnectionsResponse) SetStatusCode(v int32) *DescribeInstanceNewConnectionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceNewConnectionsResponse) SetBody(v *DescribeInstanceNewConnectionsResponseBody) *DescribeInstanceNewConnectionsResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceNewConnectionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
