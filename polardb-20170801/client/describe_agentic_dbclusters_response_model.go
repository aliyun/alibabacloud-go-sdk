// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAgenticDBClustersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAgenticDBClustersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAgenticDBClustersResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAgenticDBClustersResponseBody) *DescribeAgenticDBClustersResponse
	GetBody() *DescribeAgenticDBClustersResponseBody
}

type DescribeAgenticDBClustersResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAgenticDBClustersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAgenticDBClustersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticDBClustersResponse) GoString() string {
	return s.String()
}

func (s *DescribeAgenticDBClustersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAgenticDBClustersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAgenticDBClustersResponse) GetBody() *DescribeAgenticDBClustersResponseBody {
	return s.Body
}

func (s *DescribeAgenticDBClustersResponse) SetHeaders(v map[string]*string) *DescribeAgenticDBClustersResponse {
	s.Headers = v
	return s
}

func (s *DescribeAgenticDBClustersResponse) SetStatusCode(v int32) *DescribeAgenticDBClustersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAgenticDBClustersResponse) SetBody(v *DescribeAgenticDBClustersResponseBody) *DescribeAgenticDBClustersResponse {
	s.Body = v
	return s
}

func (s *DescribeAgenticDBClustersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
