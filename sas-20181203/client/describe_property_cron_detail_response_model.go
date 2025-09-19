// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePropertyCronDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePropertyCronDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePropertyCronDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribePropertyCronDetailResponseBody) *DescribePropertyCronDetailResponse
	GetBody() *DescribePropertyCronDetailResponseBody
}

type DescribePropertyCronDetailResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePropertyCronDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePropertyCronDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyCronDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribePropertyCronDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePropertyCronDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePropertyCronDetailResponse) GetBody() *DescribePropertyCronDetailResponseBody {
	return s.Body
}

func (s *DescribePropertyCronDetailResponse) SetHeaders(v map[string]*string) *DescribePropertyCronDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribePropertyCronDetailResponse) SetStatusCode(v int32) *DescribePropertyCronDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePropertyCronDetailResponse) SetBody(v *DescribePropertyCronDetailResponseBody) *DescribePropertyCronDetailResponse {
	s.Body = v
	return s
}

func (s *DescribePropertyCronDetailResponse) Validate() error {
	return dara.Validate(s)
}
