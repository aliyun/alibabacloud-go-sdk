// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInternetOpenDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInternetOpenDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInternetOpenDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInternetOpenDetailResponseBody) *DescribeInternetOpenDetailResponse
	GetBody() *DescribeInternetOpenDetailResponseBody
}

type DescribeInternetOpenDetailResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInternetOpenDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInternetOpenDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInternetOpenDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeInternetOpenDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInternetOpenDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInternetOpenDetailResponse) GetBody() *DescribeInternetOpenDetailResponseBody {
	return s.Body
}

func (s *DescribeInternetOpenDetailResponse) SetHeaders(v map[string]*string) *DescribeInternetOpenDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeInternetOpenDetailResponse) SetStatusCode(v int32) *DescribeInternetOpenDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInternetOpenDetailResponse) SetBody(v *DescribeInternetOpenDetailResponseBody) *DescribeInternetOpenDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeInternetOpenDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
