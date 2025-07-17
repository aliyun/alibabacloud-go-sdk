// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScalingGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeScalingGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeScalingGroupsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeScalingGroupsResponseBody) *DescribeScalingGroupsResponse
	GetBody() *DescribeScalingGroupsResponseBody
}

type DescribeScalingGroupsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeScalingGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeScalingGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeScalingGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeScalingGroupsResponse) GetBody() *DescribeScalingGroupsResponseBody {
	return s.Body
}

func (s *DescribeScalingGroupsResponse) SetHeaders(v map[string]*string) *DescribeScalingGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeScalingGroupsResponse) SetStatusCode(v int32) *DescribeScalingGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScalingGroupsResponse) SetBody(v *DescribeScalingGroupsResponseBody) *DescribeScalingGroupsResponse {
	s.Body = v
	return s
}

func (s *DescribeScalingGroupsResponse) Validate() error {
	return dara.Validate(s)
}
