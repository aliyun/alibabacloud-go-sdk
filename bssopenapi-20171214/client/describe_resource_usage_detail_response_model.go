// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceUsageDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeResourceUsageDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeResourceUsageDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeResourceUsageDetailResponseBody) *DescribeResourceUsageDetailResponse
	GetBody() *DescribeResourceUsageDetailResponseBody
}

type DescribeResourceUsageDetailResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeResourceUsageDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeResourceUsageDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceUsageDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourceUsageDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeResourceUsageDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeResourceUsageDetailResponse) GetBody() *DescribeResourceUsageDetailResponseBody {
	return s.Body
}

func (s *DescribeResourceUsageDetailResponse) SetHeaders(v map[string]*string) *DescribeResourceUsageDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourceUsageDetailResponse) SetStatusCode(v int32) *DescribeResourceUsageDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResourceUsageDetailResponse) SetBody(v *DescribeResourceUsageDetailResponseBody) *DescribeResourceUsageDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeResourceUsageDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
