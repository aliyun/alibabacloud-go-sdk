// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOssObjectDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOssObjectDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOssObjectDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeOssObjectDetailResponseBody) *DescribeOssObjectDetailResponse
	GetBody() *DescribeOssObjectDetailResponseBody
}

type DescribeOssObjectDetailResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOssObjectDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOssObjectDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeOssObjectDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeOssObjectDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOssObjectDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOssObjectDetailResponse) GetBody() *DescribeOssObjectDetailResponseBody {
	return s.Body
}

func (s *DescribeOssObjectDetailResponse) SetHeaders(v map[string]*string) *DescribeOssObjectDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeOssObjectDetailResponse) SetStatusCode(v int32) *DescribeOssObjectDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOssObjectDetailResponse) SetBody(v *DescribeOssObjectDetailResponseBody) *DescribeOssObjectDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeOssObjectDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
