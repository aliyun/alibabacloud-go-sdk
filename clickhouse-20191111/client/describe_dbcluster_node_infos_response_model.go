// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterNodeInfosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBClusterNodeInfosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBClusterNodeInfosResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBClusterNodeInfosResponseBody) *DescribeDBClusterNodeInfosResponse
	GetBody() *DescribeDBClusterNodeInfosResponseBody
}

type DescribeDBClusterNodeInfosResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBClusterNodeInfosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBClusterNodeInfosResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterNodeInfosResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterNodeInfosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBClusterNodeInfosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBClusterNodeInfosResponse) GetBody() *DescribeDBClusterNodeInfosResponseBody {
	return s.Body
}

func (s *DescribeDBClusterNodeInfosResponse) SetHeaders(v map[string]*string) *DescribeDBClusterNodeInfosResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterNodeInfosResponse) SetStatusCode(v int32) *DescribeDBClusterNodeInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterNodeInfosResponse) SetBody(v *DescribeDBClusterNodeInfosResponseBody) *DescribeDBClusterNodeInfosResponse {
	s.Body = v
	return s
}

func (s *DescribeDBClusterNodeInfosResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
