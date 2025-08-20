// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClustersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBClustersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBClustersResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBClustersResponseBody) *DescribeDBClustersResponse
	GetBody() *DescribeDBClustersResponseBody
}

type DescribeDBClustersResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBClustersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBClustersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClustersResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBClustersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBClustersResponse) GetBody() *DescribeDBClustersResponseBody {
	return s.Body
}

func (s *DescribeDBClustersResponse) SetHeaders(v map[string]*string) *DescribeDBClustersResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClustersResponse) SetStatusCode(v int32) *DescribeDBClustersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClustersResponse) SetBody(v *DescribeDBClustersResponseBody) *DescribeDBClustersResponse {
	s.Body = v
	return s
}

func (s *DescribeDBClustersResponse) Validate() error {
	return dara.Validate(s)
}
