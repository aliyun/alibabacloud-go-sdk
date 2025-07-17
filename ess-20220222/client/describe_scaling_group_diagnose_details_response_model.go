// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScalingGroupDiagnoseDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeScalingGroupDiagnoseDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeScalingGroupDiagnoseDetailsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeScalingGroupDiagnoseDetailsResponseBody) *DescribeScalingGroupDiagnoseDetailsResponse
	GetBody() *DescribeScalingGroupDiagnoseDetailsResponseBody
}

type DescribeScalingGroupDiagnoseDetailsResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeScalingGroupDiagnoseDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeScalingGroupDiagnoseDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingGroupDiagnoseDetailsResponse) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupDiagnoseDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeScalingGroupDiagnoseDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeScalingGroupDiagnoseDetailsResponse) GetBody() *DescribeScalingGroupDiagnoseDetailsResponseBody {
	return s.Body
}

func (s *DescribeScalingGroupDiagnoseDetailsResponse) SetHeaders(v map[string]*string) *DescribeScalingGroupDiagnoseDetailsResponse {
	s.Headers = v
	return s
}

func (s *DescribeScalingGroupDiagnoseDetailsResponse) SetStatusCode(v int32) *DescribeScalingGroupDiagnoseDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScalingGroupDiagnoseDetailsResponse) SetBody(v *DescribeScalingGroupDiagnoseDetailsResponseBody) *DescribeScalingGroupDiagnoseDetailsResponse {
	s.Body = v
	return s
}

func (s *DescribeScalingGroupDiagnoseDetailsResponse) Validate() error {
	return dara.Validate(s)
}
