// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterInfoListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClusterInfoListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClusterInfoListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClusterInfoListResponseBody) *DescribeClusterInfoListResponse
	GetBody() *DescribeClusterInfoListResponseBody
}

type DescribeClusterInfoListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClusterInfoListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClusterInfoListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterInfoListResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterInfoListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClusterInfoListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClusterInfoListResponse) GetBody() *DescribeClusterInfoListResponseBody {
	return s.Body
}

func (s *DescribeClusterInfoListResponse) SetHeaders(v map[string]*string) *DescribeClusterInfoListResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterInfoListResponse) SetStatusCode(v int32) *DescribeClusterInfoListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterInfoListResponse) SetBody(v *DescribeClusterInfoListResponseBody) *DescribeClusterInfoListResponse {
	s.Body = v
	return s
}

func (s *DescribeClusterInfoListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
