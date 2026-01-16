// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClusterDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClusterDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClusterDetailResponseBody) *DescribeClusterDetailResponse
	GetBody() *DescribeClusterDetailResponseBody
}

type DescribeClusterDetailResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClusterDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClusterDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClusterDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClusterDetailResponse) GetBody() *DescribeClusterDetailResponseBody {
	return s.Body
}

func (s *DescribeClusterDetailResponse) SetHeaders(v map[string]*string) *DescribeClusterDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterDetailResponse) SetStatusCode(v int32) *DescribeClusterDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterDetailResponse) SetBody(v *DescribeClusterDetailResponseBody) *DescribeClusterDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeClusterDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
