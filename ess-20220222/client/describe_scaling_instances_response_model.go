// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScalingInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeScalingInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeScalingInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeScalingInstancesResponseBody) *DescribeScalingInstancesResponse
	GetBody() *DescribeScalingInstancesResponseBody
}

type DescribeScalingInstancesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeScalingInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeScalingInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeScalingInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeScalingInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeScalingInstancesResponse) GetBody() *DescribeScalingInstancesResponseBody {
	return s.Body
}

func (s *DescribeScalingInstancesResponse) SetHeaders(v map[string]*string) *DescribeScalingInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeScalingInstancesResponse) SetStatusCode(v int32) *DescribeScalingInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScalingInstancesResponse) SetBody(v *DescribeScalingInstancesResponseBody) *DescribeScalingInstancesResponse {
	s.Body = v
	return s
}

func (s *DescribeScalingInstancesResponse) Validate() error {
	return dara.Validate(s)
}
