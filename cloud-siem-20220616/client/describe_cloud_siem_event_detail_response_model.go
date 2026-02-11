// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudSiemEventDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCloudSiemEventDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCloudSiemEventDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCloudSiemEventDetailResponseBody) *DescribeCloudSiemEventDetailResponse
	GetBody() *DescribeCloudSiemEventDetailResponseBody
}

type DescribeCloudSiemEventDetailResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudSiemEventDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudSiemEventDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudSiemEventDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemEventDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCloudSiemEventDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCloudSiemEventDetailResponse) GetBody() *DescribeCloudSiemEventDetailResponseBody {
	return s.Body
}

func (s *DescribeCloudSiemEventDetailResponse) SetHeaders(v map[string]*string) *DescribeCloudSiemEventDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudSiemEventDetailResponse) SetStatusCode(v int32) *DescribeCloudSiemEventDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudSiemEventDetailResponse) SetBody(v *DescribeCloudSiemEventDetailResponseBody) *DescribeCloudSiemEventDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeCloudSiemEventDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
