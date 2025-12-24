// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScalingActivityStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeScalingActivityStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeScalingActivityStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeScalingActivityStatisticsResponseBody) *DescribeScalingActivityStatisticsResponse
	GetBody() *DescribeScalingActivityStatisticsResponseBody
}

type DescribeScalingActivityStatisticsResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeScalingActivityStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeScalingActivityStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingActivityStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeScalingActivityStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeScalingActivityStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeScalingActivityStatisticsResponse) GetBody() *DescribeScalingActivityStatisticsResponseBody {
	return s.Body
}

func (s *DescribeScalingActivityStatisticsResponse) SetHeaders(v map[string]*string) *DescribeScalingActivityStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeScalingActivityStatisticsResponse) SetStatusCode(v int32) *DescribeScalingActivityStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScalingActivityStatisticsResponse) SetBody(v *DescribeScalingActivityStatisticsResponseBody) *DescribeScalingActivityStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribeScalingActivityStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
