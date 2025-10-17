// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterEndpointsZonalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBClusterEndpointsZonalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBClusterEndpointsZonalResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBClusterEndpointsZonalResponseBody) *DescribeDBClusterEndpointsZonalResponse
	GetBody() *DescribeDBClusterEndpointsZonalResponseBody
}

type DescribeDBClusterEndpointsZonalResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBClusterEndpointsZonalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBClusterEndpointsZonalResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterEndpointsZonalResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterEndpointsZonalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBClusterEndpointsZonalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBClusterEndpointsZonalResponse) GetBody() *DescribeDBClusterEndpointsZonalResponseBody {
	return s.Body
}

func (s *DescribeDBClusterEndpointsZonalResponse) SetHeaders(v map[string]*string) *DescribeDBClusterEndpointsZonalResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterEndpointsZonalResponse) SetStatusCode(v int32) *DescribeDBClusterEndpointsZonalResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterEndpointsZonalResponse) SetBody(v *DescribeDBClusterEndpointsZonalResponseBody) *DescribeDBClusterEndpointsZonalResponse {
	s.Body = v
	return s
}

func (s *DescribeDBClusterEndpointsZonalResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
