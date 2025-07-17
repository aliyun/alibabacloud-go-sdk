// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScalingConfigurationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeScalingConfigurationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeScalingConfigurationsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeScalingConfigurationsResponseBody) *DescribeScalingConfigurationsResponse
	GetBody() *DescribeScalingConfigurationsResponseBody
}

type DescribeScalingConfigurationsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeScalingConfigurationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeScalingConfigurationsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingConfigurationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeScalingConfigurationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeScalingConfigurationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeScalingConfigurationsResponse) GetBody() *DescribeScalingConfigurationsResponseBody {
	return s.Body
}

func (s *DescribeScalingConfigurationsResponse) SetHeaders(v map[string]*string) *DescribeScalingConfigurationsResponse {
	s.Headers = v
	return s
}

func (s *DescribeScalingConfigurationsResponse) SetStatusCode(v int32) *DescribeScalingConfigurationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScalingConfigurationsResponse) SetBody(v *DescribeScalingConfigurationsResponseBody) *DescribeScalingConfigurationsResponse {
	s.Body = v
	return s
}

func (s *DescribeScalingConfigurationsResponse) Validate() error {
	return dara.Validate(s)
}
