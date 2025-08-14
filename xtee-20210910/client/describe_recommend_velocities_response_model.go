// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecommendVelocitiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRecommendVelocitiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRecommendVelocitiesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRecommendVelocitiesResponseBody) *DescribeRecommendVelocitiesResponse
	GetBody() *DescribeRecommendVelocitiesResponseBody
}

type DescribeRecommendVelocitiesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRecommendVelocitiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRecommendVelocitiesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecommendVelocitiesResponse) GoString() string {
	return s.String()
}

func (s *DescribeRecommendVelocitiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRecommendVelocitiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRecommendVelocitiesResponse) GetBody() *DescribeRecommendVelocitiesResponseBody {
	return s.Body
}

func (s *DescribeRecommendVelocitiesResponse) SetHeaders(v map[string]*string) *DescribeRecommendVelocitiesResponse {
	s.Headers = v
	return s
}

func (s *DescribeRecommendVelocitiesResponse) SetStatusCode(v int32) *DescribeRecommendVelocitiesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRecommendVelocitiesResponse) SetBody(v *DescribeRecommendVelocitiesResponseBody) *DescribeRecommendVelocitiesResponse {
	s.Body = v
	return s
}

func (s *DescribeRecommendVelocitiesResponse) Validate() error {
	return dara.Validate(s)
}
