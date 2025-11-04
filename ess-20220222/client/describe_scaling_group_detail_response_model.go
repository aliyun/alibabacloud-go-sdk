// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScalingGroupDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeScalingGroupDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeScalingGroupDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeScalingGroupDetailResponseBody) *DescribeScalingGroupDetailResponse
	GetBody() *DescribeScalingGroupDetailResponseBody
}

type DescribeScalingGroupDetailResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeScalingGroupDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeScalingGroupDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingGroupDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeScalingGroupDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeScalingGroupDetailResponse) GetBody() *DescribeScalingGroupDetailResponseBody {
	return s.Body
}

func (s *DescribeScalingGroupDetailResponse) SetHeaders(v map[string]*string) *DescribeScalingGroupDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeScalingGroupDetailResponse) SetStatusCode(v int32) *DescribeScalingGroupDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScalingGroupDetailResponse) SetBody(v *DescribeScalingGroupDetailResponseBody) *DescribeScalingGroupDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeScalingGroupDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
