// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClusterListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClusterListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClusterListResponseBody) *DescribeClusterListResponse
	GetBody() *DescribeClusterListResponseBody
}

type DescribeClusterListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClusterListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClusterListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterListResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClusterListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClusterListResponse) GetBody() *DescribeClusterListResponseBody {
	return s.Body
}

func (s *DescribeClusterListResponse) SetHeaders(v map[string]*string) *DescribeClusterListResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterListResponse) SetStatusCode(v int32) *DescribeClusterListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterListResponse) SetBody(v *DescribeClusterListResponseBody) *DescribeClusterListResponse {
	s.Body = v
	return s
}

func (s *DescribeClusterListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
