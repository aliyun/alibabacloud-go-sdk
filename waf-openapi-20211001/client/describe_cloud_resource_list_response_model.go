// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudResourceListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCloudResourceListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCloudResourceListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCloudResourceListResponseBody) *DescribeCloudResourceListResponse
	GetBody() *DescribeCloudResourceListResponseBody
}

type DescribeCloudResourceListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudResourceListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudResourceListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudResourceListResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudResourceListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCloudResourceListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCloudResourceListResponse) GetBody() *DescribeCloudResourceListResponseBody {
	return s.Body
}

func (s *DescribeCloudResourceListResponse) SetHeaders(v map[string]*string) *DescribeCloudResourceListResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudResourceListResponse) SetStatusCode(v int32) *DescribeCloudResourceListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudResourceListResponse) SetBody(v *DescribeCloudResourceListResponseBody) *DescribeCloudResourceListResponse {
	s.Body = v
	return s
}

func (s *DescribeCloudResourceListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
