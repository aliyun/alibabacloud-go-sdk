// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecommendVariablesVelocityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRecommendVariablesVelocityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRecommendVariablesVelocityResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRecommendVariablesVelocityResponseBody) *DescribeRecommendVariablesVelocityResponse
	GetBody() *DescribeRecommendVariablesVelocityResponseBody
}

type DescribeRecommendVariablesVelocityResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRecommendVariablesVelocityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRecommendVariablesVelocityResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecommendVariablesVelocityResponse) GoString() string {
	return s.String()
}

func (s *DescribeRecommendVariablesVelocityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRecommendVariablesVelocityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRecommendVariablesVelocityResponse) GetBody() *DescribeRecommendVariablesVelocityResponseBody {
	return s.Body
}

func (s *DescribeRecommendVariablesVelocityResponse) SetHeaders(v map[string]*string) *DescribeRecommendVariablesVelocityResponse {
	s.Headers = v
	return s
}

func (s *DescribeRecommendVariablesVelocityResponse) SetStatusCode(v int32) *DescribeRecommendVariablesVelocityResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRecommendVariablesVelocityResponse) SetBody(v *DescribeRecommendVariablesVelocityResponseBody) *DescribeRecommendVariablesVelocityResponse {
	s.Body = v
	return s
}

func (s *DescribeRecommendVariablesVelocityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
