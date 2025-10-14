// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterKubeConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClusterKubeConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClusterKubeConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClusterKubeConfigResponseBody) *DescribeClusterKubeConfigResponse
	GetBody() *DescribeClusterKubeConfigResponseBody
}

type DescribeClusterKubeConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClusterKubeConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClusterKubeConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterKubeConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterKubeConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClusterKubeConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClusterKubeConfigResponse) GetBody() *DescribeClusterKubeConfigResponseBody {
	return s.Body
}

func (s *DescribeClusterKubeConfigResponse) SetHeaders(v map[string]*string) *DescribeClusterKubeConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterKubeConfigResponse) SetStatusCode(v int32) *DescribeClusterKubeConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterKubeConfigResponse) SetBody(v *DescribeClusterKubeConfigResponseBody) *DescribeClusterKubeConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeClusterKubeConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
