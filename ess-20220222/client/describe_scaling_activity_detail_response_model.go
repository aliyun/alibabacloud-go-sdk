// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScalingActivityDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeScalingActivityDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeScalingActivityDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeScalingActivityDetailResponseBody) *DescribeScalingActivityDetailResponse
	GetBody() *DescribeScalingActivityDetailResponseBody
}

type DescribeScalingActivityDetailResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeScalingActivityDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeScalingActivityDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingActivityDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeScalingActivityDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeScalingActivityDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeScalingActivityDetailResponse) GetBody() *DescribeScalingActivityDetailResponseBody {
	return s.Body
}

func (s *DescribeScalingActivityDetailResponse) SetHeaders(v map[string]*string) *DescribeScalingActivityDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeScalingActivityDetailResponse) SetStatusCode(v int32) *DescribeScalingActivityDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScalingActivityDetailResponse) SetBody(v *DescribeScalingActivityDetailResponseBody) *DescribeScalingActivityDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeScalingActivityDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
