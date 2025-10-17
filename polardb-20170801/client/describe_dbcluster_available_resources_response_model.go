// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterAvailableResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBClusterAvailableResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBClusterAvailableResourcesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBClusterAvailableResourcesResponseBody) *DescribeDBClusterAvailableResourcesResponse
	GetBody() *DescribeDBClusterAvailableResourcesResponseBody
}

type DescribeDBClusterAvailableResourcesResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBClusterAvailableResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBClusterAvailableResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterAvailableResourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAvailableResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBClusterAvailableResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBClusterAvailableResourcesResponse) GetBody() *DescribeDBClusterAvailableResourcesResponseBody {
	return s.Body
}

func (s *DescribeDBClusterAvailableResourcesResponse) SetHeaders(v map[string]*string) *DescribeDBClusterAvailableResourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterAvailableResourcesResponse) SetStatusCode(v int32) *DescribeDBClusterAvailableResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterAvailableResourcesResponse) SetBody(v *DescribeDBClusterAvailableResourcesResponseBody) *DescribeDBClusterAvailableResourcesResponse {
	s.Body = v
	return s
}

func (s *DescribeDBClusterAvailableResourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
