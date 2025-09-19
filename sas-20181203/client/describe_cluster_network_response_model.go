// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterNetworkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClusterNetworkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClusterNetworkResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClusterNetworkResponseBody) *DescribeClusterNetworkResponse
	GetBody() *DescribeClusterNetworkResponseBody
}

type DescribeClusterNetworkResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClusterNetworkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClusterNetworkResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNetworkResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterNetworkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClusterNetworkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClusterNetworkResponse) GetBody() *DescribeClusterNetworkResponseBody {
	return s.Body
}

func (s *DescribeClusterNetworkResponse) SetHeaders(v map[string]*string) *DescribeClusterNetworkResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterNetworkResponse) SetStatusCode(v int32) *DescribeClusterNetworkResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterNetworkResponse) SetBody(v *DescribeClusterNetworkResponseBody) *DescribeClusterNetworkResponse {
	s.Body = v
	return s
}

func (s *DescribeClusterNetworkResponse) Validate() error {
	return dara.Validate(s)
}
