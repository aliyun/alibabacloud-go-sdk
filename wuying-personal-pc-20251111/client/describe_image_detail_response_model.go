// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeImageDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeImageDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeImageDetailResponseBody) *DescribeImageDetailResponse
	GetBody() *DescribeImageDetailResponseBody
}

type DescribeImageDetailResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImageDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImageDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeImageDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeImageDetailResponse) GetBody() *DescribeImageDetailResponseBody {
	return s.Body
}

func (s *DescribeImageDetailResponse) SetHeaders(v map[string]*string) *DescribeImageDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageDetailResponse) SetStatusCode(v int32) *DescribeImageDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImageDetailResponse) SetBody(v *DescribeImageDetailResponseBody) *DescribeImageDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeImageDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
