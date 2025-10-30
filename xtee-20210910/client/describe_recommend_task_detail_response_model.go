// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecommendTaskDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRecommendTaskDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRecommendTaskDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRecommendTaskDetailResponseBody) *DescribeRecommendTaskDetailResponse
	GetBody() *DescribeRecommendTaskDetailResponseBody
}

type DescribeRecommendTaskDetailResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRecommendTaskDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRecommendTaskDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecommendTaskDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeRecommendTaskDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRecommendTaskDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRecommendTaskDetailResponse) GetBody() *DescribeRecommendTaskDetailResponseBody {
	return s.Body
}

func (s *DescribeRecommendTaskDetailResponse) SetHeaders(v map[string]*string) *DescribeRecommendTaskDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeRecommendTaskDetailResponse) SetStatusCode(v int32) *DescribeRecommendTaskDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRecommendTaskDetailResponse) SetBody(v *DescribeRecommendTaskDetailResponseBody) *DescribeRecommendTaskDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeRecommendTaskDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
