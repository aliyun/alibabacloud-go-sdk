// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceClusterListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceClusterListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceClusterListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceClusterListResponseBody) *DescribeInstanceClusterListResponse
	GetBody() *DescribeInstanceClusterListResponseBody
}

type DescribeInstanceClusterListResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceClusterListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceClusterListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceClusterListResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceClusterListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceClusterListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceClusterListResponse) GetBody() *DescribeInstanceClusterListResponseBody {
	return s.Body
}

func (s *DescribeInstanceClusterListResponse) SetHeaders(v map[string]*string) *DescribeInstanceClusterListResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceClusterListResponse) SetStatusCode(v int32) *DescribeInstanceClusterListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceClusterListResponse) SetBody(v *DescribeInstanceClusterListResponseBody) *DescribeInstanceClusterListResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceClusterListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
