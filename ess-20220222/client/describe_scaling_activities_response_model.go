// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScalingActivitiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeScalingActivitiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeScalingActivitiesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeScalingActivitiesResponseBody) *DescribeScalingActivitiesResponse
	GetBody() *DescribeScalingActivitiesResponseBody
}

type DescribeScalingActivitiesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeScalingActivitiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeScalingActivitiesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingActivitiesResponse) GoString() string {
	return s.String()
}

func (s *DescribeScalingActivitiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeScalingActivitiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeScalingActivitiesResponse) GetBody() *DescribeScalingActivitiesResponseBody {
	return s.Body
}

func (s *DescribeScalingActivitiesResponse) SetHeaders(v map[string]*string) *DescribeScalingActivitiesResponse {
	s.Headers = v
	return s
}

func (s *DescribeScalingActivitiesResponse) SetStatusCode(v int32) *DescribeScalingActivitiesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScalingActivitiesResponse) SetBody(v *DescribeScalingActivitiesResponseBody) *DescribeScalingActivitiesResponse {
	s.Body = v
	return s
}

func (s *DescribeScalingActivitiesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
