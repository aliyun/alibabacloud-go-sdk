// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecommendTaskPageListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRecommendTaskPageListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRecommendTaskPageListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRecommendTaskPageListResponseBody) *DescribeRecommendTaskPageListResponse
	GetBody() *DescribeRecommendTaskPageListResponseBody
}

type DescribeRecommendTaskPageListResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRecommendTaskPageListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRecommendTaskPageListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecommendTaskPageListResponse) GoString() string {
	return s.String()
}

func (s *DescribeRecommendTaskPageListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRecommendTaskPageListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRecommendTaskPageListResponse) GetBody() *DescribeRecommendTaskPageListResponseBody {
	return s.Body
}

func (s *DescribeRecommendTaskPageListResponse) SetHeaders(v map[string]*string) *DescribeRecommendTaskPageListResponse {
	s.Headers = v
	return s
}

func (s *DescribeRecommendTaskPageListResponse) SetStatusCode(v int32) *DescribeRecommendTaskPageListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRecommendTaskPageListResponse) SetBody(v *DescribeRecommendTaskPageListResponseBody) *DescribeRecommendTaskPageListResponse {
	s.Body = v
	return s
}

func (s *DescribeRecommendTaskPageListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
