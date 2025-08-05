// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePostpayTrafficDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePostpayTrafficDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePostpayTrafficDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribePostpayTrafficDetailResponseBody) *DescribePostpayTrafficDetailResponse
	GetBody() *DescribePostpayTrafficDetailResponseBody
}

type DescribePostpayTrafficDetailResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePostpayTrafficDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePostpayTrafficDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePostpayTrafficDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribePostpayTrafficDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePostpayTrafficDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePostpayTrafficDetailResponse) GetBody() *DescribePostpayTrafficDetailResponseBody {
	return s.Body
}

func (s *DescribePostpayTrafficDetailResponse) SetHeaders(v map[string]*string) *DescribePostpayTrafficDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribePostpayTrafficDetailResponse) SetStatusCode(v int32) *DescribePostpayTrafficDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePostpayTrafficDetailResponse) SetBody(v *DescribePostpayTrafficDetailResponseBody) *DescribePostpayTrafficDetailResponse {
	s.Body = v
	return s
}

func (s *DescribePostpayTrafficDetailResponse) Validate() error {
	return dara.Validate(s)
}
