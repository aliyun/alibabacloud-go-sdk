// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExposedInstanceDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeExposedInstanceDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeExposedInstanceDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeExposedInstanceDetailResponseBody) *DescribeExposedInstanceDetailResponse
	GetBody() *DescribeExposedInstanceDetailResponseBody
}

type DescribeExposedInstanceDetailResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExposedInstanceDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExposedInstanceDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeExposedInstanceDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeExposedInstanceDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeExposedInstanceDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeExposedInstanceDetailResponse) GetBody() *DescribeExposedInstanceDetailResponseBody {
	return s.Body
}

func (s *DescribeExposedInstanceDetailResponse) SetHeaders(v map[string]*string) *DescribeExposedInstanceDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeExposedInstanceDetailResponse) SetStatusCode(v int32) *DescribeExposedInstanceDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExposedInstanceDetailResponse) SetBody(v *DescribeExposedInstanceDetailResponseBody) *DescribeExposedInstanceDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeExposedInstanceDetailResponse) Validate() error {
	return dara.Validate(s)
}
