// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEventDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEventDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEventDetailResponseBody) *DescribeEventDetailResponse
	GetBody() *DescribeEventDetailResponseBody
}

type DescribeEventDetailResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEventDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEventDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeEventDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEventDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEventDetailResponse) GetBody() *DescribeEventDetailResponseBody {
	return s.Body
}

func (s *DescribeEventDetailResponse) SetHeaders(v map[string]*string) *DescribeEventDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeEventDetailResponse) SetStatusCode(v int32) *DescribeEventDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEventDetailResponse) SetBody(v *DescribeEventDetailResponseBody) *DescribeEventDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeEventDetailResponse) Validate() error {
	return dara.Validate(s)
}
