// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClustersZonalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBClustersZonalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBClustersZonalResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBClustersZonalResponseBody) *DescribeDBClustersZonalResponse
	GetBody() *DescribeDBClustersZonalResponseBody
}

type DescribeDBClustersZonalResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBClustersZonalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBClustersZonalResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClustersZonalResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersZonalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBClustersZonalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBClustersZonalResponse) GetBody() *DescribeDBClustersZonalResponseBody {
	return s.Body
}

func (s *DescribeDBClustersZonalResponse) SetHeaders(v map[string]*string) *DescribeDBClustersZonalResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClustersZonalResponse) SetStatusCode(v int32) *DescribeDBClustersZonalResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClustersZonalResponse) SetBody(v *DescribeDBClustersZonalResponseBody) *DescribeDBClustersZonalResponse {
	s.Body = v
	return s
}

func (s *DescribeDBClustersZonalResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
